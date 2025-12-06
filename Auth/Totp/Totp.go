package Totp

// 1. User logs in
// 2. User opens setup-mfa page
// 3. Backend generates key, saves it encrypted in db
// 4. Backend sends key + QR code to frontend
// 5. Frontend requests validation code from user
// 6. On successful validation, enable TOTP for that user
// 7. Until successful validation, disable TOTP for that user

import (
	"Polybub/Data/Models"
	"Polybub/Utilities"
	"errors"

	"github.com/pquerna/otp/totp"

	"bytes"
	"image/png"
)

func encryptTotpKey(TotpKey string) (string, error) {
	return TotpKey, nil
}

func decryptTotpKey(EncryptedTotpKey string) (string, error) {
	return EncryptedTotpKey, nil
}

func BeginTotp(user Models.User) (bytes.Buffer, string, error) {
	// Create a key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      Utilities.GlobalConfig.Domain,
		AccountName: user.Username,
	})
	if err != nil {
		return bytes.Buffer{}, "", err
	}

	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		return bytes.Buffer{}, "", err
	}
	png.Encode(&buf, img)

	// display the QR code to the user.
	return buf, key.Secret(), err
}

func CheckTotp(user Models.User, code string) error {
	key, err := decryptTotpKey(user.TotpKey)
	if err != nil {
		return err
	}

	valid := totp.Validate(code, key)
	if valid {
		return nil
	} else {
		return errors.New("invalid code")
	}
}

func KeyMatchTotp(key string, code string) error {
	_, err := decryptTotpKey(key)
	if err != nil {
		return err
	}

	valid := totp.Validate(code, key)
	if valid {
		return nil
	} else {
		return errors.New("invalid key")
	}
}
