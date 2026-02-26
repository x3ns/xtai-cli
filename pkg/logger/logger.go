package logger

import (
	"log/slog"
	"os"
)

// Logger is the default structured logger for xtai. It wraps the standard
// library slog implementation so that call sites depend on a small surface
// area that can be swapped or configured centrally.
var Logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

