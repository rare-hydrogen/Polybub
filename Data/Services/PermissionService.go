package Services

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data"
	"Polybub/Data/Models"
	"context"
)

func CreatePermission(ctx context.Context, data Models.Permission) (Models.Permission, error) {
	var db = Data.GetConnection().WithContext(ctx)

	err := db.Model(&Models.Permission{}).
		Save(&data).
		Error
	if err != nil {
		return Models.Permission{}, err
	}

	return data, nil
}

func CreateDefaultPermissions(ctx context.Context, userId int32) error {
	defaultPermissions := []Models.Permission{
		OAuth2.NewPerm("MfaCode", true, true, true, false),
		OAuth2.NewPerm("Dashboard", true, false, false, false),
		OAuth2.NewPerm("FooBars", true, true, true, true),
	}

	for i := 0; i < len(defaultPermissions); i++ {
		defaultPermissions[i].UserId = userId
	}

	var db = Data.GetConnection().WithContext(ctx)

	err := db.Model(&[]Models.Permission{}).
		Save(defaultPermissions).
		Error
	if err != nil {
		return err
	}

	return nil
}

func ReadSinglePermission(ctx context.Context, id int32) (Models.Permission, error) {
	var db = Data.GetConnection().WithContext(ctx)

	single := Models.Permission{}
	err := db.Model(&Models.Permission{}).
		Where("Id = ?", id).
		First(&single).
		Error
	if err != nil {
		return Models.Permission{}, err
	}

	return single, nil
}

func ReadUsersPermissions(ctx context.Context, userId int32) ([]Models.Permission, error) {
	var db = Data.GetConnection().WithContext(ctx)

	many := []Models.Permission{}
	err := db.
		Where("UserId = ?", userId).
		Find(&many).Error
	if err != nil {
		return []Models.Permission{}, err
	}

	return many, nil
}

func UpdatePermission(ctx context.Context, data Models.Permission) (Models.Permission, error) {
	var db = Data.GetConnection().WithContext(ctx)

	single := Models.Permission{}
	err := db.Model(&Models.Permission{}).
		Where("Id = ?", data.Id).
		Updates(data).
		First(&single).
		Error
	if err != nil {
		return Models.Permission{}, err
	}

	return single, nil
}

func SoftDeletePermission(ctx context.Context, id int32) error {
	var db = Data.GetConnection().WithContext(ctx)

	var data = &Models.Permission{
		Id: id,
	}

	err := db.Model(&Models.Permission{}).
		Where("Id = ?", data.Id).
		Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}
