// go:generate go run schema_generator.go && gofmt -s ../*

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mehanizm/iuliia-go"
	"gopkg.in/yaml.v2"
)

const (
	dirWithSchemasConst = "../schemas"
	fileToGenerateConst = "../schemas.go"
)

var schemaTpl = template.Must(template.New("schemaTpl").Parse(`
// {{.GetName}} schema
var {{.GetName}} = &Schema{
	Name: "{{.Name}}",
	Mapping: map[string]string{
		{{- range $key, $value := .Mapping }}
		"{{ $key }}": "{{ $value }}",
		{{- end }}
	},
	PrevMapping: map[string]string{
		{{- range $key, $value := .PrevMapping }}
		"{{ $key }}": "{{ $value }}",
		{{- end }}
	},
	NextMapping: map[string]string{
		{{- range $key, $value := .NextMapping }}
		"{{ $key }}": "{{ $value }}",
		{{- end }}
	},
	EndingMapping: map[string]string{
		{{- range $key, $value := .EndingMapping }}
		"{{ $key }}": "{{ $value }}",
		{{- end }}
	},
}`))

var schemaTestTpl = template.Must(template.New("schemaTestTpl").Parse(`
func Test_{{ .GetName }}(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{{- range $key, $value := .TestCases}}
		{
			name: "{{ $key }}",
			in:   "{{ index $value 0 }}",
			out:  "{{ index $value 1 }}",
		},
		{{- end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := {{ .GetName }}.Translate(tt.in)
			if err != nil {
				t.Errorf("{{ .GetName }} get an err:\n%v", err)
			}
			if got != tt.out {
				fmt.Println({{ .GetName }})
				t.Errorf("{{ .GetName }} got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}`))

type schemable interface {
	String() string
	GetName() string
}

// SchemaTest struct of the schema test case
type SchemaTest struct {
	Name      string               `yaml:"schema_name"`
	TestCases map[string][2]string `yaml:"test_cases"`
}

// GetName get title name of the schema
func (s *SchemaTest) GetName() string {
	return strings.Title(s.Name)
}

func (s *SchemaTest) String() string {
	var res bytes.Buffer
	if err := schemaTestTpl.Execute(&res, s); err != nil {
		log.Fatal(err)
	}
	return res.String()
}

// SchemaMain composite to implement interface
type SchemaMain iuliia.Schema

// GetName get title name of the schema
func (s *SchemaMain) GetName() string {
	return strings.Title(s.Name)
}

func (s *SchemaMain) String() string {
	var res bytes.Buffer
	if err := schemaTpl.Execute(&res, s); err != nil {
		log.Fatal(err)
	}
	return res.String()
}

func printSchemaToBuffer(fileToRead string, schema schemable, destinationToWrite *bytes.Buffer) error {
	schemaFileData, err := ioutil.ReadFile(fileToRead)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(schemaFileData, schema)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(destinationToWrite, schema.String())
	if err != nil {
		return err
	}
	return nil
}

func printToFileWithFmt(buffer *bytes.Buffer, fileToWrite *os.File) error {
	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(fileToWrite, string(formatted))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	dirWithSchemas := ""
	if len(os.Args) >= 2 {
		dirWithSchemas = os.Args[1]
	}
	if dirWithSchemas == "" {
		dirWithSchemas = dirWithSchemasConst
	}
	fileToGenerate := ""
	if len(os.Args) >= 3 {
		fileToGenerate = os.Args[2]
	}
	if fileToGenerate == "" {
		fileToGenerate = fileToGenerateConst
	}
	fileTestToGenerate := strings.Replace(fileToGenerate, ".go", "_test.go", -1)

	schemaFiles, err := ioutil.ReadDir(dirWithSchemas)
	if err != nil {
		log.Fatal(err)
	}
	generatedFile, err := os.Create(fileToGenerate)
	if err != nil {
		log.Fatal(err)
	}
	generatedTestFile, err := os.Create(fileTestToGenerate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(generatedFile, "// Package iuliia do not edit, generated file")
	fmt.Fprintln(generatedFile, "package iuliia")

	fmt.Fprintln(generatedTestFile, "// Package iuliia do not edit, generated file")
	fmt.Fprintln(generatedTestFile, "package iuliia")
	fmt.Fprintln(generatedTestFile) // empty line
	fmt.Fprintln(generatedTestFile, `import (
	"fmt"
	"testing"
)

// spell-checker: disable`)

	var b, bt []byte
	buffer := bytes.NewBuffer(b)
	bufferTest := bytes.NewBuffer(bt)
	for _, schemaFile := range schemaFiles {
		var err error
		var schema schemable
		var destinationToWrite *bytes.Buffer
		fileToRead := filepath.Join(dirWithSchemas, schemaFile.Name())
		fmt.Println(fileToRead)
		switch {
		case strings.Contains(schemaFile.Name(), "test"):
			schema = &SchemaTest{}
			destinationToWrite = bufferTest
		default:
			schema = &SchemaMain{}
			destinationToWrite = buffer
		}
		err = printSchemaToBuffer(fileToRead, schema, destinationToWrite)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = printToFileWithFmt(buffer, generatedFile)
	if err != nil {
		log.Fatal(err)
	}
	err = printToFileWithFmt(bufferTest, generatedTestFile)
	if err != nil {
		log.Fatal(err)
	}
}
