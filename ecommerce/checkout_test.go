package ecommerce_test

import (
	"testing"

	"github.com/smunot/letsGo/ecommerce"
	"github.com/stretchr/testify/assert"
)

func TestFinalCost(t *testing.T) {

	inputs := []ecommerce.Input{
		{Item: "bread", Quantity: 3},
		{Item: "butter", Quantity: 3},
	}

	e := 62.081250
	assert.Equal(t, e, ecommerce.FinalCost(inputs))
	// fmt.Printf("The total is %f \n", ecommerce.FinalCost(inputs))
	// fmt.Println("The total is: ", ecommerce.FinalCost(inputs))

}
