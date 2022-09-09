package response

import (
	"net/http"

	"github.com/go-chi/render"

	"primedivident/pkg/errorn"
	"primedivident/pkg/logger"
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

func (h Respond) Decode(v interface{}) error {
	return render.Decode(h.r, v)
}

func (h Respond) Err(err error) {
	var errorRespond ErrorRespond

	e, ok := err.(errorn.Errorn)
	if ok {
		errorRespond = FindErrorType(e)
	} else {
		errorRespond = InternalError(err)
	}

	logger.GetLogger().
		ExtraError(err).
		Errorf("%+v", errorRespond.Errors)

	h.SetHeader("Content-Type", "application/json; charset=utf-8")
	if err := render.Render(h.w, h.r, errorRespond); err != nil {
		panic(err)
	}
}
