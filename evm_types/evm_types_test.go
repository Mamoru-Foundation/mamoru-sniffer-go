package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func BenchmarkBlockchainDataCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(NewTransactionData([]Transaction{
			{
				Index:      1,
				BlockIndex: 2,
				Hash:       []byte{3, 4, 5},
				Type:       6,
				Nonce:      7,
				Status:     "completed",
				Timestamp:  42,
				From:       []byte{7, 8, 9},
				To:         []byte{7, 8, 9},
			},
			{
				Index:      1,
				BlockIndex: 2,
				Hash:       []byte{3, 4, 5},
				Type:       6,
				Nonce:      7,
				Status:     "completed",
				Timestamp:  42,
				From:       []byte{7, 8, 9},
				To:         []byte{7, 8, 9},
			},
		}))

		_ = builder.Finish("some-id", "some-hash", time.Now())
	}
}

func TestNewTransactionBatch(t *testing.T) {
	data := NewTransactionData([]Transaction{
		{
			Index:      1,
			BlockIndex: 2,
			Hash:       []byte{3, 4, 5},
			Type:       6,
			Nonce:      7,
			Status:     "completed",
			Timestamp:  42,
			From:       []byte{7, 8, 9},
			To:         []byte{7, 8, 9},
		},
		{
			Index:      1,
			BlockIndex: 2,
			Hash:       []byte{3, 4, 5},
			Type:       6,
			Nonce:      7,
			Status:     "completed",
			Timestamp:  42,
			From:       []byte{7, 8, 9},
			To:         []byte{7, 8, 9},
		},
	})

	assert.NotNil(t, data)
}

func TestBlockchainDataCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

	builder.AddData(NewTransactionData([]Transaction{
		{
			Index:      1,
			BlockIndex: 2,
			Hash:       []byte{3, 4, 5},
			Type:       6,
			Nonce:      7,
			Status:     "completed",
			Timestamp:  42,
			From:       []byte{7, 8, 9},
			To:         []byte{7, 8, 9},
		},
		{
			Index:      1,
			BlockIndex: 2,
			Hash:       []byte{3, 4, 5},
			Type:       6,
			Nonce:      7,
			Status:     "completed",
			Timestamp:  42,
			From:       []byte{7, 8, 9},
			To:         []byte{7, 8, 9},
		},
	}))

	_ = builder.Finish("some-id", "some-hash", time.Now())
}
