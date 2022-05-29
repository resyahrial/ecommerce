package gvalidator

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ValidationFn func(fl validator.FieldLevel) bool

var trans ut.Translator
var Validator *validator.Validate

func init() {
	Validator = validator.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, _ = uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(Validator, trans)
}

func Validate(data interface{}) (string, bool) {
	err := Validator.Struct(data)
	if err == nil {
		return "", true
	}

	errs := translateError(err, trans)
	return strings.Join(errs, ", "), false
}

func translateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr.Error())
	}
	return errs
}
