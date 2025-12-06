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
	permBytes, err := json.Marshal(permissions)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	pw := gzip.NewWriter(&buf)
	_, err = pw.Write(permBytes)
	if err != nil {
		return "", err
	}
	pw.Close()

	compressedPermissions := base64.StdEncoding.EncodeToString(buf.Bytes())
	return compressedPermissions, nil
}

func DecompressPermsFromClaims(claims jwt.MapClaims) ([]Models.Permission, error) {
	decoded, err := base64.StdEncoding.DecodeString(claims["prm"].(string))
	if err != nil {
		return []Models.Permission{}, err
	}

	re := bytes.NewReader(decoded)
	gzre, err := gzip.NewReader(re)
	if err != nil {
		return []Models.Permission{}, err
	}
	output, err := io.ReadAll(gzre)
	if err != nil {
		return []Models.Permission{}, err
	}

	var perms []Models.Permission
	err = json.Unmarshal(output, &perms)
	if err != nil {
		return []Models.Permission{}, err
	}

	return perms, nil
}
