package Tests

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Models"
	"Polybub/Tests/TestHelpers"
	"Polybub/Utilities"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func Test_Given_Valid_When_NewJwt_Then_ReturnsTokenString(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	name := "asdf"
	userId := int32(1)
	userGroup := int32(1)
	permissions := []Models.Permission{
		OAuth2.NewPerm("FooBar", true, true, true, true),
	}

	// Act
	s, err := OAuth2.NewJwt(name, userId, userGroup, permissions)
	length := len(s)
	bytesize := unsafe.Sizeof(s)

	// Assert
	assert.True(t, bytesize < 20)
	assert.True(t, length > 500)
	assert.Equal(t, nil, err)
}
