package iuliia

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func isLetter(c rune) bool {
	return unicode.IsLetter(c)
}

// splitSentence splits the sentences
// by only the words and non-words correctly
func splitSentence(source string) []string {
	// first element "0" already in the slice
	chunks := make([]int, 1, 32)

	wasLetter := false
	for i, rune := range source {
		switch {
		case isLetter(rune) && !wasLetter:
			wasLetter = true
			if i == 0 {
				continue
			}
			chunks = append(chunks, i)
		case !isLetter(rune) && wasLetter:
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

// getPairs convert rune triplet to two string pairs
func getPairs(prev, curr, next rune) (prevPair string, nextPair string) {
	if prev != rune(0) {
		prevPair += string(prev)
	}
	prevPair += string(curr)
	nextPair += string(curr)
	if next != rune(0) {
		nextPair += string(next)
	}
	return
}

// letterReader struct to read rune triplets
type letterReader struct {
	wordInRune       []rune
	index            int
	size             int
	prev, curr, next string
}

// letterReader constructor
func newLetterReader(in string) *letterReader {
	return &letterReader{
		wordInRune: []rune(in),
		index:      0,
		size:       utf8.RuneCountInString(in),
	}
}

// readLetters reads rune triplets from string
// return io.IOF if string is end
func (lr *letterReader) readLetters() (rune, rune, rune, error) {
	if lr.size == 0 {
		return rune(0), rune(0), rune(0), io.EOF
	}
	if lr.size == 1 {
		return rune(0), lr.wordInRune[0], rune(0), io.EOF
	}
	if lr.index == lr.size-1 {
		return lr.wordInRune[lr.size-2], lr.wordInRune[lr.size-1], rune(0), io.EOF
	}
	if lr.index == 0 {
		lr.index++
		return rune(0), lr.wordInRune[0], lr.wordInRune[1], nil
	}
	lr.index++
	return lr.wordInRune[lr.index-2], lr.wordInRune[lr.index-1], lr.wordInRune[lr.index], nil
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
