//go:build ignore

package main

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
	"log"
)

func main() {
	// generate db diagram
	err := entc.Generate("./schema", &gen.Config{}, entc.Extensions(entviz.Extension{}))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	// generate OpenApi spec
	ex, err := entoas.NewExtension()
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
