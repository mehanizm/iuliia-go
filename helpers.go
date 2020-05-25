package iuliia

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func isCyrillic(c rune) bool {
	return unicode.Is(unicode.Cyrillic, c)
}

// splitSentence splits the sentences
// by only the words and non-words correctly
func splitSentence(source string) []string {
	// first element "0" already in the slice
	chunks := make([]int, 1, 32)

	wasLetter := false
	for i, rune := range source {
		switch {
		case isCyrillic(rune) && !wasLetter:
			wasLetter = true
			if i == 0 {
				continue
			}
			chunks = append(chunks, i)
		case !isCyrillic(rune) && wasLetter:
			wasLetter = false
			chunks = append(chunks, i)
		default:
			continue
		}
	}

	// add last element
	chunks = append(chunks, len(source))

	res := make([]string, len(chunks)-1)
	for i := 0; i < len(chunks)-1; i++ {
		res[i] = source[chunks[i]:chunks[i+1]]
	}
	return res
}

// splitWord to stem and ending
func splitWord(word string) (string, string) {
	endingLength := 2
	lenOfWord := utf8.RuneCountInString(word)
	if lenOfWord <= endingLength {
		return word, ""
	}
	wordInRune := []rune(word)
	return string(wordInRune[:lenOfWord-endingLength]),
		string(wordInRune[lenOfWord-endingLength:])
}

// capitalize uppers only first letter of the string
func capitalize(in string) string {
	firstLetter, size := utf8.DecodeRuneInString(in)
	return string(unicode.ToUpper(firstLetter)) + in[size:]
}

// sliding window 3 size on word
func readLetters(in string) [][]rune {
	if in == "" {
		return [][]rune{make([]rune, 3)}
	}
	inRune := append([]rune(in), make([]rune, 2)...)
	copy(inRune[1:], inRune[:])
	inRune[0] = rune(0)
	res := make([][]rune, 0, len(inRune)-2)
	for i := 0; i < len(inRune)-2; i++ {
		res = append(res, inRune[i:i+3])
	}
	return res
}

// SchemaPrinter prints schemas line by line
func SchemaPrinter(schemas map[string]*Schema) string {
	res := strings.Builder{}
	sortedKeys := make([]string, 0, len(schemas))
	for k := range schemas {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	for _, schemaName := range sortedKeys {
		schema := schemas[schemaName]
		num := 20 - utf8.RuneCountInString(schema.Name)
		res.WriteString(
			fmt.Sprintf(
				"%s:%s%s\n",
				schema.Name,
				strings.Repeat(" ", num),
				schema.Desc,
			),
		)

	}
	return res.String()
}
