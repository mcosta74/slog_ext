# slogext

![test workflow](https://github.com/mcosta74/slogext/actions/workflows/test.yml/badge.svg)

Extension to the standard Go log/slog package

The package offers two utilites:

- `New(io.Writer ...Option)`: creates a logger with specific options
- `NewNullLogger()`: allows to create a logger that do not produce output
