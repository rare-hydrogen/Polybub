package Tests

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Models"
	"Polybub/Tests/TestHelpers"
	"Polybub/Utilities"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func Test_Given_Valid_When_CompressPermsForClaims_Then_ReturnsCompressedString(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	perms := Models.Permission{
		Id:       1,
		UserId:   1,
		Name:     "Foobar",
		IsCreate: true,
		IsRead:   true,
		IsUpdate: true,
		IsDelete: true,
	}

	// Act
	s, err := OAuth2.CompressPermsForClaims([]Models.Permission{perms})

	// Assert
	assert.True(t, len(s) > 20)
	assert.Equal(t, nil, err)
}

func Test_Given_Valid_When_DecompressPermsFromClaims_Then_ReturnsPerms(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	perm := Models.Permission{
		Id:       1,
		UserId:   1,
		Name:     "Foobar",
		IsCreate: true,
		IsRead:   true,
		IsUpdate: true,
		IsDelete: true,
	}

	// Act
	tokenString, _ := OAuth2.NewJwt("asdf", 1, 1, []Models.Permission{perm})
	jwtObj, _ := OAuth2.ParseJwt(tokenString)
	claims := jwtObj.Claims.(jwt.MapClaims)
	dPerms, err := OAuth2.DecompressPermsFromClaims(claims)

	// Assert
	assert.True(t, perm.Id == dPerms[0].Id)
	assert.True(t, perm.UserId == dPerms[0].UserId)
	assert.True(t, perm.Name == dPerms[0].Name)
	assert.True(t, perm.IsCreate == dPerms[0].IsCreate)
	assert.True(t, perm.IsRead == dPerms[0].IsRead)
	assert.True(t, perm.IsUpdate == dPerms[0].IsUpdate)
	assert.True(t, perm.IsDelete == dPerms[0].IsDelete)
	assert.Equal(t, nil, err)
}
