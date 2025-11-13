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
	pd, err := os.ReadFile("./Certs/private.pem")
	if err != nil {
		return nil, errors.New("missing private key")
	}

	block, _ := pem.Decode(pd)
	if block == nil {
		return nil, errors.New("failed to decode private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key")
	}

	return privateKey.(*rsa.PrivateKey), nil
}

func NewJwt(name string, userId int32, userGroup int32, permissions []Models.Permission) (string, error) {
	key, err := readPrivateKey()
	if err != nil {
		return "", err
	}

	compressedPermissions, err := CompressPermsForClaims(permissions)
	if err != nil {
		return "", err
	}

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

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}

func ParseJwt(tokenString string) (jwt.Token, error) {
	key, err := readPrivateKey()
	if err != nil {
		return jwt.Token{}, err
	}

	publicKey := &key.PublicKey
	var expression = func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	}

	jwtToken, err := jwt.Parse(tokenString, expression)
	if err != nil {
		return jwt.Token{}, err
	}

	if !jwtToken.Valid {
		return jwt.Token{}, errors.New("invalid token")
	}

	return *jwtToken, nil
}
