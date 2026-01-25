package OAuth2

import (
	"Polybub/Data/Models"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/golang-jwt/jwt/v5"
)

func CompressPermsForClaims(permissions []Models.Permission) (string, error) {
	// Get bytes
	permBytes, err := json.Marshal(permissions)
	if err != nil {
		return "", err
	}

	// Compress bytes with gzip
	var buf bytes.Buffer
	pw := gzip.NewWriter(&buf)
	_, err = pw.Write(permBytes)
	if err != nil {
		return "", err
	}
	pw.Close()

	// Encode and return as a string
	compressedPermissions := base64.StdEncoding.EncodeToString(buf.Bytes())
	return compressedPermissions, nil
}

func DecompressPermsFromClaims(claims jwt.MapClaims) ([]Models.Permission, error) {
	// Read the perms claim from the token to bytes
	decoded, err := base64.StdEncoding.DecodeString(claims["prm"].(string))
	if err != nil {
		return []Models.Permission{}, err
	}

	// Decompress the bytes
	re := bytes.NewReader(decoded)
	gzre, err := gzip.NewReader(re)
	if err != nil {
		return []Models.Permission{}, err
	}
	output, err := io.ReadAll(gzre)
	if err != nil {
		return []Models.Permission{}, err
	}

	// Unmarshal the bytes back to perms
	var perms []Models.Permission
	err = json.Unmarshal(output, &perms)
	if err != nil {
		return []Models.Permission{}, err
	}

	return perms, nil
}
