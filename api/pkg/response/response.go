package response

import (
	"net/http"

	"github.com/go-chi/render"

	"primedivident/pkg/logger"
	"primedivident/pkg/utils"
	"primedivident/pkg/validator"
)

type Responder interface {
	Http(w http.ResponseWriter, r *http.Request) Responder
	SetHeader(key string, value string) Responder
	WriteHeader(httpStatus int)
	Json(httpStatus int, data any)
	Any(httpStatus int, data any)
	Err(err error)
	Redirect(url string, httpStatus ...int)
	Decode(v any) error
	DecodeValidate(v any) error
}

type (
	Response struct {
		Data any `json:"data"`
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

func (h Respond) Json(httpStatus int, data any) {
	render.Status(h.request, httpStatus)
	render.JSON(h.writer, h.request, Response{Data: data})
}

func (h Respond) Any(httpStatus int, data any) {
	render.Status(h.request, httpStatus)
	render.Respond(h.writer, h.request, Response{Data: data})
}

func (h Respond) Redirect(url string, httpStatus ...int) {
	http.Redirect(h.writer, h.request, url, utils.ByDefault(http.StatusMovedPermanently, httpStatus...))
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
	errorResponse := NewByError(err)

	h.logger.Errorf("%s", errorResponse.Error)

	h.SetHeader("Content-Type", "application/json; charset=utf-8")
	if err := render.Render(h.writer, h.request, errorResponse); err != nil {
		panic(err)
	}
}
