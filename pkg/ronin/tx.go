package ronin

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Transaction struct {
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       uint64          `json:"blockNumber"`
	TimeStamp         uint64          `json:"timestamp"`
	From              common.Address  `json:"from"`
	Type              hexutil.Uint64  `json:"type"`
	ContractAddress   common.Address  `json:"contractAddress"`
	EffectiveGasPrice hexutil.Big     `json:"effectiveGasPrice"`
	Bloom             string          `json:"bloom"`
	Status            uint64          `json:"status"`
	Gas               hexutil.Uint64  `json:"gas"`
	GasPrice          *hexutil.Big    `json:"gasPrice"`
	GasUsed           uint64          `json:"gasUsed"`
	CumulativeGasUsed uint64          `json:"cumulativeGasUsed"`
	Hash              common.Hash     `json:"hash"`
	Input             string          `json:"input"`
	Nonce             hexutil.Uint64  `json:"nonce"`
	To                *common.Address `json:"to"`
	TransactionIndex  hexutil.Uint    `json:"transactionIndex"`
	Value             *hexutil.Big    `json:"value"`
	V                 *hexutil.Big    `json:"v"`
	R                 *hexutil.Big    `json:"r"`
	S                 *hexutil.Big    `json:"s"`
	PublishedTime     int64           `json:"publishedTime" rlp:"-"`
}
