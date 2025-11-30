package Totp

// 1. User logs in
// 2. User opens setup-mfa page
// 3. Backend generates key, saves it encrypted in db
// 4. Backend sends key + QR code to frontend
// 5. Frontend requests validation code from user
// 6. On successful validation, enable TOTP for that user
// 7. Until successful validation, disable TOTP for that user

import (
	"Polybub/Utilities"

	"github.com/pquerna/otp/totp"

	"bytes"
	"image/png"
)

func BeginTotp(username string) (bytes.Buffer, string, error) {
	// Create a key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      Utilities.GlobalConfig.Domain,
		AccountName: username,
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

/*
func FinishTotp() {
	// Now Validate that the user's successfully added the passcode.
	fmt.Println("Validating TOTP...")
	passcode := promptForPasscode()
	valid := totp.Validate(passcode, key.Secret())
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passcode!")
		os.Exit(1)
	}
}
*/
