package utils_validator

import (
	"reflect"
	"regexp"
	"time"

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

func IsDateAfterNow(fl validator.FieldLevel) bool {

	if fl.Field().Type() != reflect.TypeOf(time.Time{}) {
		println("Expected time.Time type, got:", fl.Field().Type().String())
		return false
	}
	// Tenta fazer o parse da string no formato RFC3339
	inputTime := fl.Field().Interface().(time.Time)
	// Compara com o horário atual
	if inputTime.After(time.Now()) {
		return true
	}

	return false
}
