package RequestMiddleware

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Utilities/Logger"
	"bytes"
	"context"
	"io"
	"net/http"
)

// Smuggle the request details back into the request's context
func LogHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := UpdateContextFromRequest(req)
		newReq := req.WithContext(ctx)
		h.ServeHTTP(w, newReq)
	}

	return http.HandlerFunc(fn)
}

// Given a request, produce a new context from the old one with details
func UpdateContextFromRequest(req *http.Request) context.Context {
	// Begin with existing context
	ctx := req.Context()

	// Create starting details
	rdv := Logger.RequestDetails{
		UserId:      0,
		UserName:    "",
		UserGroupId: 0,
		Verb:        req.Method,
		Endpoint:    req.RequestURI,
		QueryString: req.URL.RawQuery,
		Body:        "",
	}

	// Use token to update details
	ts, err := OAuth2.GetTokenStringFromHeader(req)
	if err == nil {
		claims, err := OAuth2.GetClaimsFromTokenString(ts)
		if err == nil {
			rdv.UserId = claims.Subject
			rdv.UserName = claims.Name
			rdv.UserGroupId = claims.Audience
		}
	}

	// Parse the body from the request
	buf, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewReader(buf)) // Allows buf to be read twice
	rdv.Body = string(buf)

	// Finally, return the new context with requestDetails added
	return context.WithValue(ctx, "requestDetails", rdv)
}
