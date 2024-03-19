package cosmos

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func BenchmarkCosmosCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createCosmosCtx()
	}
}

func TestCosmosCtxBuilder(t *testing.T) {
	err := os.Setenv("RUST_LOG", "info")
	require.NoError(t, err)
	t.Log("Set RUST_LOG to info")
	// This should not panic
	// Call it once to initialize the logger
	mamoru_sniffer.InitLogger(func(entry mamoru_sniffer.LogEntry) {
		// Put the entry into the project's log
		t.Log(entry)
	})
	_ = createCosmosCtx()
}

func createCosmosCtx() CosmosCtx {
	builder := NewCosmosCtxBuilder()

	builder.SetBlock(Block{
		Seq:                                0,
		Height:                             0,
		Hash:                               "some_hash",
		VersionBlock:                       0,
		VersionApp:                         0,
		ChainId:                            "",
		Time:                               0,
		LastBlockIdHash:                    "some_hash",
		LastBlockIdPartSetHeaderTotal:      0,
		LastBlockIdPartSetHeaderHash:       "some_hash",
		LastCommitHash:                     "some_hash",
		DataHash:                           "some_hash",
		ValidatorsHash:                     "some_hash",
		NextValidatorsHash:                 "some_hash",
		ConsensusHash:                      "some_hash",
		AppHash:                            "some_hash",
		LastResultsHash:                    "some_hash",
		EvidenceHash:                       "some_hash",
		ProposerAddress:                    "some_hash",
		LastCommitInfoRound:                0,
		ConsensusParamUpdatesBlockMaxBytes: 0,
		ConsensusParamUpdatesBlockMaxGas:   0,
		ConsensusParamUpdatesEvidenceMaxAgeNumBlocks: 0,
		ConsensusParamUpdatesEvidenceMaxAgeDuration:  0,
		ConsensusParamUpdatesEvidenceMaxBytes:        0,
		ConsensusParamUpdatesValidatorPubKeyTypes:    "",
		ConsensusParamUpdatesVersionApp:              0,
	})

	builder.AppendTxs([]Transaction{
		{
			Seq:  0,
			Tx:   nil,
			Code: 0,
			Data: nil,
		},
		{
			Seq:  1,
			Tx:   []byte{1, 2, 3, 4, 5},
			Code: 2,
			Data: []byte{1, 2, 3, 4, 5},
		},
	})

	builder.AppendEvents([]Event{
		{
			Seq:       0,
			EventType: "",
		},
		{
			Seq:       1,
			EventType: "string",
		},
	})

	builder.AppendMisbehaviors([]Misbehavior{
		{
			Seq:              1,
			BlockSeq:         1,
			Typ:              "some_type",
			ValidatorPower:   1,
			ValidatorAddress: "some_address",
			Height:           1,
			Time:             1,
			TotalVotingPower: 1,
		},
	})

	builder.AppendValidatorUpdates([]ValidatorUpdate{
		{
			Seq:    0,
			PubKey: nil,
			Power:  0,
		}, {
			Seq:    0,
			PubKey: []byte{1, 2, 3, 4, 5},
			Power:  0,
		},
	})

	builder.AppendEventAttributes([]EventAttribute{
		{
			Seq:      0,
			EventSeq: 0,
			Key:      "",
			Value:    "",
			Index:    false,
		},
	})

	builder.SetTxData("some-id", "some-hash")
	builder.SetBlockData("some-id", "some-hash")
	builder.SetMempoolSource()
	builder.SetStatistics(1, 2, 3, 4)

	return builder.Finish()
}

func TestCosmosCtxBuilderSetsTxData(t *testing.T) {
	builder := NewCosmosCtxBuilder()
	builder.SetTxData("txId", "txHash")
	builder.SetBlockData("blockId", "blockHash")
	builder.SetStatistics(1, 2, 3, 4)
	assert.NotNil(t, builder)
}
