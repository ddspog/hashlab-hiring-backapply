package today

import (
	"testing"

	"github.com/ddsgok/bdd"
	"github.com/ddsgok/gql"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server/info"
)

// Feature Check if Today is User's birthday
// - As a developer,
// - I want to be able to use GQL to fetch user information,
// - So that I can discover if today is user's birthday.
func Test_Use_gql_to_check_if_today_is_Bday_of_User(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Today string `yaml:"today"`
		ID    string `yaml:"id"`
	}{}, &struct {
		Bday string `yaml:"bday"`
	}{}

	given(t, "today is %[input.today]q, and we have a testing GraphQLDB on info.TestGraphQLServer", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)

		MockNow(input.Today)
		defer ResetMocks()

		gql.ConnectAt(info.TestGraphQLServer)
		defer gql.Disconnect()

		ans, err := IsUserBDay(input.ID)
		when("ans, err := IsUserBDay(%[input.id]q) is called", func(it bdd.It) {
			it("shouldn't return any error", func(assert bdd.Assert) {
				assert.Nil(err)
			})

			if Date() == gold.Bday {
				it("should return true", func(assert bdd.Assert) {
					assert.True(ans)
				})
			} else {
				it("should return false", func(assert bdd.Assert) {
					assert.False(ans)
				})
			}
		})
	})
}

// Feature Check if Today is Black Friday
// - As a developer,
// - I want to be able to verify current data,
// - So that I can discover if today is Black Friday.
func Test_Check_if_today_is_Black_Friday(t *testing.T) {
	given := bdd.Sentences().Golden()

	input, gold := &struct {
		Today string `yaml:"today"`
	}{}, &struct{}{}

	given(t, "today is %[input.today]q, and we have a testing GraphQLDB on info.TestGraphQLServer", func(when bdd.When, golden bdd.Golden) {
		golden.Load(input, gold)

		MockNow(input.Today)
		defer ResetMocks()

		gql.ConnectAt(info.TestGraphQLServer)
		defer gql.Disconnect()

		ans := IsBlackFriday()
		when("ans := IsBlackFriday() is called", func(it bdd.It) {
			if Date() == "11-25" {
				it("should return true", func(assert bdd.Assert) {
					assert.True(ans)
				})
			} else {
				it("should return false", func(assert bdd.Assert) {
					assert.False(ans)
				})
			}
		})
	})
}
