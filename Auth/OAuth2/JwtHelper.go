package OAuth2

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"Polybub/Utilities"
)

func GetTokenStringFromHeader(req *http.Request) (string, error) {
	// Get the token string from the cookie
	tokenString, err := req.Cookie(Utilities.GlobalConfig.CookieName)
	if err != nil {
		return "", errors.New("cookie error detected")
	}
	if tokenString.Value == "" {
		return "", errors.New("empty cookie detected")
	}

	return tokenString.Value, nil
}

func GetClaimsFromTokenString(tokenString string) (Claims, error) {
	// Get the jwt object from the token string
	jwtObj, err := ParseJwt(tokenString)
	if err != nil {
		return Claims{}, err
	}

	// Turn the compressed claims from the jwt into claims we can use
	claims := jwtObj.Claims.(jwt.MapClaims)
	permissions, err := DecompressPermsFromClaims(claims)
	if err != nil {
		return Claims{}, err
	}

	// Map the values from the claims to a custom claims object with decompressed perms
	var c Claims
	c.Name = claims["nme"].(string)
	c.Issuer = claims["iss"].(string)
	c.Subject = int32(claims["sub"].(float64))
	c.Audience = int32(claims["aud"].(float64))
	c.Expiration = time.Unix(int64(claims["exp"].(float64)), 0)
	c.NotBefore = time.Unix(int64(claims["nbf"].(float64)), 0)
	c.IssuedAt = time.Unix(int64(claims["iat"].(float64)), 0)
	c.Permissions = permissions

	return c, nil
}

func GetClaimsFromRequest(w http.ResponseWriter, req *http.Request) (Claims, error) {
	tokenString, err := GetTokenStringFromHeader(req)
	if err != nil {
		return Claims{}, err
	}

	claims, err := GetClaimsFromTokenString(tokenString)
	if err != nil {
		return Claims{}, err
	}

	return claims, nil
}

func StoreTokenAndRedirect(w http.ResponseWriter, tokenString string, page string) {
	redirectURL := Utilities.GetBaseUrl(Utilities.GlobalConfig) + "/" + page

	// Make the cookie and put it in the user's browser
	cookie := &http.Cookie{
		Name:     Utilities.GlobalConfig.CookieName,
		Value:    tokenString,
		Path:     "/",
		Domain:   Utilities.GetDomain(Utilities.GlobalConfig),
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
		Secure:   Utilities.GlobalConfig.IsSecure,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	// Redirect to wherever the user was actually supposed to go
	js := `<script>window.location.href = '` + redirectURL + `';</script>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, js)
}

func DeleteTokenAndRedirect(w http.ResponseWriter, page string) {
	redirectURL := Utilities.GetBaseUrl(Utilities.GlobalConfig) + "/" + page

	// Make a bad cookie to overwrite the user's
	cookie := &http.Cookie{
		Name:     Utilities.GlobalConfig.CookieName,
		Value:    "",
		Path:     "/",
		Domain:   Utilities.GetDomain(Utilities.GlobalConfig),
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteStrictMode,
		Secure:   Utilities.GlobalConfig.IsSecure,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	// Redirect to wherever the user was actually supposed to go
	js := `<script>window.location.href = '` + redirectURL + `';</script>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, js)
}
