package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Validator struct
type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

// ValidationError represent error from validator
type ValidationError struct {
	Namespace string `json:"namespace,omitempty"`
	Field     string `json:"field,omitempty"`
	Error     string `json:"error,omitempty"`
}

// NewValidator create an instance of Validator Struct
func NewValidator() *Validator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	_validator := &Validator{
		validate: validator.New(),
		trans:    trans,
	}

	_validator.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	_validator.validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	return _validator
}

// FormatErrors generate list of error
func (v *Validator) FormatErrors(errs validator.ValidationErrors) []ValidationError {
	validationErrorMessages := errs.Translate(v.trans)

	messages := make([]ValidationError, 0)
	for _, err := range errs {
		messages = append(messages, ValidationError{
			Namespace: err.Namespace(),
			Field:     err.Field(),
			Error:     validationErrorMessages[err.Namespace()],
		})
	}

	return messages
}

// ValidateStruct func
func (v *Validator) ValidateStruct(input interface{}) []ValidationError {
	err := v.validate.Struct(input)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errResp := v.FormatErrors(validationErrors)
		return errResp
	}

	return []ValidationError{}
}
