// +build ignore

package main

import (
	"log"

	"github.com/rh/for-sake-of-testing/apps/goservd/pkg/sitedata"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(sitedata.Assets, vfsgen.Options{
		PackageName:  "sitedata",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})

	if err != nil {
		log.Fatalln(err)
	}
}
