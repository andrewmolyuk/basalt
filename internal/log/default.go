package log

import (
	"fmt"
	"github.com/andrewmolyuk/basalt/internal/exitor"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

const (
	levelDBG = "DBG"
	levelINF = "INF"
	levelWRN = "WRN"
	levelERR = "ERR"
)

var colors = map[string]string{
	levelDBG: "",
	levelINF: colorCyan,
	levelWRN: colorYellow,
	levelERR: colorRed,
}

// Ensure defaultLogger implements ILog interface
var _ ILogger = (*defaultLogger)(nil)

type defaultLogger struct {
	secrets   []string
	debugMode bool
	lock      sync.Mutex
	exitor    exitor.IExitor
}

func (l *defaultLogger) print(level string, args ...interface{}) {
	output := ""

	if len(args) == 1 {
		output = fmt.Sprint(args...)
	} else {
		format := fmt.Sprintf("%s", args[0])
		output = fmt.Sprintf(format, args[1:]...)
	}

	if l.debugMode && level == levelDBG {
		skip := 2
		_, file, no, ok := runtime.Caller(skip)
		for ok {
			if !strings.Contains(file, "/pkg/log/") {
				file = strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
				output = fmt.Sprintf("(%s:%d) %s", file, no, output)
				break
			}
			skip++
			_, file, no, ok = runtime.Caller(skip)
		}
	}

	output = fmt.Sprintf("%s [%s] %s", time.Now().Format("2006/01/02 15:04:05.000"), level, output)

	for _, secret := range l.secrets {
		output = strings.Replace(output, secret, "*****", -1)
	}

	output = fmt.Sprint(colors[level], output, colorReset)

	l.lock.Lock()
	defer l.lock.Unlock()
	fmt.Println(output)
}

// Debug prints debug message if debug mode is enabled
func (l *defaultLogger) Debug(args ...interface{}) {
	if l.debugMode {
		l.print(levelDBG, args...)
	}
}

// Info prints info message
func (l *defaultLogger) Info(args ...interface{}) {
	l.print(levelINF, args...)
}

// Warn prints warning message
func (l *defaultLogger) Warn(args ...interface{}) {
	l.print(levelWRN, args...)
}

// Error prints error message and exits with code 1
func (l *defaultLogger) Error(args ...interface{}) {
	l.print(levelERR, args...)
	l.exitor.Exit(1)
}

// New creates a new instance of defaultLogger implementing ILogger interface
func New(debugMode bool, secrets []string, exitor exitor.IExitor) ILogger {
	return &defaultLogger{
		secrets:   secrets,
		debugMode: debugMode,
		exitor:    exitor,
	}
}
