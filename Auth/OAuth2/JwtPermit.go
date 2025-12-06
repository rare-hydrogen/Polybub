package OAuth2

import (
	"Polybub/Data/Models"
	"Polybub/Utilities"
	"errors"
	"net/http"
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
	return func(w http.ResponseWriter, req *http.Request) {
		tokenString, err := GetTokenStringFromHeader(req)
		if err != nil {
			failHandle(w, http.StatusUnauthorized, "Unauthorised.")
			return
		}

		claims, err := GetClaimsFromTokenString(tokenString)
		if err != nil {
			failHandle(w, http.StatusUnauthorized, "Unauthorised.")
			return
		}

		GlobalClaims = claims

		if userGroup != nil {
			if claims.Audience != *userGroup {
				failHandle(w, http.StatusForbidden, "Forbidden.")
				return
			}
		}

		hasPerm := CheckPerm(perm, claims.Permissions)

		if !hasPerm {
			failHandle(w, http.StatusForbidden, "Forbidden.")
			return
		}

		handler(w, req)
	}
}

func JwtPermit(mux *http.ServeMux, path string, handler http.HandlerFunc, perm Models.Permission, userGroup *int32) {
	// nil userGroup is public, meaning any group
	authedFunc := getLockHandler(handler, userGroup, perm)
	mux.HandleFunc(path, authedFunc)
}

func JwtPermitRequest(req *http.Request, perm Models.Permission, userGroup *int32) (int, error) {
	tokenString, err := GetTokenStringFromHeader(req)
	if err != nil {
		return http.StatusUnauthorized, errors.New("missing token")
	}

	claims, err := GetClaimsFromTokenString(tokenString)
	if err != nil {
		return http.StatusUnauthorized, errors.New("cannot read token")
	}

	GlobalClaims = claims
	if userGroup != nil {
		if claims.Audience != *userGroup {
			return http.StatusForbidden, errors.New("missing group")
		}
	}

	hasPerm := CheckPerm(perm, claims.Permissions)
	if !hasPerm {
		return http.StatusForbidden, errors.New("missing permission")
	}

	return 0, nil
}
