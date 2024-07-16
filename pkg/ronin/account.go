package ronin

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// this struct is copied from ronin-subscriber to work with kafka
// because ronin-subscriber is private repo so copy is easier.
type DirtyAccount struct {
	Address       common.Address `json:"address"`
	Nonce         uint64         `json:"nonce"`
	Balance       *hexutil.Big   `json:"balance"`
	Root          common.Hash    `json:"root"`
	CodeHash      common.Hash    `json:"codeHash"`
	BlockNumber   uint64         `json:"blockNumber"`
	BlockHash     common.Hash    `json:"blockHash"`
	Deleted       bool           `json:"deleted"`
	Suicided      bool           `json:"suicided"`
	DirtyCode     bool           `json:"dirtyCode"`
	Index         uint           `json:"index" rlp:"-"`
	PublishedTime int64          `json:"publishedTime" rlp:"-"`
}
