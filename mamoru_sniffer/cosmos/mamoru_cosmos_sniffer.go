package cosmos

import "C"
import (
	"fmt"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"unsafe"
)

type SnifferCosmos struct {
	*generated_bindings.FfiSnifferT
}

// Connects to Validation Chain.
// The input parameters must be passed using environment variables.
func CosmosConnect() (*SnifferCosmos, error) {
	snifferResult := generated_bindings.NewSniffer()
	errMessage := (*C.char)(unsafe.Pointer(generated_bindings.SnifferResultGetErrorMessage(snifferResult)))

	if errMessage != nil {
		return nil, fmt.Errorf("failed to create Sniffer: %s", C.GoString(errMessage))
	}

	ffiSniffer := generated_bindings.SnifferResultGetSniffer(snifferResult)

	return &SnifferCosmos{ffiSniffer}, nil
}

// Sends Cosmos data to Mamoru
func (s SnifferCosmos) ObserveCosmosData(data CosmosCtx) {
	generated_bindings.CosmosSnifferObserveData(s.FfiSnifferT, data.FfiCosmosBlockchainDataCtxT)
}
