package Totp

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

	// Display the QR code to the user.
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
