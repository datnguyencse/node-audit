package ronin

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Log struct {
	Address       common.Address `json:"address"`
	Topics        []common.Hash  `json:"topics"`
	Data          hexutil.Bytes  `json:"data"`
	BlockNumber   uint64         `json:"blockNumber"`
	TxHash        common.Hash    `json:"transactionHash"`
	TxIndex       uint           `json:"transactionIndex"`
	BlockHash     common.Hash    `json:"blockHash"`
	Index         uint           `json:"logIndex"`
	Removed       bool           `json:"removed"`
	TimeStamp     uint64         `json:"timestamp"`
	PublishedTime int64          `json:"publishedTime" rlp:"-"`
}
