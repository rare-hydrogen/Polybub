package Services

import (
	"Polybub/Data"
	"Polybub/Data/Models"
	"crypto/rand"
)

func AddResetKeyThenDeleteOthers(userId int32) error {
	var db = Data.GetConnection()

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

	return nil
}

func GetResetKey(userId int32) (string, error) {
	var db = Data.GetConnection()

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

func DeleteAllResetKeys(userId int32) error {
	var db = Data.GetConnection()

	err := db.Model(&Models.UserPasswordReset{}).
		Where("UserId = ?", userId).
		Delete(&Models.UserPasswordReset{}).
		Error
	if err != nil {
		return err
	}

	return nil
}
