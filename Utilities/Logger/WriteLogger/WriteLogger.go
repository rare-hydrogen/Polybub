package WriteLogger

import (
	"Polybub/Utilities/Logger"
	"context"
	"log/slog"
	"net/http"
)

func WriteLogger(ctx context.Context, message string, statuses ...int) {
	log := Logger.GetLogger()

	// Handle cases where statuses aren't given
	level := 0
	if len(statuses) > 0 {
		level = statuses[0]
	}

	// System Errors
	if level >= 500 {
		LogRequestWithContext(log, ctx).Error(message)
		return
	}

	// User Errors
	if level >= 300 {
		LogRequestWithContext(log, ctx).Warn(message)
		return
	}

	// Non-Errors
	LogRequestWithContext(log, ctx).Info(message)
}

// Call the logger after every request is processed
func WriteLoggerHandler(w http.ResponseWriter, r *http.Request) {
	write := func(statuses ...int) {
		WriteLogger(r.Context(), r.Response.Status, statuses...)
	}
	write()
}

func LogRequestWithContext(log *slog.Logger, ctx context.Context) *slog.Logger {
	rdv, ok := ctx.Value("requestDetails").(Logger.RequestDetails)
	if ok != true {
		// If the middleware didn't add requestDetails, log it and continue
		return slog.With(slog.String("Logging Error", "Nil Context"))
	}

	return log.With(
		slog.Int64("UserId", int64(rdv.UserId)),
		slog.String("UserName", rdv.UserName),
		slog.Int64("UserGroupId", int64(rdv.UserGroupId)),
		slog.String("Verb", rdv.Verb),
		slog.String("Endpoint", rdv.Endpoint),
		slog.String("QueryString", rdv.QueryString),
		slog.String("Body", rdv.Body),
	)
}
