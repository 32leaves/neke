package parser

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestParseExamples(t *testing.T) {
	exampledir := "../../examples"
	files, err := ioutil.ReadDir(exampledir)
	if err != nil {
		t.Error(err)
		return
	}

	for _, file := range files {
		fn := path.Join(exampledir, file.Name())
		r, err := os.Open(fn)
		if err != nil {
			t.Errorf("unable to open file: %s", file.Name())
			continue
		}
		defer r.Close()

		neke := &AST{}
		err = NewParser().Parse(r, neke)
		if err != nil {
			t.Errorf("parse error in %s: %v", file.Name(), err)
			continue
		}
	}
}
