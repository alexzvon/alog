package logger

import (
	"github.com.alexzvon.alog/internal/helper"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const (
	infoLevel = "info"
	warnLevel = "warn"
	errLevel  = "err"
)

type Logger interface {
	Info(string) error
	Warn(string) error
	Error(string) error
}

type sLogger struct {
	mu sync.Mutex
}

func (l *sLogger) Info(mes string) error {
	return l.writer(mes, infoLevel)
}

func (l *sLogger) Warn(mes string) error {
	return l.writer(mes, warnLevel)
}

func (l *sLogger) Error(mes string) error {
	return l.writer(mes, errLevel)
}

func (l *sLogger) writer(mes, level string) error {
	t := time.Now().Format(time.RFC822)
	mes = helper.ConCat(t, " - ", mes, "\n")

	l.mu.Lock()
	_, err := os.Stdout.WriteString(mes)
	l.mu.Unlock()

	if err != nil {
		return errors.Wrap(err, "cannot write")
	}

	return nil
}

func New() Logger {
	return &sLogger{}
}
