package Logger

import (
	"log/slog"
	"os"
)

func GetLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
