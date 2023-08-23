package model

type Tests struct {
	Id        int     `json:"id"`
	Supplier  string  `json:"supplier"`
	Price     float64 `json:"price"`
	Type      string  `json:"type"`
	Taxes     float64 `json:"taxes"`
	Discounts float64 `json:"discounts"`
}

func NewTests() *[8]Tests {
	tests := [8]Tests{
		{Id: 0, Supplier: "Betina", Price: 500, Type: "Internacional", Taxes: 15, Discounts: 5},
		{Id: 1, Supplier: "Betina", Price: 2000.50, Type: "Internacional", Taxes: 10, Discounts: 4},
		{Id: 2, Supplier: "Betina", Price: 8000, Type: "Nacional", Taxes: 40, Discounts: 20},
		{Id: 3, Supplier: "Almir", Price: 12000, Type: "Internacional", Taxes: 30, Discounts: 14},
		{Id: 4, Supplier: "Almir", Price: 16000, Type: "Internacional", Taxes: 50, Discounts: 20},
		{Id: 5, Supplier: "Almir", Price: 60000, Type: "Internacional", Taxes: 50, Discounts: 24},
		{Id: 6, Supplier: "Cecília", Price: 35000, Type: "Internacional", Taxes: 35, Discounts: 30.5},
		{Id: 7, Supplier: "Clarice", Price: 500, Type: "Nacional", Taxes: -1, Discounts: -1},
	}

	return &tests
}
