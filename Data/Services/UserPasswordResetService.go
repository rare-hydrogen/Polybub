package Services

import (
	"Polybub/Data"
	"Polybub/Data/Models"
	"Polybub/Utilities/Smtp"
	"context"
	"crypto/rand"
)

func AddResetKeyThenDeleteOthers(ctx context.Context, userId int32, givenEmail string) error {
	var db = Data.GetConnection().WithContext(ctx)

	data := Models.UserPasswordReset{
		UserId:   userId,
		ResetKey: rand.Text(),
	}

	err := db.Model(&Models.UserPasswordReset{}).
		Save(&data).
		Error
	if err != nil {
		return err
	}

	err = db.Model(&Models.UserPasswordReset{}).
		Where("UserId = ?", userId).
		Where("ResetKey IS NOT ?", data.ResetKey).
		Delete(&Models.UserPasswordReset{}).
		Error
	if err != nil {
		return err
	}

	_, err = Smtp.SendgridEmail(data.UserId, data.ResetKey, givenEmail)
	if err != nil {
		return err
	}

	return nil
}

func GetResetKey(ctx context.Context, userId int32) (string, error) {
	var db = Data.GetConnection().WithContext(ctx)

	single := Models.UserPasswordReset{}
	err := db.Model(&Models.UserPasswordReset{}).
		Where("UserId = ?", userId).
		First(&single).
		Error
	if err != nil {
		return "", err
	}

	return single.ResetKey, nil
}

func DeleteAllResetKeys(ctx context.Context, userId int32) error {
	var db = Data.GetConnection().WithContext(ctx)

	err := db.Model(&Models.UserPasswordReset{}).
		Where("UserId = ?", userId).
		Delete(&Models.UserPasswordReset{}).
		Error
	if err != nil {
		return err
	}

	return nil
}
