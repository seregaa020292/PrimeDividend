package helpers

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"primedivident/pkg/logger"
)

func GetLogEntry(r *http.Request) logger.Logger {
	entry := middleware.GetLogEntry(r).(logger.Logger)
	return entry
}
