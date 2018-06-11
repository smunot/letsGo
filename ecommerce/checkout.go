package ecommerce

type (
	product struct {
		item  string
		price float64
	}

	products = []product
)

var ecommerceStore = products{
	{"bread", 10.50},
	{"butter", 8.75},
}

// Input is a struct.
type Input struct {
	Item     string
	Quantity int
}

const taxRate = 7.5

// FinalCost is used to get the total cost of the purchase post tax.
func FinalCost(inputs []Input) (total float64) {

	for _, input := range inputs {
		for _, p := range ecommerceStore {
			if input.Item == p.item {
				total += p.price * float64(input.Quantity)
			}

		}
	}

	total += (total * taxRate) / 100

	return
}
