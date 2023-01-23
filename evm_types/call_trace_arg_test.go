package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func BenchmarkCallTraceArgDataCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(NewCallTraceArgData([]CallTraceArg{
			{
				CallTraceSeq: 1,
				Depth:        0,
				TxIndex:      1,
				BlockIndex:   1,
				Arg:          "string",
			},
			{
				CallTraceSeq: 0,
				Depth:        0,
				TxIndex:      0,
				BlockIndex:   0,
				Arg:          "",
			},
		}))

		_ = builder.Finish("some-id", "some-hash", time.Now())
	}
}

func TestCallTraceArgDataCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()
	data := NewCallTraceArgData([]CallTraceArg{
		{
			CallTraceSeq: 1,
			TxIndex:      1,
			BlockIndex:   1,
			Depth:        1,
			Arg:          "string",
		},
		{
			CallTraceSeq: 2,
			Depth:        0,
			TxIndex:      0,
			BlockIndex:   0,
			Arg:          "",
		},
	})

	assert.NotNil(t, data)

	builder.AddData(data)

	_ = builder.Finish("some-id", "some-hash", time.Now())
}
