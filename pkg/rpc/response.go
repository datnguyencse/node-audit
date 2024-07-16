package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	Call         = "CALL"
	Create       = "CREATE"
	Create2      = "CREATE2"
	DelegateCall = "DELEGATECALL"

	InternalTransactionTransfer         = "transfer"
	InternalTransactionContractCall     = "call"
	InternalTransactionContractCreation = "create"
)

type BlockResponse struct {
	Number           hexutil.Uint64   `json:"number"`
	Hash             common.Hash      `json:"hash"`
	ParentHash       common.Hash      `json:"parentHash"`
	UncleHash        common.Hash      `json:"sha3Uncles"`
	Coinbase         *common.Address  `json:"miner"`
	StateRoot        common.Hash      `json:"stateRoot"`
	ReceiptsRoot     common.Hash      `json:"receiptsRoot"`
	TransactionsRoot common.Hash      `json:"transactionsRoot"`
	MixHash          common.Hash      `json:"mixHash"`
	LogsBloom        string           `json:"logsBloom"`
	Difficulty       *hexutil.Big     `json:"difficulty"`
	Nonce            types.BlockNonce `json:"nonce"`
	GasLimit         hexutil.Uint64   `json:"gasLimit"`
	GasUsed          hexutil.Uint64   `json:"gasUsed"`
	ExtraData        string           `json:"extraData"`
	Size             hexutil.Uint64   `json:"size"`
	Timestamp        hexutil.Uint64   `json:"timestamp"` //epoch second
	TotalDifficulty  *hexutil.Big     `json:"totalDifficulty"`
	Transactions     []common.Hash    `json:"transactions"`
	Uncles           []common.Hash    `json:"uncles"`
}

func (b *BlockResponse) BlockNumber() uint64 {
	return uint64(b.Number)
}

func (b *BlockResponse) BlockHash() common.Hash {
	return b.Hash
}

func (b *BlockResponse) BlockTimestamp() uint64 {
	return uint64(b.Timestamp)
}
