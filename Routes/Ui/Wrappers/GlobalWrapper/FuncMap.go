package GlobalWrapper

import (
	"html/template"

	"github.com/Rhymond/go-money"
)

var FuncMap = template.FuncMap{
	"moneyDisplay": func(amount int64, currency string) string {
		return money.New(amount, currency).Display()
	},
}
