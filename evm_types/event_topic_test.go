package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func BenchmarkEventTopicDataCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(NewEventTopicData([]EventTopic{
			{
				Index: 0,
				Topic: "",
			},
			{
				Index: 1,
				Topic: "string",
			},
		}))

		_ = builder.Finish("some-id", "some-hash", time.Now())
	}
}

func TestEventTopicDataCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()
	data := NewEventTopicData([]EventTopic{
		{
			Index: 0,
			Topic: "",
		},
		{
			Index: 1,
			Topic: "string",
		},
	})

	assert.NotNil(t, data)

	builder.AddData(data)

	_ = builder.Finish("some-id", "some-hash", time.Now())
}
