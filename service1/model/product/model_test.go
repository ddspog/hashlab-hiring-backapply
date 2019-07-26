package product

import (
	"testing"

	"github.com/ddsgok/bdd"
	"github.com/ddsgok/gql"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server/info"
)

// Feature Fetch product by ID
// - As a developer,
// - I want to be able to use GQL to fetch product information,
// - So that I can use its data.
func Test_Fetch_product_by_ID(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		ID string `yaml:"id"`
	}{}, &struct {
		Price int64 `yaml:"price"`
	}{}

	given(t, "a testing GraphQLDB on info.TestGraphQLServer", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)

		gql.ConnectAt(info.TestGraphQLServer)
		defer gql.Disconnect()

		prod, err := FetchByID(input.ID)
		when("prod, err := FetchByID(%[input.id]q) is called", func(it bdd.It) {
			if gold.Price == 0 {
				it("should return an 'no product found' error", func(assert bdd.Assert) {
					if assert.Error(err) {
						assert.Contains(err.Error(), "no product found")
					}
				})
			} else {
				it("shouldn't return any error", func(assert bdd.Assert) {
					assert.Nil(err)
				})
	
				it("should have prod.ID return %[input.id]q", func(assert bdd.Assert) {
					assert.Equal(prod.ID, input.ID)
				})
	
				it("should have prod.PriceInCents return %[golden.price]d", func(assert bdd.Assert) {
					assert.Equal(prod.PriceInCents, gold.Price)
				})
			}

		})
	})
}
