package ci_tests

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	answerExt   = ".a"
	testTimeout = time.Second
)

type internalTestError struct {
	reason error
}

func (e *internalTestError) Error() string {
	return "internal ci test error: " + e.reason.Error()
}

func (e *internalTestError) Unwrap() error {
	return e.reason
}

func TestService(t *testing.T) {
	t.Parallel()

	bin := filepath.Join(os.Getenv("BUILD_BIN"), "service")

	entries, err := os.ReadDir("testdata")
	require.NoError(t, err, internalTestError{
		reason: err,
	})

	cases := make([]string, 0, len(entries)/2)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) == answerExt {
			continue
		}

		cases = append(cases, filepath.Join("testdata", entry.Name()))
	}

	for ttNum, caseFile := range cases {
		t.Run(fmt.Sprintf("test-2-1-case-%d", ttNum+1), func(t *testing.T) {
			t.Parallel()

			input, err := os.Open(caseFile)
			require.NoError(t, err, internalTestError{
				reason: err,
			})
			defer input.Close()

			expected, err := os.ReadFile(caseFile + answerExt)
			require.NoError(t, err, internalTestError{
				reason: err,
			})

			ctx, closeFn := context.WithTimeout(context.TODO(), testTimeout)
			defer closeFn()

			cmd := exec.CommandContext(ctx, bin)
			cmd.Stdin = input
			output, err := cmd.CombinedOutput()
			require.NoError(t, err, "Failed to execute bin. Non-zero exit status is not expected.")
			require.Equal(t, expected, output, "Result is not equal.")
		})
	}
}
