package response

import (
	"net/http"

	"github.com/go-chi/render"

	"primedivident/pkg/errorn"
	"primedivident/pkg/logger"
	"primedivident/pkg/utils"
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

func (h Respond) Redirect(url string, httpStatus ...int) {
	http.Redirect(h.w, h.r, url, utils.ByDefault(http.StatusMovedPermanently, httpStatus...))
}

func (h Respond) Decode(v interface{}) error {
	return render.Decode(h.r, v)
}

func (h Respond) Err(err error) {
	h.SetHeader("Content-Type", "application/json; charset=utf-8")
	if err := render.Render(h.w, h.r, ErrRender(err)); err != nil {
		panic(err)
	}
}

func ErrRender(err error) ErrorRespond {
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

	return errorRespond
}
