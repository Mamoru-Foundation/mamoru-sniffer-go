package tests

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer/cosmos"
	"github.com/stretchr/testify/require"
	"os"
	"sync"
	"testing"
)

var once sync.Once

func BenchmarkCosmosSnifferSmoke(b *testing.B) {
	sniffer, err := testCosmosSniffer()

	require.Nil(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx := createCosmosCtx()

		sniffer.ObserveCosmosData(ctx)
	}
}

func TestCosmosSnifferSmoke(t *testing.T) {
	_ = os.Setenv("RUST_LOG", "info")

	once.Do(func() {
		mamoru_sniffer.InitLogger(func(entry mamoru_sniffer.LogEntry) {
			t.Log(entry)
		})
	})

	sniffer, err := testCosmosSniffer()

	require.Nil(t, err)
	require.NotNil(t, sniffer)

	ctx := createCosmosCtx()

	sniffer.ObserveCosmosData(ctx)
}

var cosmosSniffer *cosmos.SnifferCosmos

func testCosmosSniffer() (*cosmos.SnifferCosmos, error) {
	_ = os.Setenv("MAMORU_CHAIN_TYPE", "OSMOSIS_TESTNET")
	_ = os.Setenv("MAMORU_ENDPOINT", "http://localhost:9090")
	_ = os.Setenv("MAMORU_PRIVATE_KEY", "SKCUszUFUg+s7eRXWPkrg0lOFwHgPpHbg8PHNuqOEk0=")
	_ = os.Setenv("MAMORU_CHAIN_ID", "validationchain")
	_ = os.Setenv("MAMORU_STATISTICS_SEND_INTERVAL_SECS", "1")

	mutex.Lock()

	var err error
	if cosmosSniffer == nil {
		cosmosSniffer, err = cosmos.CosmosConnect()
	}

	mutex.Unlock()

	return cosmosSniffer, err
}

func createCosmosCtx() cosmos.CosmosCtx {
	builder := cosmos.NewCosmosCtxBuilder()

	builder.SetBlock(cosmos.Block{
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

	builder.AppendValidatorUpdates([]cosmos.ValidatorUpdate{
		{
			Seq:    0,
			PubKey: nil,
			Power:  0,
		},
		{
			Seq:    0,
			PubKey: nil,
			Power:  0,
		},
	})

	builder.AppendMisbehaviors([]cosmos.Misbehavior{
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
		{
			Seq:              0,
			BlockSeq:         0,
			Typ:              "",
			ValidatorPower:   0,
			ValidatorAddress: "some_address2",
			Height:           0,
			Time:             0,
			TotalVotingPower: 0,
		},
	})

	builder.AppendVoteInfos([]cosmos.VoteInfo{
		{
			Seq:             0,
			BlockSeq:        0,
			Validator:       nil,
			SignedLastBlock: false,
		},
		{
			Seq:             0,
			BlockSeq:        0,
			Validator:       nil,
			SignedLastBlock: false,
		},
	})

	builder.AppendTxs([]cosmos.Transaction{
		{
			Seq:       0,
			Tx:        nil,
			Code:      0,
			Data:      nil,
			Log:       "",
			Info:      "",
			GasWanted: 0,
			GasUsed:   0,
			Codespace: "",
		},
		{
			Seq:       0,
			Tx:        nil,
			Code:      0,
			Data:      nil,
			Log:       "",
			Info:      "",
			GasWanted: 0,
			GasUsed:   0,
			Codespace: "",
		},
	})

	builder.AppendEvents([]cosmos.Event{
		{
			Seq:       0,
			EventType: "",
		},
		{
			Seq:       0,
			EventType: "",
		},
	})

	builder.AppendEventAttributes([]cosmos.EventAttribute{
		{
			Seq:      0,
			EventSeq: 0,
			Key:      "",
			Value:    "",
			Index:    false,
		},
		{
			Seq:      0,
			EventSeq: 0,
			Key:      "",
			Value:    "",
			Index:    false,
		},
	})
	builder.SetMempoolSource()

	builder.SetStatistics(1, 20, 30, 40)
	builder.SetBlockData("1", "0xtest_block_hash")
	builder.SetTxData("1", "0xtest_block_hash")

	return builder.Finish()
}
