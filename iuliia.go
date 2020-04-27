package iuliia

import (
	"bytes"
	"io"
	"log"
	"strings"
	"text/template"
)

var baseMapping = map[string]string{
	"а": "a",
	"б": "b",
	"в": "v",
	"г": "g",
	"д": "d",
	"е": "e",
	"з": "z",
	"и": "i",
	"к": "k",
	"л": "l",
	"м": "m",
	"н": "n",
	"о": "o",
	"п": "p",
	"р": "r",
	"с": "s",
	"т": "t",
	"у": "u",
	"ф": "f",
}

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
}
`))

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
}
`))

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

// Schema base Schema struct
type Schema struct {
	Name          string            `yaml:"schema_name"`
	Mapping       map[string]string `yaml:"mapping"`
	PrevMapping   map[string]string `yaml:"prev_mapping"`
	NextMapping   map[string]string `yaml:"next_mapping"`
	EndingMapping map[string]string `yaml:"ending_mapping"`
	isBuilt       bool
}

func (s *Schema) build() *Schema {
	mapping := make(map[string]string)
	for key, value := range baseMapping {
		mapping[key] = value
		mapping[capitalize(key)] = capitalize(value)
	}
	for key, value := range s.Mapping {
		mapping[key] = value
		mapping[capitalize(key)] = capitalize(value)
	}
	s.Mapping = mapping

	for key, value := range s.PrevMapping {
		s.PrevMapping[capitalize(key)] = value
		s.PrevMapping[strings.ToUpper(key)] = capitalize(value)
	}

	for key, value := range s.NextMapping {
		s.NextMapping[strings.ToUpper(key)] = capitalize(value)
		s.NextMapping[capitalize(key)] = capitalize(value)
	}

	for key, value := range s.EndingMapping {
		s.EndingMapping[strings.ToUpper(key)] = strings.ToUpper(value)
	}

	s.isBuilt = true

	return s
}

// GetName get title name of the schema
func (s *Schema) GetName() string {
	return strings.Title(s.Name)
}

func (s *Schema) String() string {
	var res bytes.Buffer
	if err := schemaTpl.Execute(&res, s); err != nil {
		log.Fatal(err)
	}
	return res.String()
}

func (s *Schema) translateLetter(prev, curr, next rune) string {
	prevPair, nextPair := getPairs(prev, curr, next)
	if translated, existInPrev := s.PrevMapping[prevPair]; existInPrev {
		return translated
	}
	if translated, existInNext := s.NextMapping[nextPair]; existInNext {
		return translated
	}
	if translated, existInCurr := s.Mapping[string(curr)]; existInCurr {
		return translated
	}
	return string(curr)
}

func (s *Schema) translateEnding(ending string) (string, bool) {
	if translated, existInEnding := s.EndingMapping[ending]; existInEnding {
		return translated, true
	}
	return ending, false
}

func (s *Schema) translateLetters(word string) (string, error) {
	var res strings.Builder
	wordReader := newLetterReader(word)
	for {
		prev, curr, next, isLast := wordReader.ReadLetters()
		_, err := res.WriteString(s.translateLetter(prev, curr, next))
		if err != nil || (isLast != nil && isLast != io.EOF) {
			return res.String(), err
		}
		if isLast == io.EOF {
			return res.String(), nil
		}
	}
}

func (s *Schema) translateWord(word string) (string, error) {
	stem, ending := splitWord(word)
	translatedEnding, isEndingTranslated := s.translateEnding(ending)
	if isEndingTranslated {
		translatedStem, err := s.translateLetters(stem)
		if err != nil {
			return "", err
		}
		return translatedStem + translatedEnding, nil
	}
	translatedWord, err := s.translateLetters(word)
	if err != nil {
		return "", err
	}
	return translatedWord, nil
}

// Translate translates due to schema
func (s *Schema) Translate(source string) (string, error) {
	if !s.isBuilt {
		s.build()
	}
	translated := make([]string, 0)
	for _, word := range strings.Split(source, " ") {
		translatedWord, err := s.translateWord(word)
		if err != nil {
			return "", err
		}
		translated = append(translated, translatedWord)
	}
	return strings.Join(translated, " "), nil
}
