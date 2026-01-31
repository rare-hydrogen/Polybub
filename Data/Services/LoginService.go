package Services

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Auth/Totp"
	"Polybub/Data"
	"Polybub/Data/Models"
	"context"
	"errors"

	"github.com/google/uuid"
)

func getJwtFromUser(ctx context.Context, user Models.User) (Models.User, string, error) {
	permissions, err := ReadUsersPermissions(ctx, user.Id)
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

func Login(ctx context.Context, username string, password string) (Models.User, string, error) {
	user, err := UnsafeGetUserByUsername(ctx, username)
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

	return getJwtFromUser(ctx, user)
}

func MfaLogin(ctx context.Context, userId int32, code string) (Models.User, string, error) {
	user, err := UnsafeReadSingleUser(ctx, userId)
	if err != nil {
		return Models.User{}, "", err
	}

	err = Totp.CheckTotp(user, code)
	if err != nil {
		return Models.User{}, "", err
	}

	return getJwtFromUser(ctx, user)
}

func MfaUpdate(ctx context.Context, userId int32, key string, code string) (Models.User, string, error) {
	// Simple check to make sure the supplied key and code work together
	// because the user could submit a key that wasn't made by this backend
	err := Totp.KeyMatchTotp(key, code)
	if err != nil {
		return Models.User{}, "", err
	}

	_, err = UpdateUser(ctx, Models.User{
		Id:      userId,
		TotpKey: key,
	})
	if err != nil {
		return Models.User{}, "", err
	}

	user, err := UnsafeReadSingleUser(ctx, userId)
	if err != nil {
		return Models.User{}, "", err
	}
	return getJwtFromUser(ctx, user)
}

func UpdatePasswordAndSalt(ctx context.Context, userId int32, password string) error {
	var db = Data.GetConnection().WithContext(ctx)
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
