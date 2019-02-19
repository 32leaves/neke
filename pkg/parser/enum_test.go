package parser

import (
	"strings"
	"testing"
)

func TestBasicEnum(t *testing.T) {
	code := `
enum Foobar {
    Hello
    World
    AnotherItem
}
`
	r := &AST{}
	err := NewParser().Parse(strings.NewReader(code), r)
	if err != nil {
		t.Error(err)
		return
	}

	if len(r.Entries) == 0 {
		t.Errorf("did not find any entry")
		return
	}
	re := r.Entries[0]
	if re == nil || re.Enum == nil {
		t.Errorf("first entry is not an enum")
		return
	}
	if re.Enum.Name != "Foobar" {
		t.Errorf("enum name is not Foobar but %s", re.Enum.Name)
	}
	if re.Enum.Values == nil || len(re.Enum.Values) != 3 {
		t.Errorf("enum does not have the correct number of values")
		return
	}
	for i, v := range []string{"Hello", "World", "AnotherItem"} {
		if re.Enum.Values[i] != v {
			t.Errorf("enum value mismatch: expected %s, actual %s", v, re.Enum.Values[i])
		}
	}
}
