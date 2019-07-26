package product

import (
	"github.com/ddsgok/gql"
	"github.com/ddsgok/str"
	"github.com/ddspog/hashlab-hiring-backapply/service1/model/req"
)

// Product represents the product modeled on this service DB. Contains
// only needed information for the service.
type Product struct {
	ID           string
	PriceInCents int64
}

// New creates a new product, using information received.
func New(id string, price int64) (p *Product) {
	p = &Product{id, price}
	return
}

// FetchByID retrieve a product from DB, using his ID as key.
func FetchByID(id string) (p *Product, err error) {
	var res gql.Response
	if res, err = req.ProductByID(id).Run(); err == nil {
		if res.Get("Product.#").Int() > 0 {
			p = New(id, res.Get("Product.0.price_in_cents").Int())
		} else {
			err = str.New("no product found with id=%v", id).Error()
		}
	}

	return
}
