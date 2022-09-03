package response

import (
	"net/http"

	"github.com/go-chi/render"

	"primedivident/internal/infrastructure/server/http/middlewares"
	"primedivident/internal/mistakes"
)

type Response struct {
	Data interface{} `json:"data"`
}

type Respond struct {
	w http.ResponseWriter
	r *http.Request
}

func New(w http.ResponseWriter, r *http.Request) Respond {
	return Respond{
		w: w,
		r: r,
	}
}

func (h Respond) SetHeader(key string, value string) Respond {
	h.w.Header().Set(key, value)
	return h
}

func (h Respond) Json(httpStatus int, data interface{}) {
	render.Status(h.r, httpStatus)
	render.JSON(h.w, h.r, Response{Data: data})
}

func (h Respond) Any(httpStatus int, data interface{}) {
	render.Status(h.r, httpStatus)
	render.Respond(h.w, h.r, Response{Data: data})
}

func (h Respond) NoContent() {
	render.NoContent(h.w, h.r)
}

func (h Respond) Err(err error) {
	var errorResponse ErrorRespond

	slugError, ok := err.(mistakes.SlugError)
	if !ok {
		errorResponse = InternalError(err.Error())
	} else {
		errorResponse = FindErrorType(slugError)
	}

	middlewares.GetLogEntry(h.r).
		ExtraError(err).
		ExtraField("error-slug", slugError.Slug()).
		Warnf(slugError.Error())

	if err := render.Render(h.w, h.r, errorResponse); err != nil {
		panic(err)
	}
}
