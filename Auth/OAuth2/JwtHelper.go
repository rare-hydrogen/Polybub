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
	jwtObj, err := ParseJwt(tokenString)
	if err != nil {
		return Claims{}, err
	}

	claims := jwtObj.Claims.(jwt.MapClaims)

	var c Claims
	c.Name = claims["nme"].(string)
	c.Issuer = claims["iss"].(string)
	c.Subject = int32(claims["sub"].(float64))
	c.Audience = int32(claims["aud"].(float64))
	c.Expiration = time.Unix(int64(claims["exp"].(float64)), 0)
	c.NotBefore = time.Unix(int64(claims["nbf"].(float64)), 0)
	c.IssuedAt = time.Unix(int64(claims["iat"].(float64)), 0)

	return c, nil
}

func StoreTokenAndRedirect(w http.ResponseWriter, tokenString string, page string) {
	redirectURL := Utilities.GetBaseUrl(Utilities.GlobalConfig) + "/" + page

	cookie := &http.Cookie{
		Name:     Utilities.GlobalConfig.CookieName,
		Value:    tokenString,
		Path:     "/",
		Domain:   Utilities.GetDomain(Utilities.GlobalConfig),
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	js := `<script>window.location.href = '` + redirectURL + `';</script>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, js)
}

func DeleteTokenAndRedirect(w http.ResponseWriter, page string) {
	redirectURL := Utilities.GetBaseUrl(Utilities.GlobalConfig) + "/" + page

	cookie := &http.Cookie{
		Name:     Utilities.GlobalConfig.CookieName,
		Value:    "",
		Path:     "/",
		Domain:   Utilities.GetDomain(Utilities.GlobalConfig),
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	js := `<script>window.location.href = '` + redirectURL + `';</script>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, js)
}
