// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 11 Jan 2023 19:38:03 CET.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated_bindings

/*
#cgo CFLAGS: -I${SRCDIR}/../packaged/include
#cgo LDFLAGS: -lmamoru_sniffer_go
#cgo darwin arm64 LDFLAGS: -framework Security -L${SRCDIR}/../packaged/lib/darwin-arm64
#cgo linux amd64 LDFLAGS: -Wl,--no-as-needed -ldl -lm -L${SRCDIR}/../packaged/lib/linux-amd64
#include <libmamoru_sniffer_go.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// NewTransactionBatch function as declared in include/libmamoru_sniffer_go.h:19
func NewTransactionBatch() *FfiTransactionBatchT {
	__ret := C.new_transaction_batch()
	__v := *(**FfiTransactionBatchT)(unsafe.Pointer(&__ret))
	return __v
}

// TransactionBatchAppend function as declared in include/libmamoru_sniffer_go.h:52
func TransactionBatchAppend(Batch *FfiTransactionBatchT, Index uint64, BlockIndex uint64, Hash SliceRefUint8T, Typ byte, Nonce uint32, Status string, Timestamp uint32, From SliceRefUint8T, To SliceRefUint8T) {
	cBatch, cBatchAllocMap := (*C.FfiTransactionBatch_t)(unsafe.Pointer(Batch)), cgoAllocsUnknown
	cIndex, cIndexAllocMap := (C.uint64_t)(Index), cgoAllocsUnknown
	cBlockIndex, cBlockIndexAllocMap := (C.uint64_t)(BlockIndex), cgoAllocsUnknown
	cHash, cHashAllocMap := Hash.PassValue()
	cTyp, cTypAllocMap := (C.uint8_t)(Typ), cgoAllocsUnknown
	cNonce, cNonceAllocMap := (C.uint32_t)(Nonce), cgoAllocsUnknown
	Status = safeString(Status)
	cStatus, cStatusAllocMap := unpackPCharString(Status)
	cTimestamp, cTimestampAllocMap := (C.uint32_t)(Timestamp), cgoAllocsUnknown
	cFrom, cFromAllocMap := From.PassValue()
	cTo, cToAllocMap := To.PassValue()
	C.transaction_batch_append(cBatch, cIndex, cBlockIndex, cHash, cTyp, cNonce, cStatus, cTimestamp, cFrom, cTo)
	runtime.KeepAlive(cToAllocMap)
	runtime.KeepAlive(cFromAllocMap)
	runtime.KeepAlive(cTimestampAllocMap)
	runtime.KeepAlive(Status)
	runtime.KeepAlive(cStatusAllocMap)
	runtime.KeepAlive(cNonceAllocMap)
	runtime.KeepAlive(cTypAllocMap)
	runtime.KeepAlive(cHashAllocMap)
	runtime.KeepAlive(cBlockIndexAllocMap)
	runtime.KeepAlive(cIndexAllocMap)
	runtime.KeepAlive(cBatchAllocMap)
}

// TransactionBatchFinish function as declared in include/libmamoru_sniffer_go.h:69
func TransactionBatchFinish(Object *FfiTransactionBatchT) *FfiBlockchainDataT {
	cObject, cObjectAllocMap := (*C.FfiTransactionBatch_t)(unsafe.Pointer(Object)), cgoAllocsUnknown
	__ret := C.transaction_batch_finish(cObject)
	runtime.KeepAlive(cObjectAllocMap)
	__v := *(**FfiBlockchainDataT)(unsafe.Pointer(&__ret))
	return __v
}

// NewSniffer function as declared in include/libmamoru_sniffer_go.h:74
func NewSniffer() *FfiSnifferResultT {
	__ret := C.new_sniffer()
	__v := *(**FfiSnifferResultT)(unsafe.Pointer(&__ret))
	return __v
}

// SnifferResultGetErrorMessage function as declared in include/libmamoru_sniffer_go.h:79
func SnifferResultGetErrorMessage(Result *FfiSnifferResultT) *byte {
	cResult, cResultAllocMap := (*C.FfiSnifferResult_t)(unsafe.Pointer(Result)), cgoAllocsUnknown
	__ret := C.sniffer_result_get_error_message(cResult)
	runtime.KeepAlive(cResultAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// SnifferResultGetSniffer function as declared in include/libmamoru_sniffer_go.h:88
func SnifferResultGetSniffer(Result *FfiSnifferResultT) *FfiSnifferT {
	cResult, cResultAllocMap := (*C.FfiSnifferResult_t)(unsafe.Pointer(Result)), cgoAllocsUnknown
	__ret := C.sniffer_result_get_sniffer(cResult)
	runtime.KeepAlive(cResultAllocMap)
	__v := *(**FfiSnifferT)(unsafe.Pointer(&__ret))
	return __v
}

// SnifferObserveData function as declared in include/libmamoru_sniffer_go.h:96
func SnifferObserveData(Sniffer *FfiSnifferT, Data *FfiBlockchainDataCtxT) {
	cSniffer, cSnifferAllocMap := (*C.FfiSniffer_t)(unsafe.Pointer(Sniffer)), cgoAllocsUnknown
	cData, cDataAllocMap := (*C.FfiBlockchainDataCtx_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	C.sniffer_observe_data(cSniffer, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cSnifferAllocMap)
}

// NewBlockchainDataCtxBuilder function as declared in include/libmamoru_sniffer_go.h:102
func NewBlockchainDataCtxBuilder() *FfiBlockchainDataCtxBuilderT {
	__ret := C.new_blockchain_data_ctx_builder()
	__v := *(**FfiBlockchainDataCtxBuilderT)(unsafe.Pointer(&__ret))
	return __v
}

// BlockchainDataCtxBuilderAddData function as declared in include/libmamoru_sniffer_go.h:111
func BlockchainDataCtxBuilderAddData(Builder *FfiBlockchainDataCtxBuilderT, Data *FfiBlockchainDataT) bool {
	cBuilder, cBuilderAllocMap := (*C.FfiBlockchainDataCtxBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cData, cDataAllocMap := (*C.FfiBlockchainData_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	__ret := C.blockchain_data_ctx_builder_add_data(cBuilder, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
	__v := (bool)(__ret)
	return __v
}

// BlockchainDataCtxBuilderFinish function as declared in include/libmamoru_sniffer_go.h:115
func BlockchainDataCtxBuilderFinish(Builder *FfiBlockchainDataCtxBuilderT, TxId string, TxHash string, TxTimestampSecs int64) *FfiBlockchainDataCtxT {
	cBuilder, cBuilderAllocMap := (*C.FfiBlockchainDataCtxBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	TxId = safeString(TxId)
	cTxId, cTxIdAllocMap := unpackPCharString(TxId)
	TxHash = safeString(TxHash)
	cTxHash, cTxHashAllocMap := unpackPCharString(TxHash)
	cTxTimestampSecs, cTxTimestampSecsAllocMap := (C.int64_t)(TxTimestampSecs), cgoAllocsUnknown
	__ret := C.blockchain_data_ctx_builder_finish(cBuilder, cTxId, cTxHash, cTxTimestampSecs)
	runtime.KeepAlive(cTxTimestampSecsAllocMap)
	runtime.KeepAlive(TxHash)
	runtime.KeepAlive(cTxHashAllocMap)
	runtime.KeepAlive(TxId)
	runtime.KeepAlive(cTxIdAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
	__v := *(**FfiBlockchainDataCtxT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueData function as declared in include/libmamoru_sniffer_go.h:128
func NewValueData(Value *FfiValueT) *FfiValueDataT {
	cValue, cValueAllocMap := (*C.FfiValue_t)(unsafe.Pointer(Value)), cgoAllocsUnknown
	__ret := C.new_value_data(cValue)
	runtime.KeepAlive(cValueAllocMap)
	__v := *(**FfiValueDataT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueBool function as declared in include/libmamoru_sniffer_go.h:131
func NewValueBool(Data bool) *FfiValueT {
	cData, cDataAllocMap := (C._Bool)(Data), cgoAllocsUnknown
	__ret := C.new_value_bool(cData)
	runtime.KeepAlive(cDataAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueU64 function as declared in include/libmamoru_sniffer_go.h:134
func NewValueU64(Data uint64) *FfiValueT {
	cData, cDataAllocMap := (C.uint64_t)(Data), cgoAllocsUnknown
	__ret := C.new_value_u64(cData)
	runtime.KeepAlive(cDataAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueList function as declared in include/libmamoru_sniffer_go.h:137
func NewValueList() *FfiValueT {
	__ret := C.new_value_list()
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// ValueListAppend function as declared in include/libmamoru_sniffer_go.h:143
func ValueListAppend(ValueList *FfiValueT, Data *FfiValueT) bool {
	cValueList, cValueListAllocMap := (*C.FfiValue_t)(unsafe.Pointer(ValueList)), cgoAllocsUnknown
	cData, cDataAllocMap := (*C.FfiValue_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	__ret := C.value_list_append(cValueList, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cValueListAllocMap)
	__v := (bool)(__ret)
	return __v
}

// NewValueStruct function as declared in include/libmamoru_sniffer_go.h:147
func NewValueStruct(Ty string) *FfiValueT {
	Ty = safeString(Ty)
	cTy, cTyAllocMap := unpackPCharString(Ty)
	__ret := C.new_value_struct(cTy)
	runtime.KeepAlive(Ty)
	runtime.KeepAlive(cTyAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// ValueStructAddField function as declared in include/libmamoru_sniffer_go.h:154
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