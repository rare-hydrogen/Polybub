package Models

import "Polybub/Data/Audit"

const TableNameUser = "Users"

type User struct {
	Id           int32  `gorm:"column:Id;type:INTEGER;primaryKey;" json:"Id"`
	FirstName    string `gorm:"column:FirstName;type:TEXT" json:"FirstName"`
	LastName     string `gorm:"column:LastName;type:TEXT" json:"LastName"`
	Username     string `gorm:"column:Username;type:TEXT" json:"Username"`
	Password     string `gorm:"column:Password;type:TEXT" json:"Password"`
	Salt         string `gorm:"column:Salt;type:TEXT" json:"Salt"`
	AccountEmail string `gorm:"column:AccountEmail;type:TEXT" json:"AccountEmail"`
	AccountPhone string `gorm:"column:AccountPhone;type:TEXT" json:"AccountPhone"`
	UserGroup    int32  `gorm:"column:UserGroup;type:INTEGER" json:"UserGroup"`
	Audit.AuditFields
}

func (*User) TableName() string {
	return TableNameUser
}
