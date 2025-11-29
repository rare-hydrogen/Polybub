package Validators

import (
	"Polybub/Data/Enums/UserGroups"
	"Polybub/Data/Models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"regexp"

	v "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
)

type PasswordDto struct {
	Password      string `json:"Password"`
	PasswordAgain string `json:"PasswordAgain"`
}

func User(t Models.User) error {
	return v.ValidateStruct(&t,
		v.Field(&t.FirstName, userRules.NameRule...),
		v.Field(&t.LastName, userRules.NameRule...),
		v.Field(&t.UserGroup, userRules.UserGroupRule...),
		v.Field(&t.Username, userRules.UsernameRule...),
		v.Field(&t.Password, userRules.PasswordRule...),
		v.Field(&t.AccountEmail, userRules.EmailRule...),
		v.Field(&t.AccountPhone, userRules.PhoneRule...),
	)
}

func UserPassword(bodyBytes []byte) error {
	passwordDto := PasswordDto{}

	d := json.NewDecoder(io.NopCloser(bytes.NewBuffer(bodyBytes)))
	err := d.Decode(&passwordDto)
	if err != nil {
		return errors.New("decoding failed")
	}

	if passwordDto.Password != passwordDto.PasswordAgain {
		return errors.New("match failed")
	}

	return nil
}

// Rules
type UserRules struct {
	NameRule      []v.Rule
	UserGroupRule []v.Rule
	UsernameRule  []v.Rule
	PasswordRule  []v.Rule
	EmailRule     []v.Rule
	PhoneRule     []v.Rule
}

var userRules = UserRules{
	NameRule: []v.Rule{
		v.Required,
		v.Length(1, 50),
	},
	UserGroupRule: []v.Rule{
		v.Required,
		v.In(UserGroups.InternalUsers, UserGroups.ExternalUsers).Error("missing user group"),
	},
	UsernameRule: []v.Rule{
		v.Required,
		v.Length(10, 50),
	},
	PasswordRule: []v.Rule{
		v.Length(10, 50),
		v.Match(regexp.MustCompile(`[a-z]`)).Error("must have a lowercase letter"),
		v.Match(regexp.MustCompile(`[A-Z]`)).Error("must have an uppercase letter"),
		v.Match(regexp.MustCompile(`[0-9]`)).Error("must have a digit"),
		v.Match(regexp.MustCompile(`[\W_]`)).Error("must have a special character"),
	},
	EmailRule: []v.Rule{
		v.Required,
		v.Length(10, 50),
		is.Email,
	},
	PhoneRule: []v.Rule{
		v.Required,
		v.Match(regexp.MustCompile(`^\+\d{11,15}$`)).Error("must match format +12223334444"),
	},
}
