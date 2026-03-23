package bor

import (
	"math/big"

	"github.com/tenderly/polygon-bor/common"
	"github.com/tenderly/polygon-bor/consensus/bor/clerk"
	"github.com/tenderly/polygon-bor/consensus/bor/statefull"
	"github.com/tenderly/polygon-bor/core/state"
	"github.com/tenderly/polygon-bor/core/types"
	"github.com/tenderly/polygon-bor/core/vm"
)

//go:generate mockgen -destination=./genesis_contract_mock.go -package=bor . GenesisContract
type GenesisContract interface {
	CommitState(event *clerk.EventRecordWithTime, state vm.StateDB, header *types.Header, chCtx statefull.ChainContext) (uint64, error)
	LastStateId(state *state.StateDB, number uint64, hash common.Hash) (*big.Int, error)
}
