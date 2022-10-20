package response

import (
	"net/http"

	"github.com/go-chi/render"

	"primedivident/pkg/logger"
	"primedivident/pkg/utils/gog"
	"primedivident/pkg/validator"
)

type Responder interface {
	Http(w http.ResponseWriter, r *http.Request) Responder
	SetHeader(key string, value string) Responder
	WriteHeader(httpStatus int)
	Json(httpStatus int, data any, meta ...any)
	Any(httpStatus int, data any, meta ...any)
	Err(err error)
	Redirect(url string, httpStatus ...int)
	Decode(v any) error
	DecodeValidate(v any) error
}

type (
	Response struct {
		Data any `json:"data"`
		Meta any `json:"meta,omitempty"`
	}
	Respond struct {
		logger    logger.Logger
		validator validator.Validator
		writer    http.ResponseWriter
		request   *http.Request
	}
)

func NewRespond(
	logger logger.Logger,
	validator validator.Validator,
) Responder {
	return Respond{
		logger:    logger,
		validator: validator,
	}
}

func NewRespondBuilder(w http.ResponseWriter, r *http.Request) Responder {
	return Respond{
		logger:    logger.GetLogger(),
		validator: validator.GetValidator(),
		writer:    w,
		request:   r,
	}
}

func (h Respond) Http(w http.ResponseWriter, r *http.Request) Responder {
	return Respond{
		logger:    h.logger,
		validator: h.validator,
		writer:    w,
		request:   r,
	}
}

func (h Respond) SetHeader(key string, value string) Responder {
	h.writer.Header().Set(key, value)
	return h
}

func (h Respond) WriteHeader(httpStatus int) {
	h.writer.WriteHeader(httpStatus)
}

func (h Respond) Json(httpStatus int, data any, meta ...any) {
	render.Status(h.request, httpStatus)
	render.JSON(h.writer, h.request, Response{
		Data: data,
		Meta: gog.ByDefault(nil, meta...),
	})
}

func (h Respond) Any(httpStatus int, data any, meta ...any) {
	render.Status(h.request, httpStatus)
	render.Respond(h.writer, h.request, Response{
		Data: data,
		Meta: gog.ByDefault(nil, meta...),
	})
}

func (h Respond) Redirect(url string, httpStatus ...int) {
	http.Redirect(h.writer, h.request, url, gog.ByDefault(http.StatusMovedPermanently, httpStatus...))
}

func (h Respond) Decode(v any) error {
	return render.Decode(h.request, v)
}

func (h Respond) DecodeValidate(v any) error {
	if err := h.Decode(v); err != nil {
		return err
	}

	return h.validator.Struct(v)
}

func (h Respond) Err(err error) {
	h.logger.Errorf(err.Error())

	h.SetHeader("Content-Type", "application/json; charset=utf-8")
	if err := render.Render(h.writer, h.request, NewErrorResponse(err)); err != nil {
		panic(err)
	}
}
