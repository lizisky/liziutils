package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// using validator library to validate a Struct
func IsValidStruct(val interface{}) error {
	myValidator := validator.New()
	return myValidator.Struct(val)
}

func IsValidEmailAddr(email string) bool {
	ret, _ := regexp.MatchString("^.+@.+$", email)
	return ret
}

func IsDigit(str string) bool {
	ret, _ := regexp.MatchString("^[0-9]*$", str)
	return ret
}

func IsPositiveNumber(src string) bool {
	var number = regexp.MustCompile("^[\\d]*$")
	return number.MatchString(src)
}
