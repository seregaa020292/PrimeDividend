package response

import (
	"net/http"

	"github.com/go-chi/render"

	"primedivident/pkg/errorn"
	"primedivident/pkg/logger"
	"primedivident/pkg/utils"
	"primedivident/pkg/validator"
)

type Responder interface {
	Http(w http.ResponseWriter, r *http.Request) Responder
	SetHeader(key string, value string) Responder
	Json(httpStatus int, data any)
	Any(httpStatus int, data any)
	NoContent()
	Redirect(url string, httpStatus ...int)
	Decode(v any) error
	DecodeValidate(v any) error
	Err(err error)
}

type (
	Response struct {
		Data any `json:"data"`
	}
	Respond struct {
		logger    logger.Logger
		validator validator.Validator

		w http.ResponseWriter
		r *http.Request
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
	respond := NewRespond(logger.GetLogger(), validator.GetValidator())
	return respond.Http(w, r)
}

func (h Respond) Http(w http.ResponseWriter, r *http.Request) Responder {
	return Respond{
		logger:    h.logger,
		validator: h.validator,

		w: w,
		r: r,
	}
}

func (h Respond) SetHeader(key string, value string) Responder {
	h.w.Header().Set(key, value)
	return h
}

func (h Respond) Json(httpStatus int, data any) {
	render.Status(h.r, httpStatus)
	render.JSON(h.w, h.r, Response{Data: data})
}

func (h Respond) Any(httpStatus int, data any) {
	render.Status(h.r, httpStatus)
	render.Respond(h.w, h.r, Response{Data: data})
}

func (h Respond) NoContent() {
	render.NoContent(h.w, h.r)
}

func (h Respond) Redirect(url string, httpStatus ...int) {
	http.Redirect(h.w, h.r, url, utils.ByDefault(http.StatusMovedPermanently, httpStatus...))
}

func (h Respond) Decode(v any) error {
	return render.Decode(h.r, v)
}

func (h Respond) DecodeValidate(v any) error {
	if err := render.Decode(h.r, v); err != nil {
		return err
	}

	return h.validator.Struct(v)
}

func (h Respond) Err(err error) {
	var errorRespond ErrorRespond

	e, ok := err.(errorn.Errorn)
	if ok {
		errorRespond = FindErrorType(e)
	} else {
		errorRespond = InternalError(err)
	}

	h.logger.ExtraError(err).Errorf("%+v", errorRespond.Errors)

	h.SetHeader("Content-Type", "application/json; charset=utf-8")
	if err := render.Render(h.w, h.r, errorRespond); err != nil {
		panic(err)
	}
}
