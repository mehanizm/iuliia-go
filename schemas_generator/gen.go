// go:generate go run schema_generator.go && gofmt -s ../*

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mehanizm/iuliia-go"
)

const (
	dirWithSchemasConst = "../schemas"
	fileToGenerateConst = "../schemas.go"
)

func escapeQuotes(input string) string {
	return strings.Replace(input, "\"", "\\\"", -1)
}

var funcMap = template.FuncMap{
	"escape": escapeQuotes,
}

var schemaTpl = template.Must(template.New("schemaTpl").Funcs(funcMap).Parse(`
// {{ .GetName }} schema
// {{ .Desc }}
// {{ .URL }}
var {{ .GetName }} = &Schema{
	Name: "{{ .Name }}",
	Desc: "{{ .Desc }}",
	Mapping: map[string]string{
		{{- range $key, $value := .Mapping }}
		"{{ escape $key }}": "{{ escape $value }}",
		{{- end }}
	},
	PrevMapping: map[string]string{
		{{- range $key, $value := .PrevMapping }}
		"{{ escape $key }}": "{{ escape $value }}",
		{{- end }}
	},
	NextMapping: map[string]string{
		{{- range $key, $value := .NextMapping }}
		"{{ escape $key }}": "{{ escape $value }}",
		{{- end }}
	},
	EndingMapping: map[string]string{
		{{- range $key, $value := .EndingMapping }}
		"{{ escape $key }}": "{{ escape $value }}",
		{{- end }}
	},
}`))

var schemaTestTpl = template.Must(template.New("schemaTestTpl").Funcs(funcMap).Parse(`
// {{ .GetName }} schema
func Test_{{ .GetName }}(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{{- range $i, $sample := .Samples}}
		{
			name: "{{ $i }}",
			in:   "{{ index $sample 0 | escape }}",
			out:  "{{ index $sample 1 | escape }}",
		},
		{{- end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{ .GetName }}.isBuilt = false
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

var schemasTpl = template.Must(template.New("schemasTpl").Parse(`
// SchemaMapping mapping of all schemas
var SchemaMapping = map[string]*Schema{
	{{- range $key, $value := . }}
	"{{ $key }}": {{ $value }},
	{{- end }}
}`))

type schemable interface {
	String() string
	StringTest() string
	GetName() string
	GetDictName() string
}

// SchemaGen composite type for generation
type SchemaGen iuliia.Schema

// GetName get title name of the schema
func (s *SchemaGen) GetName() string {
	return strings.Title(s.Name)
}

// GetDictName get name of the schema
func (s *SchemaGen) GetDictName() string {
	return s.Name
}

// String generate code for stucts
func (s *SchemaGen) String() string {
	var res bytes.Buffer
	if err := schemaTpl.Execute(&res, s); err != nil {
		log.Fatal(err)
	}
	return res.String()
}

// StringTest generate code for tests
func (s *SchemaGen) StringTest() string {
	var res bytes.Buffer
	if err := schemaTestTpl.Execute(&res, s); err != nil {
		log.Fatal(err)
	}
	return res.String()
}

func printSchemaToBuffer(
	fileToRead string, schema schemable, schemas map[string]string,
	destinationToWriteCode, destinationToWriteTest *bytes.Buffer) error {
	schemaFileData, err := ioutil.ReadFile(fileToRead)
	if err != nil {
		return err
	}
	err = json.Unmarshal(schemaFileData, schema)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(destinationToWriteCode, schema.String())
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(destinationToWriteTest, schema.StringTest())
	if err != nil {
		return err
	}
	schemas[schema.GetDictName()] = schema.GetName()
	return nil
}

func printToFileWithFmt(buffer *bytes.Buffer, fileToWrite *os.File) error {
	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("error in go formatting: %w", err)
	}
	_, err = fmt.Fprintf(fileToWrite, "%s", formatted)
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
)`)

	var b, bt []byte
	buffer := bytes.NewBuffer(b)
	bufferTest := bytes.NewBuffer(bt)
	schemas := make(map[string]string, 0)
	for _, schemaFile := range schemaFiles {
		if !strings.HasSuffix(schemaFile.Name(), ".json") {
			continue
		}
		var err error
		var schema schemable
		fileToRead := filepath.Join(dirWithSchemas, schemaFile.Name())
		fmt.Println(fileToRead)
		schema = &SchemaGen{}
		err = printSchemaToBuffer(fileToRead, schema, schemas, buffer, bufferTest)
		if err != nil {
			log.Fatal(err)
		}
	}

	var res bytes.Buffer
	if err := schemasTpl.Execute(&res, schemas); err != nil {
		log.Panicln("error in gen template for all schemas", err)
	}
	_, err = fmt.Fprint(buffer, res.String())
	if err != nil {
		log.Panicln("error in adding code for all schemas", err)
	}
	err = printToFileWithFmt(buffer, generatedFile)
	if err != nil {
		log.Panicln("error in printing to file result", err)
	}
	err = printToFileWithFmt(bufferTest, generatedTestFile)
	if err != nil {
		log.Panicln("error in printing to file test result", err)
	}
}
