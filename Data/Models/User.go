package Models

import "Polybub/Data/Audit"

const TableNameUser = "users"

type User struct {
	Id           int32  `gorm:"column:id;type:INTEGER;primaryKey;" json:"Id"`
	FirstName    string `gorm:"column:first_name;type:TEXT" json:"FirstName"`
	LastName     string `gorm:"column:last_name;type:TEXT" json:"LastName"`
	Username     string `gorm:"column:username;type:TEXT" json:"Username"`
	Password     string `gorm:"column:password;type:TEXT" json:"Password"`
	Salt         string `gorm:"column:salt;type:TEXT" json:"Salt"`
	AccountEmail string `gorm:"column:account_email;type:TEXT" json:"AccountEmail"`
	AccountPhone string `gorm:"column:account_phone;type:TEXT" json:"AccountPhone"`
	UserGroup    int32  `gorm:"column:user_group;type:INTEGER" json:"UserGroup,string"`
	TotpKey      string `gorm:"column:totp_key;type:TEXT" json:"TotpKey"`
	Audit.AuditFields
}

func (*User) TableName() string {
	return TableNameUser
}
