package Services

import (
	"Polybub/Data"
	"Polybub/Data/Models"
)

func CreatePermission(data Models.Permission) (Models.Permission, error) {
	var db = Data.GetConnection()

	err := db.Model(&Models.Permission{}).
		Save(&data).
		Error
	if err != nil {
		return Models.Permission{}, err
	}

	return data, nil
}

func ReadSinglePermission(id int32) (Models.Permission, error) {
	var db = Data.GetConnection()

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

func ReadUsersPermissions(userId int32) ([]Models.Permission, error) {
	var db = Data.GetConnection()

	many := []Models.Permission{}
	err := db.
		Where("UserId = ?", userId).
		Find(&many).Error
	if err != nil {
		return []Models.Permission{}, err
	}

	return many, nil
}

func UpdatePermission(data Models.Permission) (Models.Permission, error) {
	var db = Data.GetConnection()

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

func SoftDeletePermission(id int32) error {
	var db = Data.GetConnection()

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
