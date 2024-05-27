package safelog

import (
	"context"
	"log/slog"
)

type (
	Logger struct {
		sll *slog.Logger
	}
	logMsg  string
	keyAttr string

	Attr struct {
		Key   keyAttr
		Value slog.LogValuer
	}
)

func NewLogger(sll *slog.Logger) *Logger {
	return &Logger{sll: sll}
}

func (l Logger) Log(ctx context.Context, lv slog.Level, msg logMsg, attrs ...Attr) {
	vals := make([]slog.Attr, len(attrs))
	for i, attr := range attrs {
		vals[i] = slog.Attr{
			Key:   string(attr.Key),
			Value: attr.Value.LogValue(),
		}
	}

	l.sll.LogAttrs(ctx, lv, string(msg), vals...)
}
