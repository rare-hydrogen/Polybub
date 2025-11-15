package Tests

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Tests/TestHelpers"
	"Polybub/Utilities"
	"testing"

	money "github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func Test_Given_Valid_When_CreateFooBar_Then_ReturnsFoobar(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	TestHelpers.ApplySchema()
	fooBar := Models.FooBar{
		Name:     "test1",
		Type:     "Internal",
		Amount:   1299,
		Currency: *money.GetCurrency("USD"),
	}

	// Act
	result, err := Services.CreateFooBar(fooBar)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, fooBar.Name, result.Name)
	assert.Equal(t, fooBar.Type, result.Type)
	assert.Equal(t, fooBar.Amount, result.Amount)
	assert.Equal(t, fooBar.Currency, result.Currency)

	mon := money.New(fooBar.Amount, fooBar.Currency.Code).Display()
	assert.Equal(t, mon, "$12.99")
}

func Test_Given_Valid_When_GetSingleFooBar_Then_ReturnsFooBar(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	TestHelpers.ApplySchema()
	fooBar := Models.FooBar{
		Name:     "test1",
		Type:     "Internal",
		Amount:   1299,
		Currency: *money.GetCurrency("USD"),
	}

	// Act
	fb, err := Services.CreateFooBar(fooBar)
	assert.Nil(t, err)
	result, err := Services.ReadSingleFooBar(fb.Id)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, fooBar.Name, result.Name)
	assert.Equal(t, fooBar.Type, result.Type)
	assert.Equal(t, fooBar.Amount, result.Amount)
	assert.Equal(t, fooBar.Currency, result.Currency)

	mon := money.New(fooBar.Amount, fooBar.Currency.Code).Display()
	assert.Equal(t, mon, "$12.99")
}

func Test_Given_Valid_When_GetManyFooBar_Then_ReturnsManyFooBar(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	TestHelpers.ApplySchema()
	fooBar := Models.FooBar{
		Name:     "test1",
		Type:     "Internal",
		Amount:   1299,
		Currency: *money.GetCurrency("USD"),
	}

	// Act
	_, _ = Services.CreateFooBar(fooBar)
	_, _ = Services.CreateFooBar(fooBar)
	_, _ = Services.CreateFooBar(fooBar)
	results, err := Services.ReadManyFooBar()

	// Assert
	for i := 0; i < len(results); i++ {
		assert.Nil(t, err)
		assert.Equal(t, fooBar.Name, results[i].Name)
		assert.Equal(t, fooBar.Type, results[i].Type)
		assert.Equal(t, fooBar.Amount, results[i].Amount)
		assert.Equal(t, fooBar.Currency, results[i].Currency)
	}
}

func Test_Given_Valid_When_UpdateFooBar_Then_ReturnsFooBar(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	TestHelpers.ApplySchema()
	fooBar := Models.FooBar{
		Name:     "test1",
		Type:     "Internal",
		Amount:   1299,
		Currency: *money.GetCurrency("USD"),
	}

	// Act
	fb, err := Services.CreateFooBar(fooBar)
	assert.Nil(t, err)
	fb.Name = "test2"
	result, err := Services.UpdateFooBar(fb)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, fb.Name, result.Name)
}

func Test_Given_Valid_When_DeleteFooBar_Then_NoFooBar(t *testing.T) {
	// Arrange
	Utilities.GlobalConfig = TestHelpers.UniqueTestConfig()
	TestHelpers.ApplySchema()
	fooBar := Models.FooBar{
		Name:     "test1",
		Type:     "Internal",
		Amount:   1299,
		Currency: *money.GetCurrency("USD"),
	}

	// Act
	fb, err := Services.CreateFooBar(fooBar)
	assert.Nil(t, err)
	err = Services.SoftDeleteFooBar(fb.Id)
	assert.Nil(t, err)
	_, err = Services.ReadSingleFooBar(fb.Id)

	// Assert
	assert.Equal(t, err.Error(), "record not found")
}
