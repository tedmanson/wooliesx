package wooliesx

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strings"
)

// GetProducts returns a list of products provided by a call to a WooliesX API
func (s SDK) GetProducts(token string) (*Products, error) {
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

	return &Products{
		Products: products,
	}, nil
}

type Products struct {
	Products []Product
}

func (p Products) Sort(sortOption string) {
	switch strings.ToLower(sortOption) {
	case "ascending":
		sort.Slice(p.Products, func(i, j int) bool {
			return p.Products[i].Name < p.Products[j].Name
		})
		break

	case "descending":
		sort.Slice(p.Products, func(i, j int) bool {
			return p.Products[i].Name > p.Products[j].Name
		})
		break

	case "low":
		sort.Slice(p.Products, func(i, j int) bool {
			return p.Products[i].Price < p.Products[j].Price
		})
		break

	case "high":
		sort.Slice(p.Products, func(i, j int) bool {
			return p.Products[i].Price > p.Products[j].Price
		})
		break

	case "recommended":
		//TODO "Recommended" - this will call the "shopperHistory" resource to get a list of customers orders and needs to return based on popularity,
		break
	}
}

// Product is a single
type Product struct {
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
}
