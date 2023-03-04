package utils

import (
	"errors"
	"regexp"

	zh_CN "github.com/go-playground/locales/zh"
	uniTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

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

// using validator library to validate a Struct
// func IsValidStruct(val interface{}) error {
// 	myValidator := validator.New()
// 	return myValidator.Struct(val)
// }

// using validator library to validate a Struct
func IsValidStruct(dataStruct interface{}) error {
	zh_ch := zh_CN.New()
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	// validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
	// 	name := fld.Tag.Get("label")
	// 	return name
	// })

	uni := uniTranslator.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器注册翻译器
	zhTranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return err
}
