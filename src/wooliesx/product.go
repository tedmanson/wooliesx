package wooliesx

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/labstack/gommon/log"
)

// Product is a single prodict listing
type Product struct {
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
}

// Sale is a reference for each sale
type Sale struct {
	CustomerID int       `json:"customerId,omitempty"`
	Products   []Product `json:"products,omitempty"`
}

// GetProducts returns a list of products provided by a call to a WooliesX API
func (s SDK) GetProducts() ([]Product, error) {
	resp, err := s.client.Get(s.url + "resource/products?token=" + s.token)
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

// GetShopperHistory returns a list of sales provided by a call to a WooliesX API
func (s SDK) GetShopperHistory() ([]Sale, error) {
	resp, err := s.client.Get(s.url + "resource/shopperHistory?token=" + s.token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sales []Sale
	err = json.Unmarshal(body, &sales)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

// SortProducts will return a sorted list of products.
func (s SDK) SortProducts(products []Product, sortOption string) []Product {
	switch strings.ToLower(sortOption) {
	case "ascending":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
		break

	case "descending":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name > products[j].Name
		})
		break

	case "low":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
		break

	case "high":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price > products[j].Price
		})
		break

	case "recommended":
		sales, err := s.GetShopperHistory()
		if err != nil {
			// Better to return an unsorted list of items then no items at all
			log.Error(err)
			return products
		}

		list := findTopProducts(sales)

		var sorted []Product
		for _, s := range list {
			for _, p := range products {
				if p.Name == s.Key {
					sorted = append(sorted, p)
				}
			}
		}

		var complete []Product
		complete = sorted

		for _, p := range products {
			found := false
			for _, s := range sorted {
				if p.Name == s.Name {
					found = true
				}
			}
			if !found {
				complete = append(complete, p)
			}
		}

		products = complete

		break
	}

	return products
}

func findTopProducts(sales []Sale) []Pair {
	var productCount = make(map[string]int)

	for _, sale := range sales {
		for _, product := range sale.Products {
			if _, ok := productCount[product.Name]; ok {
				productCount[product.Name]++
			} else {
				productCount[product.Name] = 1
			}
		}
	}

	pl := make(PairList, len(productCount))
	i := 0
	for k, v := range productCount {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	return pl
}

// Pair allows for a key value representation of a map
type Pair struct {
	Key   string
	Value int
}

// PairList is the conversion of a map to struct
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
