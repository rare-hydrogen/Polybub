package Jsend

import (
	"Polybub/Utilities"
	"Polybub/Utilities/Logger/WriteLogger"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Status constants
const (
	StatusError    = "error"
	StatusFail     = "fail"
	StatusRedirect = "redirect"
	StatusSuccess  = "success"
)

// Body contains
type Body struct {
	// The status indicates the execution result of request,
	// it can be one of "success", "fail" and "error".
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
}

// CONSTRUCTORS

// New returns a success body with the given data.
func New(data interface{}) Body {
	return Body{
		Status: StatusSuccess,
		Data:   data,
	}
}

// NewFail returns a fail body with the given data.
func NewFail(data interface{}) Body {
	return Body{
		Status: StatusFail,
		Data:   data,
	}
}

// NewError returns a error body with given message.
func NewError(message string, code int, data interface{}) Body {
	return Body{
		Status:  StatusError,
		Message: message,
		Code:    code,
		Data:    data,
	}
}

// Write writes the body to http.ResponseWriter.
// If necessary, the status code can be specified through the third parameter.
func Write(w http.ResponseWriter, body Body, statuses ...int) error {
	w.Header().Set("Content-Type", "application/json")

	if len(statuses) > 0 {
		w.WriteHeader(statuses[0])
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

// Pass plain HTML back to the client without json formatting
func WriteHtml(w http.ResponseWriter, body string, statuses ...int) {
	w.Header().Set("Content-Type", "text/html")

	if len(statuses) > 0 {
		w.WriteHeader(statuses[0])
	}

	b := []byte(body)

	w.Write(b)
}

func MethodNotAllowed(w http.ResponseWriter, req *http.Request) {
	Error(w, req, errors.New("method not allowed"), "method not allowed", http.StatusMethodNotAllowed)
}

func InternalServerError(w http.ResponseWriter, req *http.Request, err error) {
	Error(w, req, err, "something went wrong", http.StatusInternalServerError)
}

// JSENDERS

// Error logs the real error and returns a given error message to the client
func Error(w http.ResponseWriter, req *http.Request, err error, publicMessage string, statuses ...int) error {
	if err == nil {
		WriteLogger.WriteLogger(req.Context(), "error was passed as nil", statuses...)
		return Write(w, NewError(publicMessage, 0, nil), statuses...)
	}

	WriteLogger.WriteLogger(req.Context(), err.Error(), statuses...)
	return Write(w, NewError(publicMessage, 0, nil), statuses...)
}

// Success writes successful body with the given data.
func SuccessRedirect(w http.ResponseWriter, req *http.Request, data interface{}, url string, statuses ...int) {
	redirectURL := Utilities.GetBaseUrl(Utilities.GlobalConfig) + "/" + url
	sm := `<script>PopToast("Success, Redirecting...", "toast notification is-primary"); `
	tm := `setTimeout(() => {window.location.href = '` + redirectURL + `';}, 1000);</script>`
	js := sm + tm

	// Note: fails to redirect with: 'hx-swap=none'
	WriteLogger.WriteLogger(req.Context(), "Success, Redirecting...", statuses...)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, js)
}

// Success writes successful body with the given data.
func Success(w http.ResponseWriter, req *http.Request, data interface{}, statuses ...int) error {
	WriteLogger.WriteLogger(req.Context(), "Success", statuses...)
	return Write(w, New(data), statuses...)
}

func Ui(w http.ResponseWriter, req *http.Request, wrappedBody string, statuses ...int) {
	WriteLogger.WriteLogger(req.Context(), "UI", statuses...)
	WriteHtml(w, wrappedBody)
}
