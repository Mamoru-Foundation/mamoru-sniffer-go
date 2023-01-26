package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func BenchmarkTransactionArgCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(NewTransactionArgData([]TransactionArg{
			{
				TxIndex: 0,
				Arg:     "",
			},
			{
				TxIndex: 1,
				Arg:     "string",
			},
		}))

		_ = builder.Finish("some-id", "some-hash", time.Now())
	}
}

func TestTransactionArgCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()
	data := NewTransactionArgData([]TransactionArg{
		{
			TxIndex: 0,
			Arg:     "",
		},
		{
			TxIndex: 1,
			Arg:     "string",
		},
	})

	assert.NotNil(t, data)

	builder.AddData(data)

	_ = builder.Finish("some-id", "some-hash", time.Now())
}
