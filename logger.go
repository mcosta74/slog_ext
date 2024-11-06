// Package slogext exposes some utilities to create slog.Logger.
package slogext

import (
	"io"
	"log/slog"
	"path/filepath"
)

type logOptions struct {
	level         slog.Level
	useUTC        bool
	addSource     bool
	addSourcePath bool
	useJSON       bool
}

// Option is a function on logger options
type Option func(*logOptions)

// WithLevel is a helper option to provide the minimum log level
func WithLevel(lvl slog.Level) Option {
	return func(lo *logOptions) {
		lo.level = lvl
	}
}

// WithUseUTC is a helper option to configure if use UTC for log messages' timestamp
func WithUseUTC(useUTC bool) Option {
	return func(lo *logOptions) {
		lo.useUTC = useUTC
	}
}

// WithSource is a helper option to configure if add source information in log messages
func WithSource(addSource bool) Option {
	return func(lo *logOptions) {
		lo.addSource = addSource
	}
}

// WithSourcePath is a helper option to configure if add the absolute path for source info
func WithSourcePath(addSourcePath bool) Option {
	return func(lo *logOptions) {
		lo.addSourcePath = addSourcePath
	}
}

// WithJSON is a helper option to configure if use JSON format for log messages
func WithJSON(useJSON bool) Option {
	return func(lo *logOptions) {
		lo.useJSON = useJSON
	}
}

// New creates a new slog.Logger.
// The new logger will use writer as output destination, default configurations can be modified using opts
func New(writer io.Writer, opts ...Option) *slog.Logger {
	options := &logOptions{}

	for _, opt := range opts {
		opt(options)
	}

	handlerOptions := &slog.HandlerOptions{
		Level:     options.level,
		AddSource: options.addSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && options.useUTC {
				a.Value = slog.TimeValue(a.Value.Time().UTC())
			}

			if a.Key == slog.SourceKey && !options.addSourcePath {
				source := a.Value.Any().(*slog.Source)
				source.File = filepath.Base(source.File)
			}
			return a
		},
	}

	var handler slog.Handler
	if options.useJSON {
		handler = slog.NewJSONHandler(writer, handlerOptions)
	} else {
		handler = slog.NewTextHandler(writer, handlerOptions)
	}
	return slog.New(handler)
}
