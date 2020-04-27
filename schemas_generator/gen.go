// go:generate go run schema_generator.go && gofmt -s ../*

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mehanizm/iuliia-go"
	"gopkg.in/yaml.v2"
)

const (
	dirWithSchemasConst = "../schemas"
	fileToGenerateConst = "../schemas.go"
)

type schemable interface {
	String() string
	GetName() string
}

func printSchemaToFile(fileToRead string, schema schemable, fileToWrite *os.File) error {
	schemaFileData, err := ioutil.ReadFile(fileToRead)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(schemaFileData, schema)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(fileToWrite, schema.String())
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
	fmt.Fprintln(generatedTestFile, `
import (
	"fmt"
	"testing"
)

// spell-checker: disable`)

	for _, schemaFile := range schemaFiles {
		var err error
		var schema schemable
		var fileToWrite *os.File
		fileToRead := filepath.Join(dirWithSchemas, schemaFile.Name())
		fmt.Println(fileToRead)
		switch {
		case strings.Contains(schemaFile.Name(), "test"):
			schema = &iuliia.SchemaTest{}
			fileToWrite = generatedTestFile
		default:
			schema = &iuliia.Schema{}
			fileToWrite = generatedFile
		}
		err = printSchemaToFile(fileToRead, schema, fileToWrite)
		if err != nil {
			log.Fatal(err)
		}
	}
}
