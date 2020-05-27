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
			"1", "Hello, mankind!", []string{"Hello, mankind!"},
		},
		{
			"2", "Привет человечество!", []string{"Привет", " ", "человечество", "!"},
		},
		{
			"3", "(привет) (привет...) привет? а", []string{"(", "привет", ") (", "привет", "...) ", "привет", "? ", "а"},
		},
	}
	funcs := []struct {
		name string
		fun  func(source string) []string
	}{
		{"regex", splitSentence},
	}
	for _, tt := range tests {
		for _, fun := range funcs {
			t.Run(tt.name, func(t *testing.T) {
				got := fun.fun(tt.source)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Test_splitSentence(), fun %s\ngot = %#v, want %#v", fun.name, got, tt.want)
				}
			})
		}
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
		{
			"3", "Ё", "Ё", "",
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
			res := readLetters(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("wrong result: want: %v was: %v", tt.out, res)
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

func Test_getPairs(t *testing.T) {
	tests := []struct {
		name     string
		in       []rune
		wantPrev string
		wantCurr string
		wantNext string
	}{
		{
			name:     "all letters",
			in:       []rune{'a', 'b', 'c'},
			wantPrev: "ab",
			wantCurr: "b",
			wantNext: "bc",
		},
		{
			name:     "only middle",
			in:       []rune{rune(0), 'b', rune(0)},
			wantPrev: "b",
			wantCurr: "b",
			wantNext: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrev, gotCurr, gotNext := getPairs(tt.in)
			if gotPrev != tt.wantPrev {
				t.Errorf("getPairs() gotPrev = %v, want %v", gotPrev, tt.wantPrev)
			}
			if gotCurr != tt.wantCurr {
				t.Errorf("getPairs() gotCurr = %v, want %v", gotCurr, tt.wantCurr)
			}
			if gotNext != tt.wantNext {
				t.Errorf("getPairs() gotNext = %v, want %v", gotNext, tt.wantNext)
			}
		})
	}
}
