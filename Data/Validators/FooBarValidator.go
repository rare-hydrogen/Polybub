package Validators

import (
	"Polybub/Data/Models"

	v "github.com/go-ozzo/ozzo-validation"
)

func FooBar(t Models.FooBar) error {
	return v.ValidateStruct(&t,
		v.Field(&t.Name, fooBarRules.NameRule...),
		v.Field(&t.Type, fooBarRules.TypeRule...),
		v.Field(&t.Amount, fooBarRules.AmountRule...),
	)
}

// Rules
type FooBarRules struct {
	NameRule   []v.Rule
	TypeRule   []v.Rule
	AmountRule []v.Rule
}

var fooBarRules = FooBarRules{
	NameRule: []v.Rule{
		v.Required,
		v.Length(1, 50),
	},
	TypeRule: []v.Rule{
		v.Required,
		v.Length(1, 50),
	},
	AmountRule: []v.Rule{
		v.Required,
		v.Min(1),
		v.Max(1000),
	},
}
