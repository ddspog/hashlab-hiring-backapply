package info

import "os"

var (
	// Name of the server.
	Name = "Service 001"

	// Host of the server.
	Host = getenv("SERVICE1_HOST", "")

	// Port used on the server.
	Port = getenv("SERVICE1_PORT", "5001")

	// GraphQLServer used to fetch data on server.
	GraphQLServer = getenv("SERVICE1_GRAPHQLSRV", "http://localhost:8080/v1/graphql")
)

// getenv fetch env var value, and prints a default if env ain't
// defined.
func getenv(e, d string) (v string) {
	v = os.Getenv(e)

	if v == "" {
		v = d
	}

	return
}
