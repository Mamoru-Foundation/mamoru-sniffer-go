package cosmos

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
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
		Hash:                               []byte{1, 2, 3, 4, 5},
		VersionBlock:                       0,
		VersionApp:                         0,
		ChainId:                            "",
		Time:                               0,
		LastBlockIdHash:                    nil,
		LastBlockIdPartSetHeaderTotal:      0,
		LastBlockIdPartSetHeaderHash:       nil,
		LastCommitHash:                     nil,
		DataHash:                           nil,
		ValidatorsHash:                     nil,
		NextValidatorsHash:                 nil,
		ConsensusHash:                      nil,
		AppHash:                            nil,
		LastResultsHash:                    nil,
		EvidenceHash:                       nil,
		ProposerAddress:                    nil,
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
			Seq:  0,
			Tx:   []byte{1, 2, 3, 4, 5},
			Code: 0,
			Data: []byte{1, 2, 3, 4, 5},
		},
	})

	builder.AppendEvents([]Event{
		{
			Seq:       0,
			EventType: "",
		},
		{
			Seq:       0,
			EventType: "string",
		},
	})

	builder.AppendMisbehaviors([]Misbehavior{
		{
			Seq:              0,
			BlockSeq:         0,
			Typ:              "",
			ValidatorPower:   0,
			ValidatorAddress: "some_address",
			Height:           0,
			Time:             0,
			TotalVotingPower: 0,
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
