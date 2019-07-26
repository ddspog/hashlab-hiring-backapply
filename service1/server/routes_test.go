package server

import (
	"context"
	"testing"

	"github.com/ddsgok/bdd"
	"github.com/ddsgok/gql"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server/info"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server/protocols"
	"github.com/ddspog/hashlab-hiring-backapply/service1/today"
)

// Feature Verify discount
// - As a developer of this project,
// - I want to enable discount verification,
// - So other services can fetch products discounts info.
func Test_Verify_discount(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Today  string `yaml:"today"`
		UserID string `yaml:"userid"`
		ProdID string `yaml:"prodid"`
	}{}, &struct {
		Price int64  `yaml:"price"`
		Bday  string `yaml:"bday"`
	}{}

	given(t, "that today is %[input.today]q a empty ctx Context, a form VerifyDiscountForm with Iduser %[input.userid]q and a Idprod %[input.prodid]q, with a s Server", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)

		today.MockNow(input.Today)
		defer today.ResetMocks()

		gql.ConnectAt(info.TestGraphQLServer)
		defer gql.Disconnect()

		s := New()
		ctx, form := context.Background(), &protocols.VerifyDiscountForm{
			Userid: input.UserID,
			Prodid: input.ProdID,
		}

		info, err := s.VerifyDiscount(ctx, form)
		when("s.VerifyDiscount(ctx, form) is called", func(it bdd.It) {

			if gold.Price == 0 {
				it("should return an 'no product found' error", func(assert bdd.Assert) {
					if assert.Error(err) {
						assert.Contains(err.Error(), "no product found")
					}
				})
			} else if gold.Bday == "" {
				it("should return an 'no user found' error", func(assert bdd.Assert) {
					if assert.Error(err) {
						assert.Contains(err.Error(), "no user found")
					}
				})
			} else {
				it("shouldn't return any error", func(assert bdd.Assert) {
					assert.Nil(err)
				})

				if today.Date() == "11-25" { // Black Friday discount: 10%
					it("should have info.GetPct() equal 0.1", func(assert bdd.Assert) {
						assert.Equal(info.GetPct(), 0.1)
					})

					it("should have info.GetValueInCents() equal 0.1 * %[golden.price]d", func(assert bdd.Assert) {
						assert.Equal(info.GetValueInCents(), int64(0.1*float32(gold.Price)))
					})
				} else if today.Date() == gold.Bday { // Birthday discount: 5%
					it("should have info.GetPct() equal 0.05", func(assert bdd.Assert) {
						assert.Equal(info.GetPct(), 0.05)
					})

					it("should have info.GetValueInCents() equal 0.05 * %[golden.price]d", func(assert bdd.Assert) {
						assert.Equal(info.GetValueInCents(), int64(0.05*float32(gold.Price)))
					})
				} else { // No Discount
					it("should have info.GetPct() equal 0", func(assert bdd.Assert) {
						assert.Equal(info.GetPct(), 0.0)
					})

					it("should have info.GetValueInCents() equal 0", func(assert bdd.Assert) {
						assert.Equal(info.GetValueInCents(), 0)
					})
				}
			}

		})
	})
}
