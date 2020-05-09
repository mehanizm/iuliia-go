// go build -o iuliia
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	iuliia "github.com/mehanizm/iuliia-go"
)

var (
	showShemas bool
	schemaName string
)

func init() {
	// Parse flags
	flag.BoolVar(&showShemas, "show", false, "showing list of the schemas")
	flag.StringVar(&schemaName, "schema", "wikipedia", "choose schema name")
	flag.Parse()
}

func main() {
	if showShemas {
		fmt.Println(iuliia.SchemaPrinter(iuliia.SchemaMapping))
		return
	}
	schema, schemaExist := iuliia.SchemaMapping[schemaName]
	if !schemaExist {
		fmt.Printf("schema name %v does not exist\n", schemaName)
		fmt.Printf("to see all schemas --show\n")
		return
	}

LOOP:
	for {
		fmt.Printf("type phrase to translate with schema %v:\n'quit' to exit or 'help' to get help\n", schemaName)
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		sentence, err := buf.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("try again")
			continue LOOP
		}
		switch {
		case strings.Contains(string(sentence), "quit"):
			fmt.Println("buy!")
			return
		case strings.Contains(string(sentence), "show"):
			fmt.Println(iuliia.SchemaPrinter(iuliia.SchemaMapping))
		case strings.Contains(string(sentence), "help"):
			fmt.Println("* show - to show all schemas\n* change schema_name - to change schema\n* quit - to quit the program")
		case strings.Contains(string(sentence), "change"):
			newSchemaName := strings.TrimSuffix(strings.Split(string(sentence), " ")[1], "\n")
			newSchema, schemaExist := iuliia.SchemaMapping[newSchemaName]
			if !schemaExist {
				fmt.Printf("schema name %v does not exist", newSchemaName)
				fmt.Printf("to see all schemas type 'show'\n\n")
				continue LOOP
			}
			fmt.Printf("schema was changed to %v\n\n", newSchemaName)
			schemaName = newSchemaName
			schema = newSchema
		default:
			translated, err := schema.Translate(string(sentence))
			if err != nil {
				fmt.Println("error in translating: ", err)
				fmt.Println("try again")
			}
			fmt.Print("> ")
			fmt.Println(translated)
		}
	}
}
