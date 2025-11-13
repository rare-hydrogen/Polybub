package Services

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data"
	"Polybub/Data/Models"
	"errors"

	"github.com/google/uuid"
)

func getUserByUsername(username string) (Models.User, error) {
	var db = Data.GetConnection()

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

func Login(username string, password string) (string, error) {
	user, err := getUserByUsername(username)
	if err != nil {
		return "", err
	}

	if user.Password == "" {
		return "", errors.New("no password")
	}

	encryptedPassword := OAuth2.EncryptPassword(password, user.Salt)

	if encryptedPassword != user.Password {
		return "", errors.New("incorrect password")
	}

	permissions, err := ReadUsersPermissions(user.Id)
	if err != nil {
		return "", err
	}

	name := user.FirstName + " " + user.LastName
	jwtString, err := OAuth2.NewJwt(name, user.Id, user.UserGroup, permissions)
	if err != nil {
		return "", err
	}

	return jwtString, nil
}

func UpdatePasswordAndSalt(userId int32, password string) error {
	var db = Data.GetConnection()
	var salt = uuid.New().String()
	var user = Models.User{
		Id:       userId,
		Password: OAuth2.EncryptPassword(password, salt),
		Salt:     salt,
	}

	err := db.Model(&Models.User{}).
		Where("Id = ?", userId).
		Updates(&user).
		Error
	if err != nil {
		return err
	}

	return nil
}
