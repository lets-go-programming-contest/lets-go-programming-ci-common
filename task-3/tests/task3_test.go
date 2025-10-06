package ci

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

type value struct {
	Num  int     `json:"num_code"`
	Char string  `json:"char_code"`
	Val  float64 `json:"value"`
}

func helperDecodeJsonValues(t *testing.T, filename string) ([]value, error) {
	t.Helper()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("open file %q: %w", filename, err)
	}

	vals := make([]value, 0)
	if err := json.Unmarshal(data, &vals); err != nil {
		return nil, fmt.Errorf("unmarshal json data: %w", err)
	}

	return vals, nil
}

var casesTestService = []struct {
	name       string
	config     string
	expected   string
	output     string
	errMessage string
}{
	{
		name:       "test_case_1",
		config:     "testdata/01_config.yaml",
		expected:   "testdata/01_expected.json",
		output:     ".output/01_output.json",
		errMessage: "",
	},
	{
		name:       "test_case_2",
		config:     "testdata/02_config.yaml",
		expected:   "testdata/02_expected.json",
		output:     ".output/02_output.json",
		errMessage: "",
	},
	{
		name:       "test_case_3",
		config:     "testdata/03_config.yaml",
		expected:   "testdata/03_expected.json",
		output:     ".output/03_output.json",
		errMessage: "",
	},
	{
		name:       "test_case_4",
		config:     "testdata/04_config.yaml",
		expected:   "testdata/04_expected.json",
		output:     ".output/subdir/one_more/04_output.json",
		errMessage: "",
	},
	{
		name:       "test_case_5",
		config:     "testdata/05_config.yaml",
		expected:   "testdata/05_expected.json",
		output:     ".output/05_output.json",
		errMessage: "",
	},
	{
		name:       "test_case_6",
		config:     "testdata/06_config.yaml",
		errMessage: "no such file or directory",
	},
	{
		name:       "test_case_7",
		config:     "testdata/07_config.yaml",
		errMessage: "did not find expected key",
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

			cmd := exec.CommandContext(ctx, bin, "--config="+tt.config)

			stdErrBuffer := bytes.NewBuffer([]byte{})
			cmd.Stderr = stdErrBuffer

			err := cmd.Run()

			if tt.errMessage != "" {
				require.Contains(t, stdErrBuffer.String(), tt.errMessage)
			} else {
				require.Empty(t, stdErrBuffer.String())
				require.NoError(t, err)

				expectedValues, err := helperDecodeJsonValues(t, tt.expected)
				require.NoError(t, err, internalTestError{
					reason: err,
				})

				outputValues, err := helperDecodeJsonValues(t, tt.output)
				require.NoError(t, err, "parse output file")

				require.Len(t, outputValues, len(expectedValues), "output len not equal")

				for i := 0; i < len(expectedValues); i++ {
					require.Equalf(t, expectedValues[i].Num, outputValues[i].Num, "invlad Num for %d element", i)
					require.Equalf(t, expectedValues[i].Char, outputValues[i].Char, "invlad Char for %d element", i)
					require.Equalf(t, expectedValues[i].Val, outputValues[i].Val, "invlad Val for %d element", i)
				}
			}
		})
	}
}
