package logger

import "log/slog"

func InitSuccess(pkg string) {
	slog.Info("Initialize successful", "package", pkg)
}
