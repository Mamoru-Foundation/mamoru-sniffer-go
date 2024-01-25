package cosmos

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type CosmosCtxBuilder struct {
	*generated_bindings.FfiCosmosBlockchainDataBuilderT
}

type CosmosCtx struct {
	*generated_bindings.FfiCosmosBlockchainDataCtxT
}

func NewCosmosCtxBuilder() CosmosCtxBuilder {
	builder := generated_bindings.NewCosmosBlockchainDataBuilder()

	return CosmosCtxBuilder{builder}
}

func (b CosmosCtxBuilder) SetTxData(txId string, txHash string) {
	generated_bindings.CosmosBlockchainDataBuilderSetTx(b.FfiCosmosBlockchainDataBuilderT, txId, txHash)
}

func (b CosmosCtxBuilder) SetBlockData(blockId string, blockHash string) {
	generated_bindings.CosmosBlockchainDataBuilderSetBlock(b.FfiCosmosBlockchainDataBuilderT, blockId, blockHash)
}

func (b CosmosCtxBuilder) SetMempoolSource() {
	generated_bindings.CosmosBlockchainDataBuilderSetMempoolSource(b.FfiCosmosBlockchainDataBuilderT)
}

func (b CosmosCtxBuilder) SetStatistics(blocks, txs, events, callTraces uint64) {
	generated_bindings.CosmosBlockchainDataBuilderSetStatistics(b.FfiCosmosBlockchainDataBuilderT, blocks, txs, events, callTraces)
}

func (b CosmosCtxBuilder) Finish() CosmosCtx {
	ctx := generated_bindings.CosmosBlockchainDataBuilderFinish(b.FfiCosmosBlockchainDataBuilderT)

	return CosmosCtx{ctx}
}

func (b CosmosCtxBuilder) SetBlock(block Block) {
	hash := mamoru_sniffer.SliceToFfi(block.Hash)
	lastBlockIdHash := mamoru_sniffer.SliceToFfi(block.LastBlockIdHash)
	lastResultsHash := mamoru_sniffer.SliceToFfi(block.LastResultsHash)
	lastCommitHash := mamoru_sniffer.SliceToFfi(block.LastCommitHash)
	appHash := mamoru_sniffer.SliceToFfi(block.AppHash)
	consensusHash := mamoru_sniffer.SliceToFfi(block.ConsensusHash)
	nextValidatorsHash := mamoru_sniffer.SliceToFfi(block.NextValidatorsHash)
	dataHash := mamoru_sniffer.SliceToFfi(block.DataHash)
	validatorsHash := mamoru_sniffer.SliceToFfi(block.ValidatorsHash)
	lastBlockIdPartSetHeaderHash := mamoru_sniffer.SliceToFfi(block.LastBlockIdPartSetHeaderHash)
	evidenceHash := mamoru_sniffer.SliceToFfi(block.EvidenceHash)
	proposerAddress := mamoru_sniffer.SliceToFfi(block.ProposerAddress)

	generated_bindings.CosmosBlockSet(
		b.FfiCosmosBlockchainDataBuilderT,
		block.Seq,
		block.Height,
		hash,
		block.VersionBlock,
		block.VersionApp,
		block.ChainId,
		block.Time,
		lastBlockIdHash,
		block.LastBlockIdPartSetHeaderTotal,
		lastBlockIdPartSetHeaderHash,
		lastCommitHash,
		dataHash,
		validatorsHash,
		nextValidatorsHash,
		consensusHash,
		appHash,
		lastResultsHash,
		evidenceHash,
		proposerAddress,
		block.LastCommitInfoRound,
		block.ConsensusParamUpdatesBlockMaxBytes,
		block.ConsensusParamUpdatesBlockMaxGas,
		block.ConsensusParamUpdatesEvidenceMaxAgeNumBlocks,
		block.ConsensusParamUpdatesEvidenceMaxAgeDuration,
		block.ConsensusParamUpdatesEvidenceMaxBytes,
		block.ConsensusParamUpdatesValidatorPubKeyTypes,
		block.ConsensusParamUpdatesVersionApp,
	)

	mamoru_sniffer.FreeFfiSlice(hash)
	mamoru_sniffer.FreeFfiSlice(lastBlockIdHash)
	mamoru_sniffer.FreeFfiSlice(lastResultsHash)
	mamoru_sniffer.FreeFfiSlice(lastCommitHash)
	mamoru_sniffer.FreeFfiSlice(appHash)
	mamoru_sniffer.FreeFfiSlice(consensusHash)
	mamoru_sniffer.FreeFfiSlice(nextValidatorsHash)
	mamoru_sniffer.FreeFfiSlice(dataHash)
	mamoru_sniffer.FreeFfiSlice(validatorsHash)
	mamoru_sniffer.FreeFfiSlice(lastBlockIdPartSetHeaderHash)
	mamoru_sniffer.FreeFfiSlice(evidenceHash)
	mamoru_sniffer.FreeFfiSlice(proposerAddress)
}

func (b CosmosCtxBuilder) AppendTxs(txs []Transaction) {
	for _, tx := range txs {
		data := mamoru_sniffer.SliceToFfi(tx.Data)
		txData := mamoru_sniffer.SliceToFfi(tx.Tx)
		generated_bindings.CosmosTransactionAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			txData,
			tx.Code,
			data,
			tx.Log,
			tx.Info,
			tx.GasUsed,
			tx.GasWanted,
			tx.Codespace,
		)
		mamoru_sniffer.FreeFfiSlice(data)
		mamoru_sniffer.FreeFfiSlice(txData)
	}
}

func (b CosmosCtxBuilder) AppendEvents(events []Event) {
	for _, event := range events {
		generated_bindings.CosmosEventAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			event.Seq,
			event.EventType,
		)
	}
}

func (b CosmosCtxBuilder) AppendEventAttributes(eventAttributes []EventAttribute) {
	for _, ea := range eventAttributes {
		generated_bindings.CosmosEventAttributeAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			ea.Seq,
			ea.EventSeq,
			ea.Key,
			ea.Value,
			ea.Index,
		)
	}
}

func (b CosmosCtxBuilder) AppendValidatorUpdates(validatorUpdates []ValidatorUpdate) {
	for _, vu := range validatorUpdates {
		pubKey := mamoru_sniffer.SliceToFfi(vu.PubKey)

		generated_bindings.CosmosValidatorUpdateAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			vu.Seq,
			pubKey,
			vu.Power,
		)
		mamoru_sniffer.FreeFfiSlice(pubKey)
	}
}

func (b CosmosCtxBuilder) AppendMisbehaviors(misbehaviors []Misbehavior) {
	for _, m := range misbehaviors {
		generated_bindings.CosmosMisbehaviorAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			m.BlockSeq,
			m.Typ,
			m.ValidatorPower,
			m.ValidatorAddress,
			m.Height,
			m.Time,
			m.TotalVotingPower,
		)
	}
}

func (b CosmosCtxBuilder) AppendVoteInfos(voteInfos []VoteInfo) {
	for _, vi := range voteInfos {
		validator := mamoru_sniffer.SliceToFfi(vi.Validator)

		generated_bindings.CosmosVoteInfosAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			vi.BlockSeq,
			validator,
			vi.SignedLastBlock,
		)
		mamoru_sniffer.FreeFfiSlice(validator)
	}
}
