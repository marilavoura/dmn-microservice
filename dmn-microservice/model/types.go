package model

// Invoice

type DecisionInput struct {
	Id				int			`json:"id"`
	Supplier		string		`json:"supplier"`
	Price			float64		`json:"price"`
	Type			string		`json:"type"`
}

type DecisionOutput struct {
	Id				int			`json:"id"`
	Taxes			float64		`json:"taxes"`
	Discounts		float64		`json:"discounts"`
}