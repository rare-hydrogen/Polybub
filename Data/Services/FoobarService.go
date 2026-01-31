package Services

import (
	"Polybub/Data"
	"Polybub/Data/Models"
	"context"
)

func CreateFooBar(ctx context.Context, data Models.FooBar) (Models.FooBar, error) {
	var db = Data.GetConnection().WithContext(ctx)

	err := db.Model(&Models.FooBar{}).
		Save(&data).
		Error
	if err != nil {
		return Models.FooBar{}, err
	}

	return data, nil
}

func ReadSingleFooBar(ctx context.Context, id int32) (Models.FooBar, error) {
	var db = Data.GetConnection().WithContext(ctx)

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

func ReadManyFooBar(ctx context.Context) ([]Models.FooBar, error) {
	var db = Data.GetConnection().WithContext(ctx)

	many := []Models.FooBar{}
	err := db.Find(&many).Error
	if err != nil {
		return []Models.FooBar{}, err
	}

	return many, nil
}

func ReadLatestFooBar(ctx context.Context) (Models.FooBar, error) {
	var db = Data.GetConnection().WithContext(ctx)

	single := Models.FooBar{}
	err := db.Model(&Models.FooBar{}).
		Last(&single).
		Error
	if err != nil {
		return Models.FooBar{}, err
	}

	return single, nil
}

func UpdateFooBar(ctx context.Context, data Models.FooBar) (Models.FooBar, error) {
	var db = Data.GetConnection().WithContext(ctx)

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

func SoftDeleteFooBar(ctx context.Context, id int32) error {
	var db = Data.GetConnection().WithContext(ctx)

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
