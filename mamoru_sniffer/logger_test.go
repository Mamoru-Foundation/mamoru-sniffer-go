package mamoru_sniffer

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestInitLogger(t *testing.T) {
	err := os.Setenv("RUST_LOG", "info")
	require.NoError(t, err)

	// This should not panic
	// Call it once to initialize the logger
	InitLogger(func(entry LogEntry) {
		// Put the entry into the project's log
		t.Log(entry)
	})
}
