package main

import (
	"log"
	"os"

	"github.com/32leaves/neke/pkg/generator"
	"github.com/32leaves/neke/pkg/parser"
)

func main() {
	r, err := os.Open("examples/helloworld.neke")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer r.Close()

	ast := &parser.AST{}
	err = parser.NewParser().Parse(r, ast)
	if err != nil {
		log.Fatalf("unable to parse file: %v", err)
	}

	err = generator.GoLang.Render(ast, os.Stdout)
	if err != nil {
		log.Fatalf("error while generating: %v", err)
	}

	os.Stdout.WriteString("// ====== TS ======\n")
	err = generator.Typescript.Render(ast, os.Stdout)
	if err != nil {
		log.Fatalf("error while generating: %v", err)
	}
}
