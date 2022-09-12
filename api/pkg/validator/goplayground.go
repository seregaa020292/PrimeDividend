package validator

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"primedivident/pkg/errorn"
	"primedivident/pkg/utils"

	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	goPlayground "github.com/go-playground/validator/v10"
	ruTranslations "github.com/go-playground/validator/v10/translations/ru"
	"github.com/google/uuid"
)

type playgroundValidator struct {
	validate   *goPlayground.Validate
	translator ut.Translator
}

func NewGoPlayground() Validator {
	var (
		language          = ru.New()
		uni               = ut.New(language, language)
		translator, found = uni.GetTranslator("ru")
	)

	if !found {
		log.Fatalln("translator not found")
	}

	validate := goPlayground.New()
	if err := ruTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
		log.Fatalln("translator not found")
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		return utils.If(name == "-", "", name)
	})

	validate.RegisterCustomTypeFunc(func(field reflect.Value) any {
		if valuer, ok := field.Interface().(uuid.UUID); ok {
			return valuer.String()
		}
		return nil
	}, uuid.UUID{})

	log.Println("Start Validator")

	return playgroundValidator{
		validate:   validate,
		translator: translator,
	}
}

func (v playgroundValidator) Struct(i any) error {
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}

	fieldErrors, ok := err.(goPlayground.ValidationErrors)
	if !ok {
		return err
	}

	return v.messages(fieldErrors)
}

func (v playgroundValidator) Var(field any, tag string) error {
	err := v.validate.Var(field, tag)
	if err == nil {
		return nil
	}

	fieldErrors, ok := err.(goPlayground.ValidationErrors)
	if !ok {
		return err
	}

	return errorn.IncorrectInput(errorn.Message{
		Error: fmt.Errorf(
			"%s%s",
			fieldErrors[0].Value(),
			fieldErrors[0].Translate(v.translator),
		),
	})
}

func (v playgroundValidator) messages(fieldErrors goPlayground.ValidationErrors) error {
	var messages []errorn.Message

	for _, fieldErr := range fieldErrors {
		messages = append(messages, errorn.Message{
			Error: fmt.Errorf("%s", fieldErr.Translate(v.translator)),
			Field: fieldErr.Field(),
		})
	}

	return errorn.IncorrectInput(messages...)
}
