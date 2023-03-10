package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkBlockchainDataCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(NewTransactionData([]Transaction{
			{
				TxIndex:    0,
				TxHash:     "",
				Type:       0,
				Nonce:      0,
				Status:     0,
				BlockIndex: 0,
				From:       "",
				To:         "",
				Value:      0,
				Fee:        0,
				GasPrice:   0,
				GasLimit:   0,
				GasUsed:    0,
				Size:       .0,
				Input:      nil,
			},
			{
				TxIndex:    1,
				TxHash:     "hash",
				Type:       1,
				Nonce:      29876,
				Status:     0,
				BlockIndex: 1,
				From:       "hash",
				To:         "",
				Value:      12345,
				Fee:        123,
				GasPrice:   0,
				GasLimit:   0,
				GasUsed:    0,
				Size:       123.1123,
				Input:      []byte{1, 2, 3, 4},
			},
		}))

		_ = builder.Finish("some-id", "some-hash")
	}
}

func TestBlockchainDataCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

	data := NewTransactionData([]Transaction{
		{
			TxIndex:    0,
			TxHash:     "",
			Type:       0,
			Nonce:      0,
			Status:     0,
			BlockIndex: 0,
			From:       "",
			To:         "",
			Value:      0,
			Fee:        0,
			GasPrice:   0,
			GasLimit:   0,
			GasUsed:    0,
			Size:       123.1123,
			Input:      []byte{1, 2, 3, 4},
		},
		{
			TxIndex:    0,
			TxHash:     "hash",
			Type:       1,
			Nonce:      29876,
			Status:     0,
			BlockIndex: 0,
			From:       "hash",
			To:         "to_hash",
			Value:      12345,
			Fee:        123,
			GasPrice:   0,
			GasLimit:   0,
			GasUsed:    0,
			Size:       123.1123,
			Input:      []byte{1, 2, 3, 4},
		},
	})

	assert.NotNil(t, data)

	builder.AddData(data)

	_ = builder.Finish("some-id", "some-hash")
}
