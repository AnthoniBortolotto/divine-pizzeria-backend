package utils_validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsName(fl validator.FieldLevel) bool {
	regex := `^([A-Za-zÀ-ÿ]+)(\s[A-Za-zÀ-ÿ]+)+$`

	nameRegex := regexp.MustCompile(regex)
	return nameRegex.MatchString(fl.Field().String())
}

func IsCep(fl validator.FieldLevel) bool {
	regex := `^\d{5}\d{3}$`

	cepRegex := regexp.MustCompile(regex)
	return cepRegex.MatchString(fl.Field().String())
}
