package info

var (
	// TodayDate is a mocking variable to make tests on API level.
	TodayDate = getenv("MOCKED_TODAY_DATE", "")

	// TestGraphQLServer contains simple data used on tests.
	TestGraphQLServer = getenv("TEST_GRAPHQLSRV", "https://hash-ddspog-apply-dbhasura.herokuapp.com/v1/graphql")
)
