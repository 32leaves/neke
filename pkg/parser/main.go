// nolint: govet, golint
// based on https://raw.githubusercontent.com/alecthomas/participle/master/_examples/protobuf/main.go
package parser

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/alecthomas/repr"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

const CommonLang = "common"

type AST struct {
	Pos lexer.Position

	Entries []*Entry `{ @@ }`
}

type Entry struct {
	Pos lexer.Position

	Interface *Interface `@@`
	Struct    *Struct    `| @@`
	Enum      *Enum      `| @@`
}

type Interface struct {
	Pos lexer.Position

	Name  string    `"interface" @Ident`
	Entry []*Method `"{" (@@)* "}"`
}

type Method struct {
	Pos lexer.Position

	Name     string `"func" @Ident`
	Request  *Types `"(" @@ ")"`
	Response *Types `"returns" "(" @@ ")"`
}

type Enum struct {
	Pos lexer.Position

	Name   string   `"enum" @Ident`
	Values []string `"{" (@Ident)* "}"`
}

type Struct struct {
	Pos lexer.Position

	Name   string   `"struct" @Ident`
	Fields []*Field `"{" (@@)* "}"`
}

type Field struct {
	Pos lexer.Position

	Optional bool `( @"optional"`
	Required bool `| @"required")`

	Name  string `@Ident`
	Types *Types `@@`
}

type Types struct {
	Pos lexer.Position

	Children []*Type `@@+`
}

func (t *Types) ByLang(lang string) string {
	if t == nil {
		return ""
	}

	for _, t := range t.Children {
		if t.Lang == lang {
			return t.Name
		}
	}
	for _, t := range t.Children {
		if t.Lang == CommonLang {
			return t.Name
		}
	}
	return ""
}

type Type struct {
	Pos lexer.Position

	Lang string `@Ident ":"`
	Name string `@Ident`
}

func NewParser() *participle.Parser {
	return participle.MustBuild(
		&AST{},
		participle.UseLookahead(1),
	)
}

var (
	cli struct {
		Files []string `required existingfile arg help:"Protobuf files."`
	}
)

func main() {
	ctx := kong.Parse(&cli)

	for _, file := range cli.Files {
		fmt.Println(file)

		r, err := os.Open(file)
		ctx.FatalIfErrorf(err, "")

		ast := &AST{}
		err = NewParser().Parse(r, ast)
		ctx.FatalIfErrorf(err, "")

		repr.Println(ast, repr.Hide(&lexer.Position{}))
	}
}
