package parser

import (
	"strings"
	"testing"
)

func TestBasicInterface(t *testing.T) {
	code := `
interface Foobar {
    func Foo (common:FooReq) returns (common:FooResp)
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
	if re == nil || re.Interface == nil {
		t.Errorf("first entry is not an interface")
		return
	}
	if re.Interface.Name != "Foobar" {
		t.Errorf("interface name is not Foobar but %s", re.Interface.Name)
	}
	if re.Interface.Entry == nil || len(re.Interface.Entry) == 0 {
		t.Errorf("interface has no functions")
		return
	}
	if re.Interface.Entry[0].Name != "Foo" {
		t.Errorf("function name is not Foo but %s", re.Interface.Entry[0].Name)
	}
	if re.Interface.Entry[0].Request == nil {
		t.Errorf("function request has no type")
	}
	if re.Interface.Entry[0].Request.ByLang("common") != "FooReq" {
		t.Errorf("function request is not FooReq but %s", re.Interface.Entry[0].Request.ByLang("common"))
	}
	if re.Interface.Entry[0].Response.ByLang("common") != "FooResp" {
		t.Errorf("function response is not FooResp but %s", re.Interface.Entry[0].Response.ByLang("common"))
	}
}
