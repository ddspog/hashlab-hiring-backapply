// +build dev

package req

import "net/http"

// Queries contains project assets.
var Queries http.FileSystem = http.Dir("queries")
