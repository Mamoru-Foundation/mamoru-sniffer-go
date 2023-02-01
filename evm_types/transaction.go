package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type Transaction struct {
	TxIndex    uint32
	TxHash     string
	Type       uint8
	Nonce      uint64
	Status     uint64
	BlockIndex uint64
	From       string
	To         string
	Value      uint64
	Fee        uint64
	GasPrice   uint64
	GasLimit   uint64
	GasUsed    uint64
	Size       float64
	Input      []byte
}

func NewTransactionData(txs []Transaction) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewTransactionBatch()

	for _, tx := range txs {
		input := sliceToFfi(tx.Input)
		generated_bindings.TransactionBatchAppend(
			batch,
			tx.TxIndex,
			tx.TxHash,
			tx.BlockIndex,
			tx.Type,
			tx.Nonce,
			tx.Status,
			tx.From,
			tx.To,
			tx.Value,
			tx.Fee,
			tx.GasPrice,
			tx.GasLimit,
			tx.GasUsed,
			input,
			tx.Size,
		)
		freeFfiSlice(input)
	}

	data := generated_bindings.TransactionBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
