package Services

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Auth/Totp"
	"Polybub/Data"
	"Polybub/Data/Models"
	"errors"

	"github.com/google/uuid"
)

func getJwtFromUser(user Models.User) (Models.User, string, error) {
	permissions, err := ReadUsersPermissions(user.Id)
	if err != nil {
		return Models.User{}, "", err
	}

	name := user.FirstName + " " + user.LastName
	jwtString, err := OAuth2.NewJwt(name, user.Id, user.UserGroup, permissions)
	if err != nil {
		return Models.User{}, "", err
	}

	return user, jwtString, nil
}

func Login(username string, password string) (Models.User, string, error) {
	user, err := UnsafeGetUserByUsername(username)
	if err != nil {
		return Models.User{}, "", err
	}

	if user.Password == "" {
		return Models.User{}, "", errors.New("no password")
	}

	encryptedPassword := OAuth2.EncryptPassword(password, user.Salt)

	if encryptedPassword != user.Password {
		return Models.User{}, "", errors.New("incorrect password")
	}

	return getJwtFromUser(user)
}

func MfaLogin(userId int32, code string) (Models.User, string, error) {
	// get the user
	user, err := UnsafeReadSingleUser(userId)
	if err != nil {
		return Models.User{}, "", err
	}

	// check if the code is valid
	err = Totp.CheckTotp(user, code)
	if err != nil {
		return Models.User{}, "", err
	}

	// if valid, return the user and jwt
	return getJwtFromUser(user)
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
