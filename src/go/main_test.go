package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	tests := []struct {
		Name           string
		Args           []string
		ExpectedStdout string
		ExpectedStderr string
	}{
		{
			Name: "when valid inputs are supplied for the participation command",
			Args: []string{"../../example-data/survey.csv", "participation"},
			ExpectedStdout: heredoc.Doc(`
				Participation

				Participants: 6
				Submitted: 5 (83.3%)
			`),
			ExpectedStderr: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// Setup
			mockStdout, readMockStdout := newMockPipe(t)
			mockStderr, readMockStderr := newMockPipe(t)

			fullArgs := []string{"main"}
			fullArgs = append(fullArgs, tc.Args...)

			// Execute
			Run(fullArgs, mockStdout, mockStderr)

			// Assert
			assert.Equal(t, tc.ExpectedStdout, readMockStdout())
			assert.Equal(t, tc.ExpectedStderr, readMockStderr())
		})
	}
}

// newMockPipe creates an *os.File that can be used in place of os.Stdout
// and os.Stderr. The *os.File instance is returned, along with a func
// that returns the contents of the pipe as a string.
func newMockPipe(t *testing.T) (writer *os.File, outputReadFunc func() string) {
	t.Helper()
	r, w, err := os.Pipe()
	require.Nil(t, err)

	return w, func() string {
		err = w.Close()
		require.Nil(t, err)

		output, err := ioutil.ReadAll(r)
		require.Nil(t, err)

		return string(output)
	}
}
