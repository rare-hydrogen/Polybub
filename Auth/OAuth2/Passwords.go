package OAuth2

import (
	"Polybub/Utilities"

	"golang.org/x/crypto/argon2"
)

func EncryptPassword(password string, salt string) string {
	pepper := Utilities.GlobalConfig.Pepper

	idKey := argon2.IDKey([]byte(password+pepper), []byte(salt), 3, 32*1024, 4, 32)

	return string(idKey)
}
