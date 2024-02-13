package cli

import "log/slog"

var logger = slog.Default()

func SetLogger(l *slog.Logger) {
	logger = l
}

func Logger() *slog.Logger {
	return logger
}
