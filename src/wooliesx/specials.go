package wooliesx

type prod struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
type quant struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
}

type special struct {
	Quantities []quant `json:"quantities"`
	Total      float32 `json:"total"`
}

func getSpecials(product Product) []special {
	var specials []special

	var q []quant
	q = append(q, quant{
		Name:     product.Name,
		Quantity: 2,
	})

	specials = append(specials, special{
		Quantities: q,
		Total:      float32(product.Price / 2),
	})

	return specials
}
