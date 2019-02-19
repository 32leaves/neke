package generator

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/32leaves/neke/pkg/parser"
	"github.com/iancoleman/strcase"
)

type Generator interface {
	Render(ast *parser.AST, out io.Writer) error
}

func (l *LanguageImpl) Render(ast *parser.AST, out io.Writer) error {
	iface, err := l.getTemplate("interface", l.iface)
	if err != nil {
		return err
	}
	strct, err := l.getTemplate("struct", l.strct)
	if err != nil {
		return err
	}
	enum, err := l.getTemplate("enum", l.enum)
	if err != nil {
		return err
	}

	for _, entry := range ast.Entries {
		if entry.Interface != nil {
			if err := iface.Execute(out, entry.Interface); err != nil {
				return err
			}
		} else if entry.Struct != nil {
			if err := strct.Execute(out, entry.Struct); err != nil {
				return err
			}
		} else if entry.Enum != nil {
			if err := enum.Execute(out, entry.Enum); err != nil {
				return err
			}
		} else {
			continue
		}

		out.Write([]byte("\n\n"))
	}

	return nil
}

func (l *LanguageImpl) getTemplate(nme, tpl string) (*template.Template, error) {
	funcs := make(map[string]interface{})
	funcs["name"] = l.namer
	funcs["type"] = l.getType
	funcs["jsonName"] = strcase.ToLowerCamel
	funcs["jsonValue"] = strcase.ToLowerCamel
	funcs["toUpper"] = strings.ToUpper

	return template.
		New(nme).
		Funcs(funcs).
		Parse(tpl)
}

func (l *LanguageImpl) getType(ts *parser.Types) (string, error) {
	t := ts.ByLang(l.name)
	if t == "" {
		return "", fmt.Errorf("no type for %s", l.name)
	}
	return t, nil
}
