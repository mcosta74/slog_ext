package slogext

import (
	"context"
	"log/slog"
)

func NewNullLogger() *slog.Logger {
	return slog.New(&nullHandler{})
}

type nullHandler struct{}

func (h *nullHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (h *nullHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *nullHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &nullHandler{}
}

func (h *nullHandler) WithGroup(name string) slog.Handler {
	return &nullHandler{}
}
