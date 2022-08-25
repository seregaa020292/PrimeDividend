package response

import (
	"github.com/go-chi/render"
	"net/http"
	"primedivident/internal/errors"
)

type Response struct {
	Data interface{} `json:"data"`
}

type HttpResponse struct {
	w http.ResponseWriter
	r *http.Request
}

func New(w http.ResponseWriter, r *http.Request) HttpResponse {
	return HttpResponse{
		w: w,
		r: r,
	}
}

func (h HttpResponse) SetHeader(key string, value string) HttpResponse {
	h.w.Header().Set(key, value)
	return h
}

func (h HttpResponse) Json(httpStatus int, data interface{}) {
	render.Status(h.r, httpStatus)
	render.JSON(h.w, h.r, Response{Data: data})
}

func (h HttpResponse) Respond(httpStatus int, data interface{}) {
	render.Status(h.r, httpStatus)
	render.Respond(h.w, h.r, Response{Data: data})
}

func (h HttpResponse) NoContent() {
	render.NoContent(h.w, h.r)
}

func (h HttpResponse) Err(err error) {
	var errorResponse ErrorResponse

	slugError, ok := err.(errors.SlugError)
	if !ok {
		errorResponse = InternalError(err.Error())
	} else {
		errorResponse = FindErrorType(slugError)
	}

	if err := render.Render(h.w, h.r, errorResponse); err != nil {
		panic(err)
	}
}
