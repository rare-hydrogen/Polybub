package Models

import "Polybub/Data/Audit"

const TableNamePermission = "permissions"

type Permission struct {
	Id       int32  `gorm:"column:id;type:INTEGER;primaryKey;" json:"Id"`
	UserId   int32  `gorm:"column:user_id;type:INTEGER" json:"UserId"`
	Name     string `gorm:"column:name;type:TEXT" json:"Name"`
	IsCreate bool   `gorm:"column:is_create;type:INTEGER" json:"IsCreate"`
	IsRead   bool   `gorm:"column:is_read;type:INTEGER" json:"IsRead"`
	IsUpdate bool   `gorm:"column:is_update;type:INTEGER" json:"IsUpdate"`
	IsDelete bool   `gorm:"column:is_delete;type:INTEGER" json:"IsDelete"`
	Audit.AuditFields
}

func (*Permission) TableName() string {
	return TableNamePermission
}
