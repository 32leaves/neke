// Copyright © 2019 Christian Weichel
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/32leaves/neke/pkg/generator"
	"github.com/32leaves/neke/pkg/parser"
	"github.com/spf13/cobra"
)

var lang string
var preamble string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates code from a neke file",
	Run: func(cmd *cobra.Command, args []string) {
		var targetLanguage *generator.LanguageImpl
		var availableLanguages = make([]string, 0)
		for k, l := range generator.Languages {
			availableLanguages = append(availableLanguages, k)
			if k == lang {
				targetLanguage = l
			}
		}

		if targetLanguage == nil {
			fmt.Fprintf(os.Stderr, "missing --language flag: must be one of %s\n", strings.Join(availableLanguages, ", "))
			os.Exit(1)
		}

		os.Stdout.WriteString(preamble)
		for _, f := range args {
			if err := generate(targetLanguage, f); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}
	},
}

func generate(lang generator.Generator, fn string) error {
	r, err := os.Open(fn)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer r.Close()

	ast := &parser.AST{}
	err = parser.NewParser().Parse(r, ast)
	if err != nil {
		return fmt.Errorf("unable to parse file: %v", err)
	}

	err = lang.Render(ast, os.Stdout)
	if err != nil {
		return fmt.Errorf("error while generating: %v", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&lang, "language", "l", "", "The language for which to generate code")
	generateCmd.Flags().StringVar(&preamble, "preamble", "", "A preamble to emit before the generated code")
}
