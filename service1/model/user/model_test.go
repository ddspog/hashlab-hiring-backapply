package user

import (
	"testing"

	"github.com/ddsgok/bdd"
	"github.com/ddsgok/gql"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server/info"
)

// Feature Fetch user by ID
// - As a developer,
// - I want to be able to use GQL to fetch user information,
// - So that I can use his data.
func Test_Fetch_user_by_ID(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		ID string `yaml:"id"`
	}{}, &struct {
		Bday string `yaml:"bday"`
	}{}

	given(t, "a testing GraphQLDB on info.TestGraphQLServer", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)

		gql.ConnectAt(info.TestGraphQLServer)
		defer gql.Disconnect()

		usr, err := FetchByID(input.ID)
		when("usr, err := FetchByID(%[input.id]q) is called", func(it bdd.It) {
			if gold.Bday == "" {
				it("should return an 'no user found' error", func(assert bdd.Assert) {
					if assert.Error(err) {
						assert.Contains(err.Error(), "no user found")
					}
				})
			} else {
				it("shouldn't return any error", func(assert bdd.Assert) {
					assert.Nil(err)
				})
	
				it("should have usr.ID return %[input.id]q", func(assert bdd.Assert) {
					assert.Equal(usr.ID, input.ID)
				})
	
				it("should have usr.ID return %[golden.bday]q", func(assert bdd.Assert) {
					assert.Equal(usr.BDay.Format("2006-01-02"), gold.Bday)
				})
			}

		})
	})
}
