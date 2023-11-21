package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_Id_Is_Blank(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{Id: "123"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{Id: "123", Price: 40}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_Final_Price(t *testing.T) {
	order := Order{Id: "123", Price: 35, Tax: 5}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.Id)
	assert.Equal(t, 35.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)
	assert.NoError(t, order.CalculateFinalPrice())
	assert.Equal(t, 40.0, order.FinalPrice)
}
