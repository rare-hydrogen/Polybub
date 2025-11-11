package Models

import "Polybub/Data/Audit"

const TableNameUserPasswordReset = "UserPasswordResets"

type UserPasswordReset struct {
	Id       int32  `gorm:"column:Id;type:INTEGER;primaryKey;" json:"Id"`
	UserId   int32  `gorm:"column:UserId;type:INTEGER" json:"UserId,string"`
	ResetKey string `gorm:"column:ResetKey;type:TEXT" json:"ResetKey"`
	Audit.AuditFields
}

func (*UserPasswordReset) TableName() string {
	return TableNameUserPasswordReset
}
