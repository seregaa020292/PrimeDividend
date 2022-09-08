package errorn

var (
	ErrorTypeUnknown        = Type{"unknown"}
	ErrorTypeAuthorization  = Type{"authorization"}
	ErrorTypeForbidden      = Type{"forbidden"}
	ErrorTypeNotFound       = Type{"not-found"}
	ErrorTypeIncorrectInput = Type{"incorrect-input"}
)

type Errorn struct {
	errorType Type
	messages  []Message
}

type Type struct {
	t string
}

type Message struct {
	Error error
	Field string
}

func New(errorType Type, messages ...Message) error {
	return Errorn{
		errorType: errorType,
		messages:  messages,
	}
}

func (e Errorn) Error() string {
	return e.errorType.t
}

func (e Errorn) ErrorType() Type {
	return e.errorType
}

func (e Errorn) Messages() []Message {
	return e.messages
}

func Unknown(messages ...Message) error {
	return New(ErrorTypeUnknown, messages...)
}

func Authorization(messages ...Message) error {
	return New(ErrorTypeAuthorization, messages...)
}

func Forbidden(messages ...Message) error {
	return New(ErrorTypeForbidden, messages...)
}

func NotFound(messages ...Message) error {
	return New(ErrorTypeNotFound, messages...)
}

func IncorrectInput(messages ...Message) error {
	return New(ErrorTypeIncorrectInput, messages...)
}
