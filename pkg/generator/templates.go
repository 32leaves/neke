package generator

import (
	"github.com/iancoleman/strcase"
)

var Languages = map[string]*LanguageImpl{
	"go":         &GoLang,
	"ts":         &Typescript,
	"ts-promise": &TypescriptWithPromise,
}

type LanguageImpl struct {
	name  string
	iface string
	strct string
	enum  string
	namer func(string) string
}

var GoLang = LanguageImpl{
	name: "go",
	iface: `type {{ name .Name }} interface {
    {{ range .Entry -}}
        {{ name .Name }}(req *{{ type .Request }}) (*{{ type .Response }}, error)
    {{ end }}
}`,
	strct: `type {{ name .Name }} struct {
    {{ range .Fields -}}
        {{ name .Name }} {{ type .Types }} ` + "`" + `json:"{{ jsonName .Name }}{{- if .Optional -}},omitempty{{- end -}}"` + "`" + `
    {{ end }}
}`,
	enum: `type {{ name .Name }} string
const (
    {{- $root := . -}}
    {{ range $idx, $val := .Values -}}
    {{ name $root.Name }}_{{ name $val }} {{ name $root.Name }} = "{{ jsonValue $val }}"
    {{ end }}
)`,
	namer: strcase.ToCamel,
}

var Typescript = LanguageImpl{
	name: "ts",
	iface: `export interface {{ name .Name }} {
    {{ range .Entry -}}
        {{ name .Name }}(req {{ type .Request }}): {{ type .Response }};
    {{ end -}}
}`,
	strct: `export interface {{ name .Name }} {
    {{ range .Fields -}}
        {{ jsonName .Name }}{{- if .Optional -}}?{{- end -}}: {{ type .Types }};
    {{ end -}}
}`,
	enum: `export enum {{ name .Name }} {
    {{ range .Values -}}
        {{ name . }} = "{{ jsonValue . }}";
    {{ end -}}
}
    `,
	namer: strcase.ToCamel,
}
var TypescriptWithPromise = LanguageImpl{
	name: Typescript.name,
	iface: `export interface {{ name .Name }} {
    {{ range .Entry -}}
        {{ name .Name }}(req {{ type .Request }}): Promise<{{ type .Response }}>;
    {{ end -}}
}`,
	strct: Typescript.strct,
	enum:  Typescript.enum,
	namer: Typescript.namer,
}
