package Models

import (
	"Polybub/Data/Audit"

	money "github.com/Rhymond/go-money"
)

const TableNameFooBar = "FooBar"

type FooBar struct {
	Id       int32          `gorm:"column:Id;type:INTEGER;primaryKey;" json:"Id"`
	Name     string         `gorm:"column:Name;type:TEXT" json:"Name"`
	Type     string         `gorm:"column:Type;type:TEXT" json:"Type"`
	Amount   int64          `gorm:"column:Amount;type:INTEGER;" json:"Amount"`
	Currency money.Currency `gorm:"column:Currency;type:TEXT;" json:"Currency"`
	Audit.AuditFields
}

func (*FooBar) TableName() string {
	return TableNameFooBar
}
