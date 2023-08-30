package utils

import (
	"testing"

	"github.com/shobky/giggs/api/schema"
)

func TestCheckValidator(t *testing.T) {
	b := &schema.SignupBody{
		Email:      "shobkyya@gmail.com",
		Password:   "passwdordAA!!",
		GivenName:  "ahmed",
		FamilyName: "shobky",
		Name:       "ahmed shobky",
	}

	err := Validate(b)

	if err != nil {
		t.Fatal(err)
	}
}
