package Services

import (
	"Polybub/Data"
	"Polybub/Data/Audit"
	"Polybub/Data/Enums/UserGroups"
	"Polybub/Data/Models"
	"context"
)

type userVariant struct {
	Id           int32  `gorm:"column:Id;type:INTEGER;primaryKey;" json:"Id"`
	FirstName    string `gorm:"column:FirstName;type:TEXT" json:"FirstName"`
	LastName     string `gorm:"column:LastName;type:TEXT" json:"LastName"`
	Username     string `gorm:"column:Username;type:TEXT" json:"Username"`
	AccountEmail string `gorm:"column:AccountEmail;type:TEXT" json:"AccountEmail"`
	AccountPhone string `gorm:"column:AccountPhone;type:TEXT" json:"AccountPhone"`
	UserGroup    int32  `gorm:"column:UserGroup;type:INTEGER" json:"UserGroup"`
	Audit.AuditFields
}

func getUserVariant(user Models.User) userVariant {
	v := userVariant{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		AccountEmail: user.AccountEmail,
		AccountPhone: user.AccountPhone,
		UserGroup:    user.UserGroup,
	}

	return v
}

func CreateUser(ctx context.Context, data Models.User) (userVariant, error) {
	var db = Data.GetConnection().WithContext(ctx)

	data.Id = 0
	data.Password = ""
	data.Salt = ""
	data.UserGroup = UserGroups.InternalUsers // Remove to allow multiple UserGroups

	err := db.Model(&Models.User{}).
		Save(&data).
		Error
	if err != nil {
		return userVariant{}, err
	}

	return getUserVariant(data), nil
}

func ReadSingleUser(ctx context.Context, id int32) (userVariant, error) {
	var db = Data.GetConnection().WithContext(ctx)

	single := Models.User{}
	err := db.Model(&Models.User{}).
		Where("Id = ?", id).
		First(&single).
		Error
	if err != nil {
		return userVariant{}, err
	}

	return getUserVariant(single), nil
}

func GetIdByEmail(ctx context.Context, email string) (int32, error) {
	var db = Data.GetConnection().WithContext(ctx)

	single := Models.User{}
	err := db.Model(&Models.User{}).
		Where("AccountEmail = ?", email).
		First(&single).
		Error
	if err != nil {
		return 0, err
	}

	return single.Id, nil
}

func GetIdByUsername(ctx context.Context, username string) (int32, error) {
	var db = Data.GetConnection().WithContext(ctx)

	single := Models.User{}
	err := db.Model(&Models.User{}).
		Where("Username = ?", username).
		First(&single).
		Error
	if err != nil {
		return 0, err
	}

	return single.Id, nil
}

func ReadManyUser(ctx context.Context) ([]userVariant, error) {
	var db = Data.GetConnection().WithContext(ctx)

	many := []Models.User{}
	err := db.Find(&many).Error
	if err != nil {
		return []userVariant{}, err
	}

	var manyVariants = []userVariant{}
	for i := 0; i < len(many); i++ {
		manyVariants = append(manyVariants, getUserVariant(many[i]))
	}

	return manyVariants, nil
}

func UpdateUser(ctx context.Context, data Models.User) (userVariant, error) {
	var db = Data.GetConnection().WithContext(ctx)

	data.Password = ""
	data.Salt = ""

	single := Models.User{}
	err := db.Model(&Models.User{}).
		Where("Id = ?", data.Id).
		Updates(data).
		First(&single).
		Error
	if err != nil {
		return userVariant{}, err
	}

	return getUserVariant(single), nil
}

func SoftDeleteUser(ctx context.Context, id int32) error {
	var db = Data.GetConnection().WithContext(ctx)

	var data = &Models.User{
		Id: id,
	}

	err := db.Model(&Models.User{}).
		Where("Id = ?", data.Id).
		Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}

// Unsafe
func UnsafeReadSingleUser(ctx context.Context, id int32) (Models.User, error) {
	var db = Data.GetConnection().WithContext(ctx)

	single := Models.User{}
	err := db.Model(&Models.User{}).
		Where("Id = ?", id).
		First(&single).
		Error
	if err != nil {
		return Models.User{}, err
	}

	return single, nil
}

func UnsafeGetUserByUsername(ctx context.Context, username string) (Models.User, error) {
	var db = Data.GetConnection().WithContext(ctx)

	user := Models.User{}
	err := db.Model(&Models.User{}).
		Where("Username = ?", username).
		First(&user).
		Error
	if err != nil {
		return user, err
	}

	return user, nil
}
