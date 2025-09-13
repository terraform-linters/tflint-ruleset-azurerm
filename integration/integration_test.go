package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"text/template"

	"github.com/google/go-cmp/cmp"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

type meta struct {
	Version string
}

type result struct {
	Issues []issue `json:"issues"`
	Errors []any   `json:"errors"`
}

type issue struct {
	Rule    any `json:"rule"`
	Message any `json:"message"`
	Range   any `json:"range"`
	Callers any `json:"callers"`

	// The following fields are ignored because they are added in TFLint v0.59.1.
	// We can uncomment this once the minimum supported version is v0.59.1+.
	// Fixed   any `json:"fixed"`
	// Fixable any `json:"fixable"`
}

func TestIntegration(t *testing.T) {
	cases := []struct {
		Name    string
		Command *exec.Cmd
		Dir     string
	}{
		{
			Name:    "basic",
			Command: exec.Command("tflint", "--format", "json", "--force"),
			Dir:     "basic",
		},
		{
			Name:    "tags",
			Command: exec.Command("tflint", "--format", "json", "--force"),
			Dir:     "tags",
		},
	}

	dir, _ := os.Getwd()
	defer os.Chdir(dir)

	for _, tc := range cases {
		testDir := dir + "/" + tc.Dir
		os.Chdir(testDir)

		var stdout, stderr bytes.Buffer
		tc.Command.Stdout = &stdout
		tc.Command.Stderr = &stderr
		if err := tc.Command.Run(); err != nil {
			t.Fatalf("Failed `%s`: %s, stdout=%s stderr=%s", tc.Name, err, stdout.String(), stderr.String())
		}

		ret, err := readResultFile(testDir)
		if err != nil {
			t.Fatalf("Failed `%s`: %s", tc.Name, err)
		}

		var expected result
		if err := json.Unmarshal(ret, &expected); err != nil {
			t.Fatalf("Failed `%s`: %s", tc.Name, err)
		}

		var got result
		if err := json.Unmarshal(stdout.Bytes(), &got); err != nil {
			t.Fatalf("Failed `%s`: %s", tc.Name, err)
		}

		if !cmp.Equal(got, expected) {
			t.Fatalf("Failed `%s`: diff=%s", tc.Name, cmp.Diff(expected, got))
		}
	}
}

func readResultFile(dir string) ([]byte, error) {
	resultFile := "result.json"
	if runtime.GOOS == "windows" {
		if _, err := os.Stat(filepath.Join(dir, "result_windows.json")); !os.IsNotExist(err) {
			resultFile = "result_windows.json"
		}
	}
	if _, err := os.Stat(fmt.Sprintf("%s.tmpl", resultFile)); !os.IsNotExist(err) {
		resultFile = fmt.Sprintf("%s.tmpl", resultFile)
	}

	if !strings.HasSuffix(resultFile, ".tmpl") {
		return os.ReadFile(filepath.Join(dir, resultFile))
	}

	want := new(bytes.Buffer)
	tmpl := template.Must(template.ParseFiles(filepath.Join(dir, resultFile)))
	if err := tmpl.Execute(want, meta{Version: project.Version}); err != nil {
		return nil, err
	}
	return want.Bytes(), nil
}
