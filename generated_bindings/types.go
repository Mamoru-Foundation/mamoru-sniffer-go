// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 19 Jun 2024 22:29:03 WEST.
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

// FfiCosmosBlockchainDataBuilderT as declared in include/libmamoru_sniffer_go.h:24
type FfiCosmosBlockchainDataBuilderT C.FfiCosmosBlockchainDataBuilder_t

// FfiCosmosBlockchainDataCtxT as declared in include/libmamoru_sniffer_go.h:52
type FfiCosmosBlockchainDataCtxT C.FfiCosmosBlockchainDataCtx_t

// FfiSnifferT as declared in include/libmamoru_sniffer_go.h:60
type FfiSnifferT C.FfiSniffer_t

// SliceRefUint8T as declared in include/libmamoru_sniffer_go.h:94
type SliceRefUint8T struct {
	Ptr            *byte
	Len            uint64
	refd8594c66    *C.slice_ref_uint8_t
	allocsd8594c66 interface{}
}

// FfiSnifferResultT as declared in include/libmamoru_sniffer_go.h:195
type FfiSnifferResultT C.FfiSnifferResult_t

// FfiValueT as declared in include/libmamoru_sniffer_go.h:212
type FfiValueT C.FfiValue_t

// FfiValueDataT as declared in include/libmamoru_sniffer_go.h:214
type FfiValueDataT C.FfiValueData_t

// FfiEvmBlockchainDataBuilderT as declared in include/libmamoru_sniffer_go.h:250
type FfiEvmBlockchainDataBuilderT C.FfiEvmBlockchainDataBuilder_t

// FfiEvmBlockchainDataCtxT as declared in include/libmamoru_sniffer_go.h:274
type FfiEvmBlockchainDataCtxT C.FfiEvmBlockchainDataCtx_t
