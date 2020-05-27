// Package iuliia main
// Copyright © 2020 Mike Berezin
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.
//go:generate go run schemas_generator/gen.go schemas schemas.go && gofmt -s .
package iuliia

import (
	"strings"
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

// Schema base Schema struct
type Schema struct {
	Name          string            `json:"name"`
	Desc          string            `json:"description"`
	URL           string            `json:"url"`
	Mapping       map[string]string `json:"mapping"`
	PrevMapping   map[string]string `json:"prev_mapping"`
	NextMapping   map[string]string `json:"next_mapping"`
	EndingMapping map[string]string `json:"ending_mapping"`
	Samples       [][]string        `json:"samples"`
	isBuilt       bool
}

// build initiate schema from raw source
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

// translateLetter translates one letter
// with respect to neighbors
func (s *Schema) translateLetter(res *strings.Builder, in []rune) {
	prev, curr, next := getPairs1(in)
	if translated, existInPrev := s.PrevMapping[prev]; existInPrev {
		res.WriteString(translated)
		return
	}
	if translated, existInNext := s.NextMapping[next]; existInNext {
		res.WriteString(translated)
		return
	}
	if translated, existInCurr := s.Mapping[curr]; existInCurr {
		res.WriteString(translated)
		return
	}
	res.WriteString(string(in[1]))
	return
}

// transalteEnding translates ending of the word
// return result and true if ending was translated
// return result and false if there is no ending in the schema
func (s *Schema) translateEnding(res *strings.Builder, ending string) bool {
	if translated, existInEnding := s.EndingMapping[ending]; existInEnding {
		res.WriteString(translated)
		return true
	}
	res.WriteString(ending)
	return false
}

func (s *Schema) endingExist(ending string) bool {
	_, existInEnding := s.EndingMapping[ending]
	return existInEnding
}

// translateLetters translates stem of the word
func (s *Schema) translateLetters(res *strings.Builder, word string) {
	for _, letters := range readLetters(word) {
		s.translateLetter(res, letters)
	}
}

// translateWord split the input
// word to stem and ending
// Translate parts and combine result
func (s *Schema) translateWord(res *strings.Builder, word string) {
	if word == "" {
		return
	}
	if !isCyrillic([]rune(word)[0]) {
		res.WriteString(word)
		return
	}
	stem, ending := splitWord(word)
	if s.endingExist(ending) {
		s.translateLetters(res, stem)
		s.translateEnding(res, ending)
		return
	}
	s.translateLetters(res, word)
	return
}

// Translate translates input strings with schema
// return error if any of the word
// was translated with error
func (s *Schema) Translate(source string) string {
	if !s.isBuilt {
		s.build()
	}
	var res strings.Builder
	for _, word := range splitSentenceUnicode(source) {
		s.translateWord(&res, word)
	}
	return res.String()
}
