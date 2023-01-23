package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type Event struct {
	Index       uint32
	TxIndex     uint32
	Address     string
	Data        []byte
	BlockNumber uint64
	TxHash      string
	BlockHash   string
}

func NewEventData(txs []Event) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewEventBatch()

	for _, tx := range txs {
		data := sliceToFfi(tx.Data)
		generated_bindings.EventBatchAppend(
			batch,
			tx.Index,
			tx.Address,
			data,
			tx.BlockNumber,
			tx.TxHash,
			tx.TxIndex,
			tx.BlockHash,
		)
		freeFfiSlice(data)
	}

	data := generated_bindings.EventBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
