package OAuth2

import (
	"Polybub/Data/Models"
	"Polybub/Utilities"
	"errors"
	"net/http"
)

func NewPerm(name string, isCreate bool, isRead bool, isUpdate bool, isDelete bool) Models.Permission {
	// Easily format a new perms object
	return Models.Permission{
		Name:     name,
		IsCreate: isCreate,
		IsRead:   isRead,
		IsUpdate: isUpdate,
		IsDelete: isDelete,
	}
}

func CheckPerm(reqPerm Models.Permission, checkPerms []Models.Permission) bool {
	// By default, not allowed
	var hasPerm = false
	var checkPerm Models.Permission

	// Find which perm matches the required perm
	// And how many perms match
	var m int
	for i := 0; i < len(checkPerms); i++ {
		if reqPerm.Name == checkPerms[i].Name {
			checkPerm = checkPerms[i]
			m++
		}
	}

	// Multiple identical permissions are not allowed
	if m > 1 {
		return hasPerm
	}

	// Unnamed permissions are not allowed
	if reqPerm.Name == "" {
		return hasPerm
	}

	// If somehow, the names don't match, block access
	if reqPerm.Name != checkPerm.Name {
		return hasPerm
	}

	// If one of the CRUD operations doesn't match, block access
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

	// If all those checks are valid, the user has at least the minimum correct permission
	return true
}

func failHandle(w http.ResponseWriter, code int, message string) {
	// Just handle this the same way every time
	realm := Utilities.GlobalConfig.Domain

	w.Header().Set("WWW-Authenticate", `Bearer realm="`+realm+`"`)
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func getLockHandler(handler http.HandlerFunc, userGroup *int32, perm Models.Permission) http.HandlerFunc {
	// Just handle this the same way
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

		if userGroup != nil {
			if claims.Audience != *userGroup {
				failHandle(w, http.StatusForbidden, "Forbidden.")
				return
			}
		}

		// Make sure the user has the required permission
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
	// Get back the normal handler after we handled perms
	mux.HandleFunc(path, authedFunc)
}

func JwtPermitRequest(req *http.Request, perm Models.Permission, userGroup *int32) (int, error) {
	// Get the token string and claims
	tokenString, err := GetTokenStringFromHeader(req)
	if err != nil {
		return http.StatusUnauthorized, errors.New("missing token")
	}

	claims, err := GetClaimsFromTokenString(tokenString)
	if err != nil {
		return http.StatusUnauthorized, errors.New("cannot read token")
	}

	if userGroup != nil {
		if claims.Audience != *userGroup {
			return http.StatusForbidden, errors.New("missing group")
		}
	}

	// Make sure the user has the required permission
	hasPerm := CheckPerm(perm, claims.Permissions)
	if !hasPerm {
		return http.StatusForbidden, errors.New("missing permission")
	}

	return 0, nil
}
