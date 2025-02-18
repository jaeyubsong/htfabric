/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validator

import (
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/txmgr"
	"github.com/hyperledger/fabric/fastfabric/dependency/deptype"
)

// Validator validates the transactions present in a block and returns a batch that should be used to update the state
type Validator interface {
	ValidateAndPrepareBatch(blockAndPvtdata *ledger.BlockAndPvtData, doMVCCValidation bool, committedTxs chan<- *deptype.Transaction) (
		*privacyenabledstate.UpdateBatch, []*txmgr.TxStatInfo, error,
	)
}

// ErrPvtdataHashMissmatch is to be thrown if the hash of a collection present in the public read-write set
// does not match with the corresponding pvt data  supplied with the block for validation
type ErrPvtdataHashMissmatch struct {
	Msg string
}

func (e *ErrPvtdataHashMissmatch) Error() string {
	return e.Msg
}
