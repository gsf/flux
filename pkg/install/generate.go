// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shurcooL/vfsgen"
)

func main() {
	usage := func() {
		fmt.Fprintf(os.Stderr, "usage: %s\n", os.Args[0])
		os.Exit(1)
	}
	if len(os.Args) != 1 {
		usage()
	}

	err := vfsgen.Generate(http.Dir("templates/"), vfsgen.Options{
		Filename:     "generated_templates.gogen.go",
		PackageName:  "install",
		VariableName: "templates",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
