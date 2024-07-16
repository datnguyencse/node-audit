package ronin

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// struct to interact with kafka
type Block struct {
	Number           uint64         `json:"number"`
	Hash             common.Hash    `json:"hash"`
	ParentHash       common.Hash    `json:"parentHash"`
	Nonce            uint64         `json:"nonce"`
	MixHash          common.Hash    `json:"mixHash"`
	LogsBloom        string         `json:"logsBloom"`
	StateRoot        common.Hash    `json:"stateRoot"`
	Miner            common.Address `json:"coinbase"`
	Difficulty       *hexutil.Big   `json:"difficulty"`
	TotalDifficulty  *hexutil.Big   `json:"totalDifficulty"`
	ExtraData        string         `json:"extraData"`
	Size             hexutil.Uint64 `json:"size"`
	GasLimit         hexutil.Uint64 `json:"gasLimit"`
	GasUsed          hexutil.Uint64 `json:"gasUsed"`
	Timestamp        hexutil.Uint64 `json:"timestamp"`
	TransactionsRoot common.Hash    `json:"transactionsRoot"`
	Transactions     []common.Hash  `json:"transactions"`
	ReceiptsRoot     common.Hash    `json:"receiptsRoot"`
}

func (b *Block) BlockNumber() uint64 {
	return b.Number
}

func (b *Block) BlockHash() common.Hash {
	return b.Hash
}

func (b *Block) BlockTimestamp() uint64 {
	return uint64(b.Timestamp)
}
