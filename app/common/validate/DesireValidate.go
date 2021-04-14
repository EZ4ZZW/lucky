package validate

import "lucky/app/helper"

var DesireValidate helper.Validator

func init() {
	rules := map[string]string{
		"user_id":        "required",
		"idcard_number":  "required",
		"nick":           "required",
		"password":       "required",
		"school":         "required",
		"major":          "required",
		"contact":        "required",
		"mail":           "required|email",
	}

	scenes := map[string][]string{
		"login": {"idcard_number", "password"},
	}
	DesireValidate.Rules = rules
	DesireValidate.Scenes = scenes
}