package parser

import (
	"strings"
	"testing"

	"github.com/alecthomas/participle/lexer"
	"github.com/go-test/deep"
)

func TestBasicStruct(t *testing.T) {
	code := `
struct Foobar {
    required firstField     common:string
    required anotherField  go:int32 ts:number
    optional optionalField go:bool  ts:boolean
}
`
	r := &AST{}
	err := NewParser().Parse(strings.NewReader(code), r)
	if err != nil {
		t.Error(err)
		return
	}

	if len(r.Entries) == 0 || r.Entries[0] == nil {
		t.Errorf("did not find any entry")
		return
	}
	if r.Entries[0].Struct == nil {
		t.Errorf("first entry is not an struct")
		return
	}

	strct := r.Entries[0].Struct
	if strct.Name != "Foobar" {
		t.Errorf("struct name is not Foobar but %s", strct.Name)
	}
	for i := range strct.Fields {
		strct.Fields[i].Pos = lexer.Position{}
		strct.Fields[i].Types.Pos = lexer.Position{}
		for j := range strct.Fields[i].Types.Children {
			strct.Fields[i].Types.Children[j].Pos = lexer.Position{}
		}
	}
	expectedFields := []*Field{
		{
			Name:     "firstField",
			Required: true,
			Types: &Types{
				Children: []*Type{
					{Lang: "common", Name: "string"},
				},
			},
		},
		{
			Name:     "anotherField",
			Required: true,
			Types: &Types{
				Children: []*Type{
					{Lang: "go", Name: "int32"},
					{Lang: "ts", Name: "number"},
				},
			},
		},
		{
			Name:     "optionalField",
			Optional: true,
			Types: &Types{
				Children: []*Type{
					{Lang: "go", Name: "bool"},
					{Lang: "ts", Name: "boolean"},
				},
			},
		},
	}
	diff := deep.Equal(strct.Fields, expectedFields)
	for _, d := range diff {
		t.Error(d)
	}
}
