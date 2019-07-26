// +build ignore

package main

import (
	"github.com/ddsgok/log"

	"github.com/ddspog/hashlab-hiring-backapply/service1/model/req"
	"github.com/shurcooL/vfsgen"
)

func main() {
	if err := vfsgen.Generate(req.Queries, vfsgen.Options{
		PackageName:  "req",
		BuildTags:    "!dev",
		VariableName: "Queries",
	}); err != nil {
		log.Fatal(err.Error())
	}
}
