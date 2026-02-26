package logger

import (
	"log/slog"
	"os"
)

func New(dev bool) *slog.Logger {
	if dev {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	}

	// Could turn into structured logging in prod
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
