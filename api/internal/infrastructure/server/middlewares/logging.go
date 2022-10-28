package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"primedividend/api/internal/config/consts"
	"primedividend/api/pkg/logger"
	"primedividend/api/pkg/utils/gog"
)

type (
	StructLogger struct {
		logger.Logger
	}
	LoggerEntry struct {
		requestFields logger.Fields
		logger.Logger
	}
)

func NewStructLogger() func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructLogger{
		Logger: logger.GetLogger(),
	})
}

func (l *StructLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	scheme := gog.If(r.TLS != nil, "https", "http")

	logFields := logger.Fields{
		"http_scheme": scheme,
		"http_proto":  r.Proto,
		"http_method": r.Method,
		"remote_addr": r.RemoteAddr,
		"user_agent":  r.UserAgent(),
		"uri":         fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
		"time":        time.Now().Format(consts.TimestampFormat),
	}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	return &LoggerEntry{
		Logger:        l.Logger,
		requestFields: logFields,
	}
}

func (l *LoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra any) {
	if status >= http.StatusBadRequest {
		l.Logger.
			ExtraFields(l.requestFields).
			ExtraFields(logger.Fields{
				"resp_status":       status,
				"resp_bytes_length": bytes,
				"resp_elapsed_ms":   float64(elapsed.Nanoseconds()) / 1000000.0,
				"resp_time":         time.Now().Format(consts.TimestampFormat),
			}).Errorf("Request fail")
	}
}

func (l *LoggerEntry) Panic(v any, stack []byte) {
	l.Logger.
		ExtraFields(l.requestFields).
		ExtraFields(logger.Fields{
			"stack": string(stack),
			"panic": fmt.Sprintf("%+v", v),
			"time":  time.Now().Format(consts.TimestampFormat),
		}).Errorf("Request panic")
}
