package ci

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TESTS_COUNT = 11

func TestMyProgram(t *testing.T) {
	t.Parallel()
	cmd := exec.Command(filepath.Join(os.Getenv("BUILD_BIN"), "sevice"))
	output, err := cmd.CombinedOutput()
	assert.NoError(t, err, "Failed to execute command: %s", cmd)
	assert.Equal(t, string("Hello world!\n"), string(output))
}
