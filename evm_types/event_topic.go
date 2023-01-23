package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type EventTopic struct {
	Index uint32
	Topic string
}

func NewEventTopicData(txs []EventTopic) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewEventTopicBatch()

	for _, tx := range txs {

		generated_bindings.EventTopicBatchAppend(
			batch,
			tx.Index,
			tx.Topic,
		)
	}

	data := generated_bindings.EventTopicBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
