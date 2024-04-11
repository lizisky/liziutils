package main

import (
	"log"

	"github.com/lizisky/liziutils/utils"
)

func test_validator() {
	// validate := validator.New()

	s1 := MyStruct{
		ID:    4,
		Code:  "blaklasdfsfsfsasj",
		EMail: "a@alkdj.cn",
		// Category: "dß",
	}
	// err := validate.Struct(s1)
	err := utils.IsValidStruct(s1)
	log.Printf("111 %+v", err)

	// s2 := MyStruct{}
	// err = validate.Struct(s2)
	// log.Printf("222 %+v", err)
}

type MyStruct struct {
	// gorm.Model
	ID       int    `validate:"gt=1"`
	Code     string `validate:"required,gt=5,lte=10"` // gt/gte/lt/lte 验证字符串长度在一个指定范围内
	EMail    string `validate:"omitempty,email"`
	Category string `validate:"omitempty,gt=5,lte=10"` // gt/gte/lt/lte 验证字符串长度在一个指定范围内
	Saled    bool
}
