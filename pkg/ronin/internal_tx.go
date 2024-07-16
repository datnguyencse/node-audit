package ronin

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type InternalTransaction struct {
	Opcode          string         `json:"opcode"`
	Order           uint64         `json:"order"`
	TransactionHash common.Hash    `json:"transactionHash"`
	Hash            common.Hash    `json:"hash"`
	Type            string         `json:"type"`
	Value           hexutil.Big    `json:"value"`
	Input           hexutil.Bytes  `json:"input"`
	Output          hexutil.Bytes  `json:"output" rlp:"-"` // rlp:- for backwards compatible while calc checksum
	From            common.Address `json:"from"`
	To              common.Address `json:"to"`
	Success         bool           `json:"success"`
	Error           string         `json:"reason"`
	Height          uint64         `json:"height"`
	BlockHash       common.Hash    `json:"blockHash"`
	// index of the internal tx in the block
	Index     uint   `json:"index"`
	TimeStamp uint64 `json:"timestamp"`
	// backward compatible for Explorer
	BlockTime     uint64 `json:"blockTime" rlp:"-"`
	PublishedTime int64  `json:"publishedTime" rlp:"-"`
}
