package gou

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSetupLogToFile(t *testing.T) {
	tmpf, err := os.CreateTemp("", "goutest")
	require.NoError(t, err, "error creating temp file")
	defer os.Remove(tmpf.Name())

	SetupLoggingFile(tmpf, "debug")
	const logStr = "hihi"
	Infof(logStr)

	// Flush file buffer to disk
	err = tmpf.Sync()
	require.NoError(t, err, "error syncing tmpf")
	time.Sleep(1 * time.Second)

	// Read tmp file and confirm log message was written
	bytes, err := os.ReadFile(tmpf.Name())
	require.NoError(t, err, "error reading temp file")

	logFileBytes := string(bytes)
	require.Contains(t, logFileBytes, logStr, "log message not found in temp file")
}
