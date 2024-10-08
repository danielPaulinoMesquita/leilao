package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"leilao/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTranslel := ut.New(en, en)
		transl, _ = enTranslel.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, transl)
	}
}

func ValidateErr(validator_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validator_err, &jsonErr) {
		return rest_err.NewNotFoundError("Invalid type error")
	} else if errors.As(validator_err, &jsonValidation) {
		errorCauses := []rest_err.Causes{}

		for _, cause := range validator_err.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Field:   cause.Field(),
				Message: cause.Translate(transl),
			})
		}

		return rest_err.NewBadRequestError("Invalid validation error", errorCauses...)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
