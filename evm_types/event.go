package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type Event struct {
	Index       uint32
	TxIndex     uint32
	Address     string
	BlockNumber uint64
	TxHash      string
	BlockHash   string
	Topic0      []byte
	Topic1      []byte
	Topic2      []byte
	Topic3      []byte
	Topic4      []byte
	Data        []byte
}

func NewEventData(txs []Event) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewEventBatch()

	for _, tx := range txs {
		topic0 := sliceToFfi(tx.Topic0)
		topic1 := sliceToFfi(tx.Topic1)
		topic2 := sliceToFfi(tx.Topic2)
		topic3 := sliceToFfi(tx.Topic3)
		topic4 := sliceToFfi(tx.Topic4)
		data := sliceToFfi(tx.Data)
		generated_bindings.EventBatchAppend(
			batch,
			tx.Index,
			tx.Address,
			tx.BlockNumber,
			tx.TxHash,
			tx.TxIndex,
			tx.BlockHash,
			topic0,
			topic1,
			topic2,
			topic3,
			topic4,
			data,
		)
		freeFfiSlice(topic0)
		freeFfiSlice(topic1)
		freeFfiSlice(topic2)
		freeFfiSlice(topic3)
		freeFfiSlice(topic4)
		freeFfiSlice(data)
	}

	data := generated_bindings.EventBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
