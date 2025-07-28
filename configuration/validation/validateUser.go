package validation

import (
	"encoding/json"
	"errors"
	errorhandler "go-jwt/errorHandler"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *errorhandler.ErrorHandler {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return errorhandler.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorsCauses := []errorhandler.Causes{}

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := errorhandler.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return errorhandler.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return errorhandler.NewBadRequestError("Error trying to convert fields")
	}
}
