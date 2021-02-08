package main

import (
	"fmt"
	"time"
)

const validUser string = "c137@onecause.com"
const validPass string = "#th@nH@rm#y#r!$100%D0p#"

type Content struct {
	// Response Struct
	Valid   bool   `json:"valid" form:"valid" query:"valid"`
	Message string `json:"message" form:"message" query:"message"`
}

type Login struct {
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"password" form:"password" query:"password"`
	Onetime  string `json:"onetime" form: "onetime" query:"onetime"`
}

func validate(test string, match string) (isValid bool) {
	fmt.Println("testing: %s \r\n against: %s", test, match)
	isValid = test == match
	return isValid
}

func generateOneTimeCompare() string {
	t := time.Now()
	fmtT := t.Format("15:04")
	compare := fmtT[0:2] + "" + fmtT[3:5]
	return compare
}

func checkLogin(name string, pass string, onetime string) bool {
	checkUser := validate(name, validUser)
	checkPass := validate(pass, validPass)
	checkOnetime := validate(onetime, generateOneTimeCompare())
	return checkUser && checkPass && checkOnetime
}
