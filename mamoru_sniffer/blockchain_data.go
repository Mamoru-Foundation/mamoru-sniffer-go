package mamoru_sniffer

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"time"
)

type BlockchainData struct {
	*generated_bindings.FfiBlockchainDataT
}

func NewBlockchainData(data *generated_bindings.FfiBlockchainDataT) BlockchainData {
	return BlockchainData{data}
}

type BlockchainDataCtx struct {
	*generated_bindings.FfiBlockchainDataCtxT
}

type BlockchainDataCtxBuilder struct {
	*generated_bindings.FfiBlockchainDataCtxBuilderT
}

func NewBlockchainDataCtxBuilder() BlockchainDataCtxBuilder {
	builder := generated_bindings.NewBlockchainDataCtxBuilder()

	return BlockchainDataCtxBuilder{builder}
}

func (b BlockchainDataCtxBuilder) AddData(data BlockchainData) bool {
	ok := generated_bindings.BlockchainDataCtxBuilderAddData(b.FfiBlockchainDataCtxBuilderT, data.FfiBlockchainDataT)

	return ok
}

func (b BlockchainDataCtxBuilder) Finish(txId string, txHash string, txTime time.Time) BlockchainDataCtx {
	ctx := generated_bindings.BlockchainDataCtxBuilderFinish(b.FfiBlockchainDataCtxBuilderT, txId, txHash, txTime.Unix())

	return BlockchainDataCtx{ctx}
}
