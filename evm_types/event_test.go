package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func BenchmarkEventDataCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(NewEventData([]Event{
			{
				Index:       1,
				TxIndex:     1,
				Address:     "string",
				Data:        []byte{1, 2, 3, 4, 5},
				BlockNumber: 1,
				TxHash:      "string",
				BlockHash:   "string",
			},
		}))

		_ = builder.Finish("some-id", "some-hash", time.Now())
	}
}

func TestEventDataCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()
	data := NewEventData([]Event{
		{
			Index:       0,
			TxIndex:     0,
			Address:     "",
			Data:        nil,
			BlockNumber: 0,
			TxHash:      "",
			BlockHash:   "",
		},
		{
			Index:       0,
			TxIndex:     0,
			Address:     "",
			Data:        []byte{1, 2, 3, 4, 5},
			BlockNumber: 0,
			TxHash:      "",
			BlockHash:   "",
		},
	})

	assert.NotNil(t, data)

	builder.AddData(data)

	_ = builder.Finish("some-id", "some-hash", time.Now())
}
