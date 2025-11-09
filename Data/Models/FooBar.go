package Models

import "Polybub/Data/Audit"

const TableNameFooBar = "FooBar"

type FooBar struct {
	Id   int32  `gorm:"column:Id;type:INTEGER;primaryKey;" json:"Id"`
	Name string `gorm:"column:Name;type:TEXT" json:"Name"`
	Type string `gorm:"column:Type;type:TEXT" json:"Type"`
	Audit.AuditFields
}

func (*FooBar) TableName() string {
	return TableNameFooBar
}
