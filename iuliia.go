// Package iuliia main
// Copyright © 2020 Mike Berezin
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.
//go:generate go run schemas_generator/gen.go schemas schemas.go && gofmt -s .
package iuliia

import (
	"fmt"
	"io"
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

// transalteEnding translates ending of the word
// return result and true if ending was translated
// return result and false if there is no ending in the schema
func (s *Schema) translateEnding(ending string) (string, bool) {
	if translated, existInEnding := s.EndingMapping[ending]; existInEnding {
		return translated, true
	}
	return ending, false
}

// translateLetters translates stem of the word
func (s *Schema) translateLetters(word string) (string, error) {
	var res strings.Builder
	wordReader := newLetterReader(word)
	for {
		prev, curr, next, isLast := wordReader.readLetters()
		_, err := res.WriteString(s.translateLetter(prev, curr, next))
		if isLast != nil && isLast != io.EOF {
			return res.String(), fmt.Errorf("error in reading words: %w", isLast)
		}
		if err != nil {
			return res.String(), fmt.Errorf("error in writing string: %w", err)
		}

		if isLast == io.EOF {
			return res.String(), nil
		}
	}
}

// translateWord split the input
// word to stem and ending
// Translate parts and combine result
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

// Translate translates input strings with schema
// return error if any of the word
// was translated with error
func (s *Schema) Translate(source string) (string, error) {
	if !s.isBuilt {
		s.build()
	}
	translated := make([]string, 0)
	for _, word := range splitSentence(source) {
		translatedWord, err := s.translateWord(word)
		if err != nil {
			return "", err
		}
		translated = append(translated, translatedWord)
	}
	return strings.Join(translated, ""), nil
}
