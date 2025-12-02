package ci

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const testTimeout = time.Second

type internalTestError struct {
	reason error
}

func (e *internalTestError) Error() string {
	return "internal ci test error: " + e.reason.Error()
}

func (e *internalTestError) Unwrap() error {
	return e.reason
}

var casesTestService = []struct {
	name     string
	tag      string
	expected string
}{
	{
		name:     "test_case_1",
		tag:      "dev",
		expected: "dev debug",
	},
	{
		name:     "test_case_2",
		tag:      "",
		expected: "dev debug",
	},
}

func TestMyProgram(t *testing.T) {
	t.Parallel()

	for _, tt := range casesTestService {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bin := filepath.Join(os.Getenv("BUILD_BIN"), "service")

			ctx, closeFn := context.WithTimeout(context.TODO(), testTimeout)
			defer closeFn()

			var flag string
			if tt.tag != "" {
				flag = "-tags" + tt.tag
			}

			cmd := exec.CommandContext(ctx, bin, flag)

			stdErrBuffer := bytes.NewBuffer([]byte{})
			cmd.Stderr = stdErrBuffer

			output, err := cmd.Output()
			require.NoError(t, err, internalTestError{
				reason: err,
			})

			require.Equal(t, tt.expected, string(output), "Output is not correct. Must be \"environment log_level\".")
		})
	}
}
