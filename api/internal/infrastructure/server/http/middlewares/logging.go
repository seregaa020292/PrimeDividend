package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"primedivident/pkg/logger"
)

type StructuredLogger struct {
	logger logger.Logger
}

func NewStructuredLogger() func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{
		logger: logger.GetLogger(),
	})
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{logger: l.logger}
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

	entry.logger.ExtraFields(logFields).Infof("Request started")

	return entry
}

type StructuredLoggerEntry struct {
	logger logger.Logger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	if status > http.StatusBadRequest {
		l.logger.ExtraFields(logger.Fields{
			"resp_status":       status,
			"resp_bytes_length": bytes,
			"resp_elapsed_ms":   float64(elapsed.Nanoseconds()) / 1000000.0,
		}).Errorf("Request fail")
	}
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.logger.ExtraFields(logger.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	}).Errorf("Request panic")
}

func GetLogEntry(r *http.Request) logger.Logger {
	entry := middleware.GetLogEntry(r).(*StructuredLoggerEntry)
	return entry.logger
}
