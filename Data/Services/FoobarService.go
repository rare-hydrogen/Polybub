package Services

import (
	"Polybub/Data"
	"Polybub/Data/Models"
)

func ReadSingleFooBar(Id int32) (Models.FooBar, error) {
	var db = Data.GetConnection()

	Single := Models.FooBar{}
	err := db.Model(&Models.FooBar{}).
		Where("Id = ?", Id).
		First(&Single).
		Error
	if err != nil {
		return Models.FooBar{}, err
	}

	return Single, nil
}

func ReadManyFooBar() ([]Models.FooBar, error) {
	var db = Data.GetConnection()

	Many := []Models.FooBar{}
	err := db.Find(&Many).Error
	if err != nil {
		return []Models.FooBar{}, err
	}

	return Many, nil
}

func CreateFooBar(data Models.FooBar) (Models.FooBar, error) {
	var db = Data.GetConnection()

	err := db.Model(&Models.FooBar{}).
		Where("Id = ?", data.Id).
		Save(&data).
		Error
	if err != nil {
		return Models.FooBar{}, err
	}

	return data, nil
}

func UpdateFooBar(data Models.FooBar) (Models.FooBar, error) {
	var db = Data.GetConnection()

	Single := Models.FooBar{}
	err := db.Model(&Models.FooBar{}).
		Where("Id = ?", data.Id).
		Updates(data).
		First(&Single).
		Error
	if err != nil {
		return Models.FooBar{}, err
	}

	return Single, nil
}

func SoftDeleteFooBar(Id int32) error {
	var db = Data.GetConnection()

	var data = &Models.FooBar{
		Id: Id,
	}

	err := db.Model(&Models.FooBar{}).
		Where("Id = ?", data.Id).
		Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}
