package log_test

import (
	"fmt"
	"github.com/andrewmolyuk/basalt/internal/exitor"
	"github.com/andrewmolyuk/basalt/internal/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"os"
	"testing"
)

// Ensure exitorMock implements IExitor interface
var _ exitor.IExitor = (*exitorMock)(nil)

type exitorMock struct {
	mock.Mock
}

func (m *exitorMock) Exit(code int) {
	m.Called(code)
}

func catchStdOut() (*os.File, *os.File, *os.File) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	return stdout, r, w
}

func releaseStdOut(stdout *os.File, r *os.File, w *os.File) string {
	os.Stdout = stdout
	err := io.Closer.Close(w)
	if err != nil {
		panic(err)
	}
	output, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func Test_Logger_New(t *testing.T) {
	// Arrange & Act
	l := log.New(true, nil, nil)

	// Assert
	assert.NotNil(t, l, "New logger should not be nil")
}

func Test_Logger_Debug_WithDebug(t *testing.T) {
	// Arrange
	l := log.New(true, nil, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Debug("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	fmt.Println(output)
	assert.Equal(t, "[DBG] (default_test.go:59) test\x1b[0m\n", output[24:], "Logger should print debug message")
}

func Test_Logger_Debug_WithoutDebug(t *testing.T) {
	// Arrange
	l := log.New(false, nil, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Debug("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, "", output, "logger should not print debug message")
}

func Test_Logger_Info(t *testing.T) {
	// Arrange
	l := log.New(false, []string{"secret1", "secret2"}, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Info("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, "[INF] test\x1b[0m\n", output[29:], "Logger should print info message")
}

func Test_Logger_Info_WithSecrets(t *testing.T) {
	// Arrange
	l := log.New(false, []string{"secret1", "secret2"}, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Info("%s%s", "secret1", "secret2")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, "[INF] **********\x1b[0m\n", output[29:], "logger should print stars instead of secret message")
}

func Test_Logger_Warn(t *testing.T) {
	// Arrange
	l := log.New(false, nil, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Warn("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, "[WRN] test\x1b[0m\n", string(output)[29:], "Logger should print warn message")
}

func Test_Logger_Error(t *testing.T) {
	// Arrange
	e := &exitorMock{}
	e.On("Exit", 1).Return()

	l := log.New(true, []string{"secret1", "secret2"}, e)
	stdout, r, w := catchStdOut()

	// Act
	l.Error("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, "[ERR] test\x1b[0m\n", string(output)[29:], "logger should print error message")
	assert.True(t, e.AssertNumberOfCalls(t, "Exit", 1), "logger should call exit")
}
