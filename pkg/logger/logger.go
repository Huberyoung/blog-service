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

type Level int8

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

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithField(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s : %d %s", file, line, f.Name())}
	}

	return ll
}

func (l *Logger) WithCallerFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1

	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s : %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}

	ll := l.clone()
	ll.callers = callers
	return ll
}
func (l *Logger) JsonFormat(level Level, msg string) map[string]any {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = msg
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; ok {
				data[k] = v
			}
		}
	}
	return data
}

func (l *Logger) Output(level Level, msg string) {
	body, _ := json.Marshal(l.JsonFormat(level, msg))
	content := string(body)
	switch level {
	case LevelDebug:
		fallthrough
	case LevelInfo:
		fallthrough
	case LevelWarning:
		fallthrough
	case LevelError:
		l.newLogger.Printf(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panicf(content)
	}
}

func (l *Logger) Info(v ...any) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) InfoF(format string, v ...any) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...any) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) FatalF(format string, v ...any) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...any) {
	l.Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) PanicF(format string, v ...any) {
	l.Output(LevelPanic, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...any) {
	l.Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) ErrorF(format string, v ...any) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}
