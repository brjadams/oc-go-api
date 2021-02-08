package main

const validUser string = "c137@onecause.com"
const validPass string = "#th@nH@rm#y#r!$100%D0p#"

type Content struct {
	// Response Struct
	Valid   bool   `json: "valid"`
	Message string `json: "message"`
}

func validate(test string, match string) (isValid bool) {
	isValid = test == match
	return isValid
}

func compareOneTime(code int) bool {
	return true
}

func checkLogin(user string, pass string) bool {
	checkUser := validate(user, validUser)
	checkPass := validate(pass, validPass)
	return checkUser && checkPass && compareOneTime(0111)
}
