# slogext

![test workflow](https://github.com/mcosta74/slogext/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/mcosta74/slogext.svg)](https://pkg.go.dev/github.com/mcosta74/slogext)

Extension to the standard Go log/slog package

The package offers two utilities:

- `New(io.Writer ...Option)`: creates a logger with specific options
- `NewNullLogger()`: allows to create a logger that do not produce output
