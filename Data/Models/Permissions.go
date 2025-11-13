package Models

import "Polybub/Data/Audit"

const TableNamePermission = "Permissions"

type Permission struct {
	Id       int32  `gorm:"column:Id;type:INTEGER;primaryKey;" json:"Id"`
	UserId   int32  `gorm:"column:UserId;type:INTEGER" json:"UserId"`
	Name     string `gorm:"column:Name;type:TEXT" json:"Name"`
	IsCreate bool   `gorm:"column:IsCreate;type:INTEGER" json:"IsCreate"`
	IsRead   bool   `gorm:"column:IsRead;type:INTEGER" json:"IsRead"`
	IsUpdate bool   `gorm:"column:IsUpdate;type:INTEGER" json:"IsUpdate"`
	IsDelete bool   `gorm:"column:IsDelete;type:INTEGER" json:"IsDelete"`
	Audit.AuditFields
}

func (*Permission) TableName() string {
	return TableNamePermission
}
