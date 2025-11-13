package OAuth2

import (
	"Polybub/Data/Models"
	"Polybub/Utilities"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var GlobalClaims Claims

func NewPerm(name string, isCreate bool, isRead bool, isUpdate bool, isDelete bool) Models.Permission {
	return Models.Permission{
		Name:     name,
		IsCreate: isCreate,
		IsRead:   isRead,
		IsUpdate: isUpdate,
		IsDelete: isDelete,
	}
}

func CheckPerm(reqPerm Models.Permission, checkPerms []Models.Permission) bool {
	var hasPerm = false
	var checkPerm Models.Permission

	var m int
	for i := 0; i < len(checkPerms); i++ {
		if reqPerm.Name == checkPerms[i].Name {
			checkPerm = checkPerms[i]
			m++
		}
	}

	if m > 1 {
		return hasPerm
	}

	if reqPerm.Name == "" {
		return hasPerm
	}

	if reqPerm.Name != checkPerm.Name {
		return hasPerm
	}

	if reqPerm.IsCreate {
		if !checkPerm.IsCreate {
			return hasPerm
		}
	}

	if reqPerm.IsRead {
		if !checkPerm.IsRead {
			return hasPerm
		}
	}

	if reqPerm.IsUpdate {
		if !checkPerm.IsUpdate {
			return hasPerm
		}
	}

	if reqPerm.IsDelete {
		if !checkPerm.IsDelete {
			return hasPerm
		}
	}

	return true
}

func failHandle(w http.ResponseWriter, code int, message string) {
	realm := Utilities.GlobalConfig.Domain

	w.Header().Set("WWW-Authenticate", `Bearer realm="`+realm+`"`)
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func getLockHandler(handler http.HandlerFunc, userGroup *int32, perm Models.Permission) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := GetTokenStringFromHeader(r)
		if err != nil {
			failHandle(w, http.StatusUnauthorized, "Unauthorised.")
			return
		}

		ec, err := GetClaimsFromTokenString(tokenString)
		if err != nil {
			failHandle(w, http.StatusUnauthorized, "Unauthorised.")
			return
		}

		GlobalClaims = ec

		if userGroup != nil {
			if ec.Audience != *userGroup {
				failHandle(w, http.StatusForbidden, "Forbidden.")
				return
			}
		}

		hasPerm := CheckPerm(perm, ec.Permissions)

		if !hasPerm {
			failHandle(w, http.StatusForbidden, "Forbidden.")
			return
		}

		handler(w, r)
	}
}

func JwtPermit(mux *http.ServeMux, path string, handler http.HandlerFunc, perm Models.Permission, userGroup *int32) {
	// nil userGroup is public, meaning any group
	authedFunc := getLockHandler(handler, userGroup, perm)
	mux.HandleFunc(path, authedFunc)
}

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
