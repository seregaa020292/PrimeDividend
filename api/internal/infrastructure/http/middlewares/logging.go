package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"primedivident/pkg/logger"
)

type StructuredLogger struct {
	logger.Logger
}

func newStructuredLogger() middlewareFunc {
	return middleware.RequestLogger(&StructuredLogger{
		logger.GetLogger(),
	})
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: l.Logger}
	logFields := logger.Fields{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method

	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger.ExtraFields(logFields).Infof("Request started")

	return entry
}

type StructuredLoggerEntry struct {
	logger.Logger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra any) {
	if status >= http.StatusBadRequest {
		l.Logger.ExtraFields(logger.Fields{
			"resp_status":       status,
			"resp_bytes_length": bytes,
			"resp_elapsed_ms":   float64(elapsed.Nanoseconds()) / 1000000.0,
		}).Errorf("Request fail")
	}
}

func (l *StructuredLoggerEntry) Panic(v any, stack []byte) {
	l.Logger.ExtraFields(logger.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	}).Errorf("Request panic")
}
