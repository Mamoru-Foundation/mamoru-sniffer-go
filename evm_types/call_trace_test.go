package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkCallTraceDataCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(NewCallTraceData([]CallTrace{
			{
				TxIndex:    0,
				BlockIndex: 0,
				Type:       "",
				From:       "",
				To:         "",
				Value:      0,
				GasLimit:   0,
				GasUsed:    0,
				Input:      []byte{1, 2, 3, 4, 5},
				Depth:      1,
			},
			{
				TxIndex:    1,
				BlockIndex: 2,
				Type:       "string",
				From:       "string",
				To:         "string",
				Value:      0,
				GasLimit:   123,
				GasUsed:    0,
				Depth:      2,
				Input:      []byte{1, 2, 3, 4, 5},
			},
		}))

		_ = builder.Finish("some-id", "some-hash")
	}
}

func TestCallTraceDataCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

	data := NewCallTraceData([]CallTrace{
		{
			Depth:      1,
			TxIndex:    0,
			BlockIndex: 0,
			Type:       "",
			From:       "",
			To:         "",
			Value:      0,
			GasLimit:   0,
			GasUsed:    0,
			Input:      []byte{1, 2, 3, 4, 5},
		},
		{
			Depth:      2,
			TxIndex:    1,
			BlockIndex: 2,
			Type:       "string",
			From:       "string",
			To:         "string",
			Value:      123,
			GasLimit:   123,
			GasUsed:    321,
			Input:      []byte{1, 2, 3, 4, 5},
		},
	})

	assert.NotNil(t, data)

	builder.AddData(data)

	_ = builder.Finish("some-id", "some-hash")
}
