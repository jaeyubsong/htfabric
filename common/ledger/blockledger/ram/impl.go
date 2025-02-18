/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ramledger

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/ledger/blockledger"
	"github.com/hyperledger/fabric/infolab"
	cb "github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
	"github.com/pkg/errors"
)

var logger = flogging.MustGetLogger("common.ledger.blockledger.ram")

type cursor struct {
	list *simpleList
}

type simpleList struct {
	lock   sync.RWMutex
	next   *simpleList
	signal chan struct{}
	block  *cb.Block
}

func (s *simpleList) getNext() *simpleList {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.next
}

func (s *simpleList) setNext(n *simpleList) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.next = n
}

type ramLedger struct {
	lock    sync.RWMutex
	maxSize int
	size    int
	oldest  *simpleList
	newest  *simpleList
}

// Next blocks until there is a new block available, or returns an error if the
// next block is no longer retrievable
func (cu *cursor) Next() (*cb.Block, cb.Status) {
	// This only loops once, as signal reading indicates non-nil next
	for {
		next := cu.list.getNext()
		if next != nil {
			cu.list = next
			return cu.list.block, cb.Status_SUCCESS
		}
		<-cu.list.signal
	}
}

// Close does nothing
func (cu *cursor) Close() {}

// Iterator returns an Iterator, as specified by a ab.SeekInfo message, and its
// starting block number
func (rl *ramLedger) Iterator(startPosition *ab.SeekPosition) (blockledger.Iterator, uint64) {
	rl.lock.RLock()
	defer rl.lock.RUnlock()

	var list *simpleList
	switch start := startPosition.Type.(type) {
	case *ab.SeekPosition_Oldest:
		oldest := rl.oldest
		list = &simpleList{
			block:  &cb.Block{Header: &cb.BlockHeader{Number: oldest.block.Header.Number - 1}},
			next:   oldest,
			signal: make(chan struct{}),
		}
		close(list.signal)
	case *ab.SeekPosition_Newest:
		newest := rl.newest
		list = &simpleList{
			block:  &cb.Block{Header: &cb.BlockHeader{Number: newest.block.Header.Number - 1}},
			next:   newest,
			signal: make(chan struct{}),
		}
		close(list.signal)
	case *ab.SeekPosition_Specified:
		oldest := rl.oldest
		specified := start.Specified.Number
		logger.Debugf("Attempting to return block %d", specified)

		// Note the two +1's here is to accommodate the 'preGenesis' block of ^uint64(0)
		if specified+1 < oldest.block.Header.Number+1 || specified > rl.newest.block.Header.Number+1 {
			logger.Debugf("Returning error iterator because specified seek was %d with oldest %d and newest %d",
				specified, rl.oldest.block.Header.Number, rl.newest.block.Header.Number)
			return &blockledger.NotFoundErrorIterator{}, 0
		}

		if specified == oldest.block.Header.Number {
			list = &simpleList{
				block:  &cb.Block{Header: &cb.BlockHeader{Number: oldest.block.Header.Number - 1}},
				next:   oldest,
				signal: make(chan struct{}),
			}
			close(list.signal)
			break
		}

		list = oldest
		for {
			if list.block.Header.Number == specified-1 {
				break
			}
			list = list.getNext() // No need for nil check, because of range check above
		}
	}
	cursor := &cursor{list: list}
	blockNum := list.block.Header.Number + 1

	// If the cursor is for pre-genesis, skip it, the block number wraps
	if blockNum == ^uint64(0) {
		cursor.Next()
		blockNum++
	}

	return cursor, blockNum
}

// Height returns the number of blocks on the ledger
func (rl *ramLedger) Height() uint64 {
	rl.lock.RLock()
	defer rl.lock.RUnlock()
	return rl.newest.block.Header.Number + 1
}

// Append appends a new block to the ledger
func (rl *ramLedger) Append(block *cb.Block) error {
	if infolab.DebugMode && infolab.TraceOrdererWriteBlock {
		_, file, no, ok := runtime.Caller(1)
		if ok {
			fmt.Printf("(common/ledger/blockledger/ram/impl.go) Append called from %s#%d\n", file, no)
		}
	}
	rl.lock.Lock()
	defer rl.lock.Unlock()

	if block.Header.Number != rl.newest.block.Header.Number+1 {
		return errors.Errorf("block number should have been %d but was %d",
			rl.newest.block.Header.Number+1, block.Header.Number)
	}

	if rl.newest.block.Header.Number+1 != 0 { // Skip this check for genesis block insertion
		if !bytes.Equal(block.Header.PreviousHash, rl.newest.block.Header.Hash()) {
			return errors.Errorf("block should have had previous hash of %x but was %x",
				rl.newest.block.Header.Hash(), block.Header.PreviousHash)
		}
	}

	rl.appendBlock(block)
	return nil
}

func (rl *ramLedger) appendBlock(block *cb.Block) {
	next := &simpleList{
		signal: make(chan struct{}),
		block:  block,
	}
	rl.newest.setNext(next)

	lastSignal := rl.newest.signal
	logger.Debugf("Sending signal that block %d has a successor", rl.newest.block.Header.Number)
	rl.newest = rl.newest.getNext()
	close(lastSignal)

	rl.size++

	if rl.size > rl.maxSize {
		logger.Debugf("RAM ledger max size about to be exceeded, removing oldest item: %d",
			rl.oldest.block.Header.Number)
		rl.oldest = rl.oldest.getNext()
		rl.size--
	}
}
