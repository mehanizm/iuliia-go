package iuliia

import (
	"strings"
	"testing"
)

func TestSchema_translateWord(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		out     string
		wantErr bool
	}{
		{
			name: "1",
			in:   "Миша",
			out:  "Misha",
		},
		{
			name: "2",
			in:   "съешь",
			out:  "syesh",
		},
		{
			name: "3",
			in:   "ещё",
			out:  "yeshchyo",
		},
		{
			name: "with ending",
			in:   "старый",
			out:  "stary",
		},
		{
			name: "special",
			in:   "Ё",
			out:  "Yo",
		},
		{
			name: "empty",
			in:   "",
			out:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res strings.Builder
			Wikipedia.build().translateWord(&res, tt.in)
			got := res.String()
			if got != tt.out {
				t.Errorf("Schema.translateWord() = %v, want %v", got, tt.out)
			}
		})
	}
}

func TestSchema_translateLetter(t *testing.T) {
	tests := []struct {
		name string
		args []rune
		want string
	}{
		{
			"current",
			[]rune{'ъ', 'е', 'ш'},
			"ye",
		},
		{
			"prev mapping",
			[]rune{rune(0), 'е', 'щ'},
			"ye",
		},
		{
			"not cyrillic",
			[]rune{rune(0), '😁', rune(0)},
			"😁",
		},
		{
			"next mapping",
			[]rune{rune(0), 'ь', 'а'},
			"y",
		},
		{
			"special",
			[]rune{rune(0), 'Ё', rune(0)},
			"Yo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res strings.Builder
			Wikipedia.build().translateLetter(&res, tt.args)
			got := res.String()
			if got != tt.want {
				t.Errorf("Schema.translateLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_translateEnding(t *testing.T) {
	tests := []struct {
		name     string
		ending   string
		want     string
		wantBool bool
	}{
		{
			"translated",
			"ый",
			"y",
			true,
		},
		{
			"translated",
			"уй",
			"уй",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res strings.Builder
			Wikipedia.build().translateEnding(&res, tt.ending)
			got := res.String()
			if got != tt.want {
				t.Errorf("Schema.translateEnding() got = %v, want %v", got, tt.want)
			}
		})
	}
}
