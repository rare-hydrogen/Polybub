package OAuth2

import (
	"Polybub/Utilities"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func EncryptPassword(password string, salt string) string {
	// Extra protection
	pepper := Utilities.GlobalConfig.Pepper

	// Strong hash and salt
	idKey := argon2.IDKey([]byte(password+pepper), []byte(salt), 3, 32*1024, 4, 32)

	return base64.StdEncoding.EncodeToString(idKey)
}
