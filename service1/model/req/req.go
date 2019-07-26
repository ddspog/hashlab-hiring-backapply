package req

//go:generate go run -tags=dev ../../queries_generate.go

import (
	"github.com/ddsgok/gql"
)

// ProductByID creates a graphql request to fetch a product by its
// id information.
func ProductByID(prodid string) (req *gql.Request) {
	req = gql.LoadRequest(Queries, "QProductByID.gql").SetHeader("Cache-Control", "no-cache").Var("prodid", prodid)
	return
}

// UserByID creates a graphql request to fetch an user by its id
// information.
func UserByID(userid string) (req *gql.Request) {
	req = gql.LoadRequest(Queries, "QUserByID.gql").SetHeader("Cache-Control", "no-cache").Var("userid", userid)
	return
}
