package Models

import (
	"Polybub/Data/Audit"
	"encoding/json"
	"strconv"
	"strings"

	money "github.com/Rhymond/go-money"
)

const TableNameFooBar = "FooBar"

type FooBar struct {
	Id       int32  `gorm:"column:Id;type:INTEGER;primaryKey;" json:"Id"`
	Name     string `gorm:"column:Name;type:TEXT" json:"Name"`
	Type     string `gorm:"column:Type;type:TEXT" json:"Type"`
	Amount   int64  `gorm:"column:Amount;type:INTEGER;" json:"amount"`
	Currency string `gorm:"column:Currency;type:TEXT;" json:"currency"`
	Audit.AuditFields
}

func (*FooBar) TableName() string {
	return TableNameFooBar
}

func (foobar *FooBar) UnmarshalJSON(data []byte) error {
	var temp struct {
		Id       int32  `json:"Id"`
		Name     string `json:"Name"`
		Type     string `json:"Type"`
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if !strings.Contains(temp.Amount, ".") {
		temp.Amount = temp.Amount + ".00"
	}

	f, err := strconv.ParseFloat(temp.Amount, 64)
	if err != nil {
		return err
	}

	// TODO: Handle other currencies
	c := "USD"
	m := money.NewFromFloat(f, c)

	foobar.Id = temp.Id
	foobar.Name = temp.Name
	foobar.Type = temp.Type
	foobar.Amount = m.Amount()
	foobar.Currency = c
	return nil
}
