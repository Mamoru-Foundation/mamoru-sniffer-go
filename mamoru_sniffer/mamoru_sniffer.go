package mamoru_sniffer

import "C"
import (
	"fmt"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"unsafe"
)

type Sniffer struct {
	*generated_bindings.FfiSnifferT
}

// Connects to Validation Chain.
// The input parameters must be passed using environment variables.
func Connect() (*Sniffer, error) {
	snifferResult := generated_bindings.NewSniffer()
	errMessage := (*C.char)(unsafe.Pointer(generated_bindings.SnifferResultGetErrorMessage(snifferResult)))

	if errMessage != nil {
		return nil, fmt.Errorf("failed to create Sniffer: %s", C.GoString(errMessage))
	}

	ffiSniffer := generated_bindings.SnifferResultGetSniffer(snifferResult)

	return &Sniffer{ffiSniffer}, nil
}

// Sends EVM data to Mamoru
func (s Sniffer) ObserveEvmData(data EvmCtx) {
	generated_bindings.EvmSnifferObserveData(s.FfiSnifferT, data.FfiEvmBlockchainDataCtxT)
}
