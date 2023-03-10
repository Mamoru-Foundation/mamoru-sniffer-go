package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
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

		_ = builder.Finish("some-id", "some-hash")
	}
}

func TestEventDataCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()
	data := NewEventData([]Event{
		{
			Index:       0,
			TxIndex:     0,
			Address:     "",
			BlockNumber: 0,
			TxHash:      "",
			BlockHash:   "",
			Topic0:      nil,
			Topic1:      nil,
			Topic2:      nil,
			Topic3:      nil,
			Topic4:      nil,
			Data:        nil,
		},
		{
			Index:       0,
			TxIndex:     0,
			Address:     "",
			BlockNumber: 0,
			TxHash:      "",
			BlockHash:   "",
			Topic0:      []byte{1, 2, 3, 4, 5},
			Topic1:      []byte{1, 2, 3, 4, 5},
			Topic2:      []byte{1, 2, 3, 4, 5},
			Topic3:      []byte{1, 2, 3, 4, 5},
			Topic4:      []byte{1, 2, 3, 4, 5},
			Data:        []byte{1, 2, 3, 4, 5},
		},
	})

	assert.NotNil(t, data)

	builder.AddData(data)

	_ = builder.Finish("some-id", "some-hash")
}
