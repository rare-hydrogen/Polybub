package Tests

import (
	"Polybub/Auth/BasicAuth"
	"Polybub/Routes"
	"Polybub/Tests/TestHelpers"
	"Polybub/Utilities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Given_Valid_When_BasicAuth_Then_Success(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	mux := Routes.AddRoutes()
	mux.HandleFunc("/asdf", BasicAuth.BasicAuth(TestHelpers.TestHandler, "username", "password"))

	// Act
	req, err := http.NewRequest("POST", "/asdf", nil)
	req.SetBasicAuth("username", "password")
	resp := httptest.NewRecorder()
	mux.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, 200, resp.Result().StatusCode)
	assert.Equal(t, nil, err)
}
