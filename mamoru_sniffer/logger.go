package mamoru_sniffer

import "C"
import (
	"encoding/json"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
)

type LogLevel uint8

const (
	LogLevelDebug   LogLevel = 0
	LogLevelInfo    LogLevel = 1
	LogLevelWarning LogLevel = 2
	LogLevelError   LogLevel = 3
)

type LogEntry struct {
	Level   LogLevel          `json:"level"`
	Message string            `json:"message"`
	Ctx     map[string]string `json:"ctx"`
}

var logger func(entry LogEntry)

//export mamoru_sniffer_logger
func mamoru_sniffer_logger(payload *C.char) {
	if logger == nil {
		return
	}

	entry := LogEntry{}
	err := json.Unmarshal([]byte(C.GoString(payload)), &entry)

	if err != nil {
		panic(err)
	}

	logger(entry)
}

func InitLogger(callback func(entry LogEntry)) {
	logger = callback

	generated_bindings.InitLogger()
}
