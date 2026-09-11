package Models

import "Polybub/Data/Audit"

const TableNameUserPasswordReset = "user_password_resets"

type UserPasswordReset struct {
	Id       int32  `gorm:"column:id;type:INTEGER;primaryKey;" json:"Id"`
	UserId   int32  `gorm:"column:user_id;type:INTEGER" json:"UserId,string"`
	ResetKey string `gorm:"column:reset_key;type:TEXT" json:"ResetKey"`
	Audit.AuditFields
}

func (*UserPasswordReset) TableName() string {
	return TableNameUserPasswordReset
}
