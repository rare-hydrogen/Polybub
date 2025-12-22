package Services

import (
	"Polybub/Data"
	"Polybub/Data/Models"
	"Polybub/Utilities"
	"Polybub/Utilities/Smtp"
	"crypto/rand"
	"strconv"
)

func AddResetKeyThenDeleteOthers(userId int32, givenEmail string) error {
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

	baseUrl := Utilities.GetBaseUrl(Utilities.GlobalConfig)
	to := "To: " + givenEmail + "\r\n"
	subject := "Subject: Password Reset - " + baseUrl + "\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	resetUrl := baseUrl + "/forgot-password?id=" + strconv.Itoa(int(data.UserId)) + "&key=" + data.ResetKey
	body := "To reset your password, click <a href=\"" + resetUrl + "\">This Link</a>" + "\r\n"
	message := []byte(to + subject + mime + body)

	err = Smtp.SendEmail([]string{givenEmail}, []byte(message))
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
