package iuliia

import (
	"testing"
)

func Test_Schema_translateLetter(t *testing.T) {
	type args struct {
		prev rune
		curr rune
		next rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{'М', 'и', 'ш'},
			"i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wikipedia.build().translateLetter(tt.args.prev, tt.args.curr, tt.args.next); got != tt.want {
				t.Errorf("Schema.translateLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_translateWord(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		out     string
		wantErr bool
	}{
		{
			name:    "1",
			in:      "Миша",
			out:     "Misha",
			wantErr: false,
		},
		{
			name:    "2",
			in:      "съешь",
			out:     "syesh",
			wantErr: false,
		},
		{
			name:    "3",
			in:      "ещё",
			out:     "yeshchyo",
			wantErr: false,
		},
		{
			name:    "with ending",
			in:      "старый",
			out:     "stary",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Wikipedia.build().translateWord(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Schema.translateWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.out {
				t.Errorf("Schema.translateWord() = %v, want %v", got, tt.out)
			}
		})
	}
}

func TestSchema_translateLetter(t *testing.T) {
	type args struct {
		prev rune
		curr rune
		next rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"current",
			args{'ъ', 'е', 'ш'},
			"ye",
		},
		{
			"prev mapping",
			args{rune(0), 'е', 'щ'},
			"ye",
		},
		{
			"not cyrillic",
			args{rune(0), '😁', rune(0)},
			"😁",
		},
		{
			"next mapping",
			args{rune(0), 'ь', 'а'},
			"y",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wikipedia.build().translateLetter(tt.args.prev, tt.args.curr, tt.args.next); got != tt.want {
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
			got, wantBool := Wikipedia.build().translateEnding(tt.ending)
			if got != tt.want {
				t.Errorf("Schema.translateEnding() got = %v, want %v", got, tt.want)
			}
			if wantBool != tt.wantBool {
				t.Errorf("Schema.translateEnding() got1 = %v, want %v", wantBool, tt.wantBool)
			}
		})
	}
}
