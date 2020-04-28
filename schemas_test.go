// Package iuliia do not edit, generated file
package iuliia

import (
	"fmt"
	"testing"
)

// spell-checker: disable

func Test_Wikipedia(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "1",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, syesh yeshchyo etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "10",
			in:   "Давыдов",
			out:  "Davydov",
		},
		{
			name: "11",
			in:   "Усолье",
			out:  "Usolye",
		},
		{
			name: "12",
			in:   "Выхухоль",
			out:  "Vykhukhol",
		},
		{
			name: "13",
			in:   "Дальнегорск",
			out:  "Dalnegorsk",
		},
		{
			name: "14",
			in:   "Ильинский",
			out:  "Ilyinsky",
		},
		{
			name: "15",
			in:   "Красный",
			out:  "Krasny",
		},
		{
			name: "16",
			in:   "Великий",
			out:  "Veliky",
		},
		{
			name: "17",
			in:   "Набережные Челны",
			out:  "Naberezhnye Chelny",
		},
		{
			name: "2",
			in:   "Россия, город Йошкар-Ола, улица Яна Крастыня",
			out:  "Rossiya, gorod Yoshkar-Ola, ulitsa Yana Krastynya",
		},
		{
			name: "3",
			in:   "Ельцин",
			out:  "Yeltsin",
		},
		{
			name: "4",
			in:   "Раздольное",
			out:  "Razdolnoye",
		},
		{
			name: "5",
			in:   "Юрьев",
			out:  "Yuryev",
		},
		{
			name: "6",
			in:   "Белкин",
			out:  "Belkin",
		},
		{
			name: "7",
			in:   "Бийск",
			out:  "Biysk",
		},
		{
			name: "8",
			in:   "Подъярский",
			out:  "Podyarsky",
		},
		{
			name: "9",
			in:   "Мусийкъонгийкоте",
			out:  "Musiykyongiykote",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Wikipedia.isBuilt = false
			got, err := Wikipedia.Translate(tt.in)
			if err != nil {
				t.Errorf("Wikipedia get an err:\n%v", err)
			}
			if got != tt.out {
				fmt.Println(Wikipedia)
				t.Errorf("Wikipedia got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}
