package iuliia

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// non-letter regex
var splitter = regexp.MustCompile(`\P{L}+`)

// splitSentenceRegex splits sentence by word boundaries.
// Returns words slice consisting of words and separators,
// ordered according to original sentence.
func splitSentenceRegex(source string) []string {
	// Go does not support unicode word boundaries in regexes (\b)
	// and it also does not support lookahead (?=), which can emulate boundaries
	// so we have to reinvent the wheel in such a creative way
	indexes := splitter.FindAllStringIndex(source, -1)
	idx := 0
	var words []string
	for _, boundary := range indexes {
		word := source[idx:boundary[0]]
		words = appendNonEmptyWord(words, word)
		separator := source[boundary[0]:boundary[1]]
		words = appendNonEmptyWord(words, separator)
		idx = boundary[1]
	}
	words = appendNonEmptyWord(words, source[idx:])
	return words
}

// merge one by one two arrays with order toggle
func merge(first, second []string, order bool) []string {
	res := make([]string, len(first)+len(second))
	shiftFirst := 0
	shiftSecond := 1
	if !order {
		shiftFirst = 1
		shiftSecond = 0
	}
	for i, el := range first {
		res[2*i+shiftFirst] = el
	}
	for i, el := range second {
		res[2*i+shiftSecond] = el
	}
	return res
}

// appendNonEmptyWord appends word to words slice if the word is not empty.
// Returns modified (or non-modified) words slice.
func appendNonEmptyWord(words []string, word string) []string {
	if word == "" {
		return words
	}
	return append(words, word)
}

func isCyrillic(c rune) bool {
	return unicode.Is(unicode.Cyrillic, c)
}

func notLetter(c rune) bool {
	return !unicode.IsLetter(c)
}

func isLetter(c rune) bool {
	return unicode.IsLetter(c)
}

func splitSentenceUnicode(source string) []string {
	words := strings.FieldsFunc(source, notLetter)
	notWords := strings.FieldsFunc(source, isLetter)
	isFirstLetter := isLetter([]rune(source)[0])

	return merge(words, notWords, isFirstLetter)
}

func splitSentenceFields(source string) []string {
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

	chunks = append(chunks, len(source))

	res := make([]string, len(chunks)-1)
	for i := 0; i < len(chunks)-1; i++ {
		res[i] = source[chunks[i]:chunks[i+1]]
	}
	return res
}

func splitSentence(source string) []string {
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
	if firstLetter == utf8.RuneError {
		return in
	}
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

func getPairs1(in []rune) (prev, curr, next string) {
	prev = strings.Trim(string(in[:2]), string(rune(0)))
	next = strings.Trim(string(in[1:]), string(rune(0)))
	curr = string(in[1])
	return
}

func getPairs2(in []rune) (prev, curr, next string) {
	if in[0] > rune(0) {
		prev += string(in[0])
	}
	prev += string(in[1])
	curr = string(in[1])
	next += string(in[1])
	if in[2] > rune(0) {
		next += string(in[2])
	}
	return
}
