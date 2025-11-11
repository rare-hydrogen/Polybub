package Services

import (
	"Polybub/Data"
	"Polybub/Data/Models"
)

func CreateFooBar(data Models.FooBar) (Models.FooBar, error) {
	var db = Data.GetConnection()

	err := db.Model(&Models.FooBar{}).
		Save(&data).
		Error
	if err != nil {
		return Models.FooBar{}, err
	}

	return data, nil
}

func ReadSingleFooBar(id int32) (Models.FooBar, error) {
	var db = Data.GetConnection()

	single := Models.FooBar{}
	err := db.Model(&Models.FooBar{}).
		Where("Id = ?", id).
		First(&single).
		Error
	if err != nil {
		return Models.FooBar{}, err
	}

	return single, nil
}

func ReadManyFooBar() ([]Models.FooBar, error) {
	var db = Data.GetConnection()

	many := []Models.FooBar{}
	err := db.Find(&many).Error
	if err != nil {
		return []Models.FooBar{}, err
	}

	return many, nil
}

func UpdateFooBar(data Models.FooBar) (Models.FooBar, error) {
	var db = Data.GetConnection()

	single := Models.FooBar{}
	err := db.Model(&Models.FooBar{}).
		Where("Id = ?", data.Id).
		Updates(data).
		First(&single).
		Error
	if err != nil {
		return Models.FooBar{}, err
	}

	return single, nil
}

func SoftDeleteFooBar(id int32) error {
	var db = Data.GetConnection()

	var data = &Models.FooBar{
		Id: id,
	}

	err := db.Model(&Models.FooBar{}).
		Where("Id = ?", data.Id).
		Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}
