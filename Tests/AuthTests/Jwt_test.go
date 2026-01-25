package Tests

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Models"
	"Polybub/Tests/TestHelpers"
	"Polybub/Utilities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Given_Valid_When_NewJwt_Then_ReturnsJwt(t *testing.T) {
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
	s, err := OAuth2.NewJwt("asdf", 1, 1, []Models.Permission{perm})

	// Assert
	assert.True(t, len(s) > 20)
	assert.Equal(t, nil, err)
}

// ParseJwt
func Test_Given_Valid_When_ParseJwt_Then_Passes(t *testing.T) {
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
	s, _ := OAuth2.NewJwt("asdf", 1, 1, []Models.Permission{perm})
	p, err := OAuth2.ParseJwt(s)

	// Assert
	assert.True(t, len(s) > 20)
	assert.True(t, p.Valid == true)
	assert.Equal(t, nil, err)
}
