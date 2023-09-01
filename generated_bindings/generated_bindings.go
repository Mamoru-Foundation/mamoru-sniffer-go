// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Fri, 01 Sep 2023 14:09:38 CEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated_bindings

/*
#cgo CFLAGS: -I${SRCDIR}/../packaged/include
#cgo LDFLAGS: -lmamoru_sniffer_go
#cgo darwin,arm64 LDFLAGS: -framework Security -L${SRCDIR}/../packaged/lib/darwin-arm64
#cgo linux,amd64 LDFLAGS: -Wl,--no-as-needed -ldl -lm -L${SRCDIR}/../packaged/lib/linux-amd64
#include <libmamoru_sniffer_go.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// NewEvmBlockchainDataBuilder function as declared in include/libmamoru_sniffer_go.h:19
func NewEvmBlockchainDataBuilder() *FfiEvmBlockchainDataBuilderT {
	__ret := C.new_evm_blockchain_data_builder()
	__v := *(**FfiEvmBlockchainDataBuilderT)(unsafe.Pointer(&__ret))
	return __v
}

// EvmBlockchainDataBuilderSetTx function as declared in include/libmamoru_sniffer_go.h:21
func EvmBlockchainDataBuilderSetTx(Builder *FfiEvmBlockchainDataBuilderT, TxId string, TxHash string) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	TxId = safeString(TxId)
	cTxId, cTxIdAllocMap := unpackPCharString(TxId)
	TxHash = safeString(TxHash)
	cTxHash, cTxHashAllocMap := unpackPCharString(TxHash)
	C.evm_blockchain_data_builder_set_tx(cBuilder, cTxId, cTxHash)
	runtime.KeepAlive(TxHash)
	runtime.KeepAlive(cTxHashAllocMap)
	runtime.KeepAlive(TxId)
	runtime.KeepAlive(cTxIdAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockchainDataBuilderSetBlock function as declared in include/libmamoru_sniffer_go.h:26
func EvmBlockchainDataBuilderSetBlock(Builder *FfiEvmBlockchainDataBuilderT, BlockId string, BlockHash string) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	BlockId = safeString(BlockId)
	cBlockId, cBlockIdAllocMap := unpackPCharString(BlockId)
	BlockHash = safeString(BlockHash)
	cBlockHash, cBlockHashAllocMap := unpackPCharString(BlockHash)
	C.evm_blockchain_data_builder_set_block(cBuilder, cBlockId, cBlockHash)
	runtime.KeepAlive(BlockHash)
	runtime.KeepAlive(cBlockHashAllocMap)
	runtime.KeepAlive(BlockId)
	runtime.KeepAlive(cBlockIdAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockchainDataBuilderSetMempoolSource function as declared in include/libmamoru_sniffer_go.h:31
func EvmBlockchainDataBuilderSetMempoolSource(Builder *FfiEvmBlockchainDataBuilderT) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	C.evm_blockchain_data_builder_set_mempool_source(cBuilder)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockchainDataBuilderSetStatistics function as declared in include/libmamoru_sniffer_go.h:38
func EvmBlockchainDataBuilderSetStatistics(Builder *FfiEvmBlockchainDataBuilderT, Blocks uint64, Transactions uint64, Events uint64, CallTraces uint64) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cBlocks, cBlocksAllocMap := (C.uint64_t)(Blocks), cgoAllocsUnknown
	cTransactions, cTransactionsAllocMap := (C.uint64_t)(Transactions), cgoAllocsUnknown
	cEvents, cEventsAllocMap := (C.uint64_t)(Events), cgoAllocsUnknown
	cCallTraces, cCallTracesAllocMap := (C.uint64_t)(CallTraces), cgoAllocsUnknown
	C.evm_blockchain_data_builder_set_statistics(cBuilder, cBlocks, cTransactions, cEvents, cCallTraces)
	runtime.KeepAlive(cCallTracesAllocMap)
	runtime.KeepAlive(cEventsAllocMap)
	runtime.KeepAlive(cTransactionsAllocMap)
	runtime.KeepAlive(cBlocksAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockchainDataBuilderFinish function as declared in include/libmamoru_sniffer_go.h:50
func EvmBlockchainDataBuilderFinish(Builder *FfiEvmBlockchainDataBuilderT) *FfiEvmBlockchainDataCtxT {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	__ret := C.evm_blockchain_data_builder_finish(cBuilder)
	runtime.KeepAlive(cBuilderAllocMap)
	__v := *(**FfiEvmBlockchainDataCtxT)(unsafe.Pointer(&__ret))
	return __v
}

// EvmSnifferObserveData function as declared in include/libmamoru_sniffer_go.h:58
func EvmSnifferObserveData(Sniffer *FfiSnifferT, Data *FfiEvmBlockchainDataCtxT) {
	cSniffer, cSnifferAllocMap := (*C.FfiSniffer_t)(unsafe.Pointer(Sniffer)), cgoAllocsUnknown
	cData, cDataAllocMap := (*C.FfiEvmBlockchainDataCtx_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	C.evm_sniffer_observe_data(cSniffer, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cSnifferAllocMap)
}

// EvmTransactionAppend function as declared in include/libmamoru_sniffer_go.h:89
func EvmTransactionAppend(Builder *FfiEvmBlockchainDataBuilderT, TxIndex uint32, TxHash string, BlockIndex uint64, Typ byte, Nonce uint64, Status uint64, From string, To string, Value uint64, Fee uint64, GasPrice uint64, GasLimit uint64, GasUsed uint64, Input SliceRefUint8T, Size float64) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cTxIndex, cTxIndexAllocMap := (C.uint32_t)(TxIndex), cgoAllocsUnknown
	TxHash = safeString(TxHash)
	cTxHash, cTxHashAllocMap := unpackPCharString(TxHash)
	cBlockIndex, cBlockIndexAllocMap := (C.uint64_t)(BlockIndex), cgoAllocsUnknown
	cTyp, cTypAllocMap := (C.uint8_t)(Typ), cgoAllocsUnknown
	cNonce, cNonceAllocMap := (C.uint64_t)(Nonce), cgoAllocsUnknown
	cStatus, cStatusAllocMap := (C.uint64_t)(Status), cgoAllocsUnknown
	From = safeString(From)
	cFrom, cFromAllocMap := unpackPCharString(From)
	To = safeString(To)
	cTo, cToAllocMap := unpackPCharString(To)
	cValue, cValueAllocMap := (C.uint64_t)(Value), cgoAllocsUnknown
	cFee, cFeeAllocMap := (C.uint64_t)(Fee), cgoAllocsUnknown
	cGasPrice, cGasPriceAllocMap := (C.uint64_t)(GasPrice), cgoAllocsUnknown
	cGasLimit, cGasLimitAllocMap := (C.uint64_t)(GasLimit), cgoAllocsUnknown
	cGasUsed, cGasUsedAllocMap := (C.uint64_t)(GasUsed), cgoAllocsUnknown
	cInput, cInputAllocMap := Input.PassValue()
	cSize, cSizeAllocMap := (C.double)(Size), cgoAllocsUnknown
	C.evm_transaction_append(cBuilder, cTxIndex, cTxHash, cBlockIndex, cTyp, cNonce, cStatus, cFrom, cTo, cValue, cFee, cGasPrice, cGasLimit, cGasUsed, cInput, cSize)
	runtime.KeepAlive(cSizeAllocMap)
	runtime.KeepAlive(cInputAllocMap)
	runtime.KeepAlive(cGasUsedAllocMap)
	runtime.KeepAlive(cGasLimitAllocMap)
	runtime.KeepAlive(cGasPriceAllocMap)
	runtime.KeepAlive(cFeeAllocMap)
	runtime.KeepAlive(cValueAllocMap)
	runtime.KeepAlive(To)
	runtime.KeepAlive(cToAllocMap)
	runtime.KeepAlive(From)
	runtime.KeepAlive(cFromAllocMap)
	runtime.KeepAlive(cStatusAllocMap)
	runtime.KeepAlive(cNonceAllocMap)
	runtime.KeepAlive(cTypAllocMap)
	runtime.KeepAlive(cBlockIndexAllocMap)
	runtime.KeepAlive(TxHash)
	runtime.KeepAlive(cTxHashAllocMap)
	runtime.KeepAlive(cTxIndexAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmCallTraceAppend function as declared in include/libmamoru_sniffer_go.h:107
func EvmCallTraceAppend(Builder *FfiEvmBlockchainDataBuilderT, Seq uint32, Depth uint32, TxIndex uint32, BlockIndex uint64, Typ string, From string, To string, Value uint64, GasLimit uint64, GasUsed uint64, Input SliceRefUint8T) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cSeq, cSeqAllocMap := (C.uint32_t)(Seq), cgoAllocsUnknown
	cDepth, cDepthAllocMap := (C.uint32_t)(Depth), cgoAllocsUnknown
	cTxIndex, cTxIndexAllocMap := (C.uint32_t)(TxIndex), cgoAllocsUnknown
	cBlockIndex, cBlockIndexAllocMap := (C.uint64_t)(BlockIndex), cgoAllocsUnknown
	Typ = safeString(Typ)
	cTyp, cTypAllocMap := unpackPCharString(Typ)
	From = safeString(From)
	cFrom, cFromAllocMap := unpackPCharString(From)
	To = safeString(To)
	cTo, cToAllocMap := unpackPCharString(To)
	cValue, cValueAllocMap := (C.uint64_t)(Value), cgoAllocsUnknown
	cGasLimit, cGasLimitAllocMap := (C.uint64_t)(GasLimit), cgoAllocsUnknown
	cGasUsed, cGasUsedAllocMap := (C.uint64_t)(GasUsed), cgoAllocsUnknown
	cInput, cInputAllocMap := Input.PassValue()
	C.evm_call_trace_append(cBuilder, cSeq, cDepth, cTxIndex, cBlockIndex, cTyp, cFrom, cTo, cValue, cGasLimit, cGasUsed, cInput)
	runtime.KeepAlive(cInputAllocMap)
	runtime.KeepAlive(cGasUsedAllocMap)
	runtime.KeepAlive(cGasLimitAllocMap)
	runtime.KeepAlive(cValueAllocMap)
	runtime.KeepAlive(To)
	runtime.KeepAlive(cToAllocMap)
	runtime.KeepAlive(From)
	runtime.KeepAlive(cFromAllocMap)
	runtime.KeepAlive(Typ)
	runtime.KeepAlive(cTypAllocMap)
	runtime.KeepAlive(cBlockIndexAllocMap)
	runtime.KeepAlive(cTxIndexAllocMap)
	runtime.KeepAlive(cDepthAllocMap)
	runtime.KeepAlive(cSeqAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockSet function as declared in include/libmamoru_sniffer_go.h:121
func EvmBlockSet(Builder *FfiEvmBlockchainDataBuilderT, BlockIndex uint64, Hash string, ParentHash string, StateRoot string, Nonce uint64, Status string, Timestamp uint64, BlockReward SliceRefUint8T, FeeRecipient string, TotalDifficulty uint64, Size float64, GasUsed uint64, GasLimit uint64) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cBlockIndex, cBlockIndexAllocMap := (C.uint64_t)(BlockIndex), cgoAllocsUnknown
	Hash = safeString(Hash)
	cHash, cHashAllocMap := unpackPCharString(Hash)
	ParentHash = safeString(ParentHash)
	cParentHash, cParentHashAllocMap := unpackPCharString(ParentHash)
	StateRoot = safeString(StateRoot)
	cStateRoot, cStateRootAllocMap := unpackPCharString(StateRoot)
	cNonce, cNonceAllocMap := (C.uint64_t)(Nonce), cgoAllocsUnknown
	Status = safeString(Status)
	cStatus, cStatusAllocMap := unpackPCharString(Status)
	cTimestamp, cTimestampAllocMap := (C.uint64_t)(Timestamp), cgoAllocsUnknown
	cBlockReward, cBlockRewardAllocMap := BlockReward.PassValue()
	FeeRecipient = safeString(FeeRecipient)
	cFeeRecipient, cFeeRecipientAllocMap := unpackPCharString(FeeRecipient)
	cTotalDifficulty, cTotalDifficultyAllocMap := (C.uint64_t)(TotalDifficulty), cgoAllocsUnknown
	cSize, cSizeAllocMap := (C.double)(Size), cgoAllocsUnknown
	cGasUsed, cGasUsedAllocMap := (C.uint64_t)(GasUsed), cgoAllocsUnknown
	cGasLimit, cGasLimitAllocMap := (C.uint64_t)(GasLimit), cgoAllocsUnknown
	C.evm_block_set(cBuilder, cBlockIndex, cHash, cParentHash, cStateRoot, cNonce, cStatus, cTimestamp, cBlockReward, cFeeRecipient, cTotalDifficulty, cSize, cGasUsed, cGasLimit)
	runtime.KeepAlive(cGasLimitAllocMap)
	runtime.KeepAlive(cGasUsedAllocMap)
	runtime.KeepAlive(cSizeAllocMap)
	runtime.KeepAlive(cTotalDifficultyAllocMap)
	runtime.KeepAlive(FeeRecipient)
	runtime.KeepAlive(cFeeRecipientAllocMap)
	runtime.KeepAlive(cBlockRewardAllocMap)
	runtime.KeepAlive(cTimestampAllocMap)
	runtime.KeepAlive(Status)
	runtime.KeepAlive(cStatusAllocMap)
	runtime.KeepAlive(cNonceAllocMap)
	runtime.KeepAlive(StateRoot)
	runtime.KeepAlive(cStateRootAllocMap)
	runtime.KeepAlive(ParentHash)
	runtime.KeepAlive(cParentHashAllocMap)
	runtime.KeepAlive(Hash)
	runtime.KeepAlive(cHashAllocMap)
	runtime.KeepAlive(cBlockIndexAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmEventAppend function as declared in include/libmamoru_sniffer_go.h:137
func EvmEventAppend(Builder *FfiEvmBlockchainDataBuilderT, Index uint32, Address string, BlockNumber uint64, TxHash string, TxIndex uint32, BlockHash string, Topic0 SliceRefUint8T, Topic1 SliceRefUint8T, Topic2 SliceRefUint8T, Topic3 SliceRefUint8T, Topic4 SliceRefUint8T, Data SliceRefUint8T) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cIndex, cIndexAllocMap := (C.uint32_t)(Index), cgoAllocsUnknown
	Address = safeString(Address)
	cAddress, cAddressAllocMap := unpackPCharString(Address)
	cBlockNumber, cBlockNumberAllocMap := (C.uint64_t)(BlockNumber), cgoAllocsUnknown
	TxHash = safeString(TxHash)
	cTxHash, cTxHashAllocMap := unpackPCharString(TxHash)
	cTxIndex, cTxIndexAllocMap := (C.uint32_t)(TxIndex), cgoAllocsUnknown
	BlockHash = safeString(BlockHash)
	cBlockHash, cBlockHashAllocMap := unpackPCharString(BlockHash)
	cTopic0, cTopic0AllocMap := Topic0.PassValue()
	cTopic1, cTopic1AllocMap := Topic1.PassValue()
	cTopic2, cTopic2AllocMap := Topic2.PassValue()
	cTopic3, cTopic3AllocMap := Topic3.PassValue()
	cTopic4, cTopic4AllocMap := Topic4.PassValue()
	cData, cDataAllocMap := Data.PassValue()
	C.evm_event_append(cBuilder, cIndex, cAddress, cBlockNumber, cTxHash, cTxIndex, cBlockHash, cTopic0, cTopic1, cTopic2, cTopic3, cTopic4, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cTopic4AllocMap)
	runtime.KeepAlive(cTopic3AllocMap)
	runtime.KeepAlive(cTopic2AllocMap)
	runtime.KeepAlive(cTopic1AllocMap)
	runtime.KeepAlive(cTopic0AllocMap)
	runtime.KeepAlive(BlockHash)
	runtime.KeepAlive(cBlockHashAllocMap)
	runtime.KeepAlive(cTxIndexAllocMap)
	runtime.KeepAlive(TxHash)
	runtime.KeepAlive(cTxHashAllocMap)
	runtime.KeepAlive(cBlockNumberAllocMap)
	runtime.KeepAlive(Address)
	runtime.KeepAlive(cAddressAllocMap)
	runtime.KeepAlive(cIndexAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// NewSniffer function as declared in include/libmamoru_sniffer_go.h:154
func NewSniffer() *FfiSnifferResultT {
	__ret := C.new_sniffer()
	__v := *(**FfiSnifferResultT)(unsafe.Pointer(&__ret))
	return __v
}

// SnifferResultGetErrorMessage function as declared in include/libmamoru_sniffer_go.h:159
func SnifferResultGetErrorMessage(Result *FfiSnifferResultT) *byte {
	cResult, cResultAllocMap := (*C.FfiSnifferResult_t)(unsafe.Pointer(Result)), cgoAllocsUnknown
	__ret := C.sniffer_result_get_error_message(cResult)
	runtime.KeepAlive(cResultAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// SnifferResultGetSniffer function as declared in include/libmamoru_sniffer_go.h:166
func SnifferResultGetSniffer(Result *FfiSnifferResultT) *FfiSnifferT {
	cResult, cResultAllocMap := (*C.FfiSnifferResult_t)(unsafe.Pointer(Result)), cgoAllocsUnknown
	__ret := C.sniffer_result_get_sniffer(cResult)
	runtime.KeepAlive(cResultAllocMap)
	__v := *(**FfiSnifferT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueData function as declared in include/libmamoru_sniffer_go.h:176
func NewValueData(Value *FfiValueT) *FfiValueDataT {
	cValue, cValueAllocMap := (*C.FfiValue_t)(unsafe.Pointer(Value)), cgoAllocsUnknown
	__ret := C.new_value_data(cValue)
	runtime.KeepAlive(cValueAllocMap)
	__v := *(**FfiValueDataT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueBool function as declared in include/libmamoru_sniffer_go.h:182
func NewValueBool(Data bool) *FfiValueT {
	cData, cDataAllocMap := (C._Bool)(Data), cgoAllocsUnknown
	__ret := C.new_value_bool(cData)
	runtime.KeepAlive(cDataAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueU64 function as declared in include/libmamoru_sniffer_go.h:185
func NewValueU64(Data uint64) *FfiValueT {
	cData, cDataAllocMap := (C.uint64_t)(Data), cgoAllocsUnknown
	__ret := C.new_value_u64(cData)
	runtime.KeepAlive(cDataAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueList function as declared in include/libmamoru_sniffer_go.h:188
func NewValueList() *FfiValueT {
	__ret := C.new_value_list()
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// ValueListAppend function as declared in include/libmamoru_sniffer_go.h:194
func ValueListAppend(ValueList *FfiValueT, Data *FfiValueT) bool {
	cValueList, cValueListAllocMap := (*C.FfiValue_t)(unsafe.Pointer(ValueList)), cgoAllocsUnknown
	cData, cDataAllocMap := (*C.FfiValue_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	__ret := C.value_list_append(cValueList, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cValueListAllocMap)
	__v := (bool)(__ret)
	return __v
}

// NewValueStruct function as declared in include/libmamoru_sniffer_go.h:198
func NewValueStruct(Ty string) *FfiValueT {
	Ty = safeString(Ty)
	cTy, cTyAllocMap := unpackPCharString(Ty)
	__ret := C.new_value_struct(cTy)
	runtime.KeepAlive(Ty)
	runtime.KeepAlive(cTyAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// ValueStructAddField function as declared in include/libmamoru_sniffer_go.h:205
func ValueStructAddField(ValueStruct *FfiValueT, Key string, Data *FfiValueT) bool {
	cValueStruct, cValueStructAllocMap := (*C.FfiValue_t)(unsafe.Pointer(ValueStruct)), cgoAllocsUnknown
	Key = safeString(Key)
	cKey, cKeyAllocMap := unpackPCharString(Key)
	cData, cDataAllocMap := (*C.FfiValue_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	__ret := C.value_struct_add_field(cValueStruct, cKey, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(Key)
	runtime.KeepAlive(cKeyAllocMap)
	runtime.KeepAlive(cValueStructAllocMap)
	__v := (bool)(__ret)
	return __v
}
