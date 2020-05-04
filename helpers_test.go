package iuliia

import (
	"reflect"
	"testing"
)

func Test_splitSentence(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   []string
	}{
		{
			"1", "Hello, mankind!", []string{"Hello", ", ", "mankind", "!"},
		},
		{
			"2", "Привет, человечество!", []string{"Привет", ", ", "человечество", "!"},
		},
		{
			"3", "(привет) (привет...) привет? а", []string{"(", "привет", ") (", "привет", "...) ", "привет", "? ", "а"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitSentence(tt.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_splitSentence() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_splitWord(t *testing.T) {
	tests := []struct {
		name  string
		word  string
		want  string
		want1 string
	}{
		{
			"1", "Миша", "Ми", "ша",
		},
		{
			"2", "Ми", "Ми", "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitWord(tt.word)
			if got != tt.want {
				t.Errorf("splitWord() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitWord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_letterReader_readLetters(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  [][]rune
	}{
		{
			name: "long",
			in:   "abcde",
			out: [][]rune{
				{rune(0), 'a', 'b'},
				{'a', 'b', 'c'},
				{'b', 'c', 'd'},
				{'c', 'd', 'e'},
				{'d', 'e', rune(0)},
			},
		},
		{
			name: "short",
			in:   "a",
			out: [][]rune{
				{rune(0), 'a', rune(0)},
			},
		},
		{
			name: "cyrillic",
			in:   "съешь",
			out: [][]rune{
				{rune(0), 'с', 'ъ'},
				{'с', 'ъ', 'е'},
				{'ъ', 'е', 'ш'},
				{'е', 'ш', 'ь'},
				{'ш', 'ь', rune(0)},
			},
		},
		{
			name: "empty",
			in:   "",
			out: [][]rune{
				{rune(0), rune(0), rune(0)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := newLetterReader(tt.in)
			i := 0
			for {
				gotPrev, gotCurr, gotNext, gotErr := lr.readLetters()
				expPrev := tt.out[i][0]
				expCurr := tt.out[i][1]
				expNext := tt.out[i][2]
				i++
				expErr := len(tt.out) == i
				if (gotErr != nil) != expErr {
					t.Errorf("letterReader.readLetters() error = %v, wantErr %v", gotErr, expErr)
					return
				}
				if gotPrev != expPrev {
					t.Errorf("letterReader.readLetters() got = %v, want %v", gotPrev, expPrev)
				}
				if gotCurr != expCurr {
					t.Errorf("letterReader.readLetters() got1 = %v, want %v", gotCurr, expPrev)
				}
				if gotNext != expNext {
					t.Errorf("letterReader.readLetters() got2 = %v, want %v", gotNext, expPrev)
				}
				if gotErr != nil {
					return
				}
			}

		})
	}
}

func TestSchemaPrinter(t *testing.T) {
	tests := []struct {
		name    string
		schemas map[string]*Schema
		want    string
	}{
		{
			name: "simple positive",
			schemas: map[string]*Schema{
				"test schema":   {Name: "schema name", Desc: "schema desc"},
				"a test schema": {Name: "a test schema", Desc: "schema desc"},
			},
			want: "a test schema:       schema desc\nschema name:         schema desc\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SchemaPrinter(tt.schemas); got != tt.want {
				t.Errorf("\nSchemaPrinter():\n%v\nwant:\n%v", got, tt.want)
			}
		})
	}
}
