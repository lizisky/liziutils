package utils

import "regexp"

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
