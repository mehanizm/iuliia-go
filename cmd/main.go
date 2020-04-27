// go build -o iuliia
package main

import (
	"fmt"
	"log"
	"os"

	iuliia "github.com/mehanizm/iuliia-go"
)

var schemaMapping = map[string]*iuliia.Schema{
	"wikipedia": iuliia.Wikipedia,
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("please, add schema name and string to translate")
	}
	schemaName := os.Args[1]
	textToTranslate := os.Args[2]
	schema, schemaExist := schemaMapping[schemaName]
	if !schemaExist {
		log.Fatalf("schema name %v does not exist", schemaName)
	}
	translated, err := schema.Translate(textToTranslate)
	if err != nil {
		log.Fatal("error in translating: ", err)
	}
	fmt.Println(translated)
}
