package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level uint8

type Fields map[string]any

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarning:
		return "warning"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NewLogger(writer io.Writer, prefix string, flag int) *Logger {
	l := log.New(writer, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) CloneWithFields(f Fields) *Logger {
	ll := l.clone()

	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

func (l *Logger) CloneWithCtx(c context.Context) *Logger {
	ll := l.clone()
	ll.ctx = c
	return ll
}

func (l *Logger) CloneWithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		s := fmt.Sprintf("%s: %d %s", file, line, f.Name())
		ll.callers = []string{s}
	}
	return ll
}

func (l *Logger) CloneWithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)

	}

	ll := l.clone()
	ll.callers = callers
	return ll
}

func (l *Logger) JsonFormat(level Level, message string) Fields {
	fields := make(Fields, len(l.fields)+4)
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			fields[k] = v
		}
	}

	fields["level"] = level.String()
	fields["time"] = time.Now().Local().UnixNano()
	fields["message"] = message
	fields["callers"] = l.callers
	return fields
}

func (l *Logger) Output(level Level, message string) {
	bytes, err := json.Marshal(l.JsonFormat(level, message))
	if err != nil {
		return
	}

	content := string(bytes)
	switch level {
	case LevelDebug:
		fallthrough
	case LevelInfo:
		fallthrough
	case LevelWarning:
		fallthrough
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Debug(v ...any) {
	l.Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) DebugF(format string, v ...any) {
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...any) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) InfoF(format string, v ...any) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(v ...any) {
	l.Output(LevelWarning, fmt.Sprint(v...))
}

func (l *Logger) WarningF(format string, v ...any) {
	l.Output(LevelWarning, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...any) {
	l.Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) ErrorF(format string, v ...any) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...any) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) FatalF(format string, v ...any) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...any) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) PanicF(format string, v ...any) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}
