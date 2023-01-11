package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type Transaction struct {
	Index      uint64
	BlockIndex uint64
	Hash       []byte
	Type       byte
	Nonce      uint32
	Status     string
	Timestamp  uint32
	From       []byte
	To         []byte
}

func NewTransactionData(txs []Transaction) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewTransactionBatch()

	for _, tx := range txs {
		hash := sliceToFfi(tx.Hash)
		from := sliceToFfi(tx.From)
		to := sliceToFfi(tx.To)

		generated_bindings.TransactionBatchAppend(
			batch,
			tx.Index,
			tx.BlockIndex,
			hash,
			tx.Type,
			tx.Nonce,
			tx.Status,
			tx.Timestamp,
			from,
			to,
		)

		freeFfiSlice(hash)
		freeFfiSlice(from)
		freeFfiSlice(to)
	}

	data := generated_bindings.TransactionBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
