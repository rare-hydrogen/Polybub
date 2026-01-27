package OAuth2

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	"Polybub/Data/Models"
	"Polybub/Utilities"

	"github.com/golang-jwt/jwt/v5"
)

func readPrivateKey() (*rsa.PrivateKey, error) {
	// Get the file
	pd, err := os.ReadFile("./.certs/private.pem")
	if err != nil {
		return nil, errors.New("missing private key")
	}

	// Decode
	block, _ := pem.Decode(pd)
	if block == nil {
		return nil, errors.New("failed to decode private key")
	}

	// Parse
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key")
	}

	return privateKey.(*rsa.PrivateKey), nil
}

func NewJwt(name string, userId int32, userGroup int32, permissions []Models.Permission) (string, error) {
	// Get the private key first
	key, err := readPrivateKey()
	if err != nil {
		return "", err
	}

	// Get compressed perms
	compressedPermissions, err := CompressPermsForClaims(permissions)
	if err != nil {
		return "", err
	}

	// Map everything to claims
	claims := jwt.MapClaims{
		"nme": name,
		"iss": Utilities.GlobalConfig.Domain,
		"sub": userId,
		"aud": userGroup,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"prm": compressedPermissions,
	}

	// Sign the token so it can't be fungible
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}

func ParseJwt(tokenString string) (jwt.Token, error) {
	// Get the private key
	key, err := readPrivateKey()
	if err != nil {
		return jwt.Token{}, err
	}

	// Get the expression for the parse
	publicKey := &key.PublicKey
	var expression = func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	}

	// Parse the token string
	jwtToken, err := jwt.Parse(tokenString, expression)
	if err != nil {
		return jwt.Token{}, err
	}

	// If token isn't expired or wasn't signed or was changed
	if !jwtToken.Valid {
		return jwt.Token{}, errors.New("invalid token")
	}

	return *jwtToken, nil
}
