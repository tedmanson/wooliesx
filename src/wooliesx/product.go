package wooliesx

import (
	"encoding/json"
	"io/ioutil"
)

// GetProducts returns a list of products provided by a call to a WooliesX API
func (s SDK) GetProducts(token string) ([]Product, error) {
	resp, err := s.client.Get(s.url + "resource/products?token=" + token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var products []Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// Product is a single
type Product struct {
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
}
