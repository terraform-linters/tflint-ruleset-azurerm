package main

import (
	"bufio"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

type metadata struct {
	RuleName   string
	RuleNameCC string
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Rule name? (e.g. azurerm_linux_virtual_machine_invalid_size): ")
	ruleName, err := buf.ReadString('\n')
	if err != nil {
		panic(err)
	}
	ruleName = strings.Trim(ruleName, "\n")

	meta := &metadata{RuleNameCC: toCamelCase(ruleName), RuleName: ruleName}

	GenerateFileWithLogs(fmt.Sprintf("rules/%s.go", ruleName), "rules/rule.go.tmpl", meta)
	GenerateFileWithLogs(fmt.Sprintf("rules/%s_test.go", ruleName), "rules/rule_test.go.tmpl", meta)
	GenerateFileWithLogs(fmt.Sprintf("docs/rules/%s.md", ruleName), "docs/rules/rule.md.tmpl", meta)

	src, err := ioutil.ReadFile("rules/provider.go")
	if err != nil {
		panic(err)
	}
	dstf, err := decorator.Parse(src)
	if err != nil {
		panic(err)
	}

	dst.Inspect(dstf, func(n dst.Node) bool {
		switch node := n.(type) {
		case *dst.CompositeLit:
			expr := &dst.CallExpr{
				Fun: &dst.Ident{
					Name: fmt.Sprintf("New%sRule", meta.RuleNameCC),
				},
			}
			expr.Decs.Before = dst.NewLine
			expr.Decs.After = dst.NewLine
			node.Elts = append(node.Elts, expr)
		}
		return true
	})

	fset, astf, err := decorator.RestoreFile(dstf)
	if err != nil {
		panic(err)
	}

	fp, err := os.OpenFile("rules/provider.go", os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	if err := format.Node(fp, fset, astf); err != nil {
		panic(err)
	}
	fmt.Println("Modified: rules/provider.go")

	fmt.Println(`
TODO:
1. Remove all "TODO" comments from generated files.
2. Write implementation of the rule.
3. Add a link to the generated documentation into docs/rules/README.md`)
}

// GenerateFile generates a new file from the passed template and metadata
func GenerateFile(fileName string, tmplName string, meta interface{}) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles(tmplName))
	err = tmpl.Execute(file, meta)
	if err != nil {
		panic(err)
	}
}

// GenerateFileWithLogs generates a new file from the passed template and metadata
// The difference from GenerateFile function is to output logs
func GenerateFileWithLogs(fileName string, tmplName string, meta interface{}) {
	GenerateFile(fileName, tmplName, meta)
	fmt.Printf("Created: %s\n", fileName)
}

var heading = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

func toCamelCase(str string) string {
	exceptions := map[string]string{
		"ip":  "IP",
		"sql": "SQL",
		"vm":  "VM",
		"os":  "OS",
		"id":  "ID",
		"tls": "TLS",
	}
	parts := strings.Split(str, "_")
	replaced := make([]string, len(parts))
	for i, s := range parts {
		conv, ok := exceptions[s]
		if ok {
			replaced[i] = conv
		} else {
			replaced[i] = s
		}
	}
	str = strings.Join(replaced, "_")

	return heading.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}
