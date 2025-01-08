// Package iuliia do not edit, generated file
package iuliia

import (
	"fmt"
	"testing"
)

// Ala_lc schema
func Test_Ala_lc(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "I͡ulii͡a, sʺeshʹ eshchё ėtikh mi͡agkikh frant͡suzskikh bulok iz Ĭoshkar-Oly, da vypeĭ altaĭskogo chai͡u",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ala_lc.isBuilt = false
			got := Ala_lc.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Ala_lc)
				t.Errorf("Ala_lc got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Ala_lc_alt schema
func Test_Ala_lc_alt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Iuliia, s\"esh' eshche etikh miagkikh frantsuzskikh bulok iz Ioshkar-Oly, da vypei altaiskogo chaiu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ala_lc_alt.isBuilt = false
			got := Ala_lc_alt.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Ala_lc_alt)
				t.Errorf("Ala_lc_alt got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Bgn_pcgn schema
func Test_Bgn_pcgn(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, s”yesh’ yeshchё etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "1",
			in:   "Россия, город Йошкар-Ола, улица Яна Крастыня",
			out:  "Rossiya, gorod Yoshkar-Ola, ulitsa Yana Krastynya",
		},
		{
			name: "2",
			in:   "Елизово",
			out:  "Yelizovo",
		},
		{
			name: "3",
			in:   "Чапаевск",
			out:  "Chapayevsk",
		},
		{
			name: "4",
			in:   "Мейеровка",
			out:  "Meyyerovka",
		},
		{
			name: "5",
			in:   "Юрьев объезд",
			out:  "Yur’yev ob”yezd",
		},
		{
			name: "6",
			in:   "Белкино",
			out:  "Belkino",
		},
		{
			name: "7",
			in:   "Ёдва",
			out:  "Yёdva",
		},
		{
			name: "8",
			in:   "Змииёвка",
			out:  "Zmiiyёvka",
		},
		{
			name: "9",
			in:   "Айёган",
			out:  "Ayyёgan",
		},
		{
			name: "10",
			in:   "Воробьёво",
			out:  "Vorob’yёvo",
		},
		{
			name: "11",
			in:   "Кебанъёль",
			out:  "Keban”yёl’",
		},
		{
			name: "12",
			in:   "Озёрный",
			out:  "Ozёrnyy",
		},
		{
			name: "13",
			in:   "Тыайа",
			out:  "Ty·ay·a",
		},
		{
			name: "14",
			in:   "Сайылык",
			out:  "Say·ylyk",
		},
		{
			name: "15",
			in:   "Ойусардах",
			out:  "Oy·usardakh",
		},
		{
			name: "16",
			in:   "Йошкар-Ола",
			out:  "Yoshkar-Ola",
		},
		{
			name: "17",
			in:   "Бийск",
			out:  "Biysk",
		},
		{
			name: "18",
			in:   "Тыэкан",
			out:  "Ty·ekan",
		},
		{
			name: "19",
			in:   "Суык-Су",
			out:  "Su·yk-Su",
		},
		{
			name: "20",
			in:   "Тында",
			out:  "Tynda",
		},
		{
			name: "21",
			in:   "Улан-Удэ",
			out:  "Ulan-Ud·e",
		},
		{
			name: "22",
			in:   "Электрогорск",
			out:  "Elektrogorsk",
		},
		{
			name: "23",
			in:   "Руэм",
			out:  "Ruem",
		},
		{
			name: "24",
			in:   "Вяртсиля",
			out:  "Vyart·silya",
		},
		{
			name: "25",
			in:   "Ташчишма",
			out:  "Tash·chishma",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bgn_pcgn.isBuilt = false
			got := Bgn_pcgn.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Bgn_pcgn)
				t.Errorf("Bgn_pcgn got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Bgn_pcgn_alt schema
func Test_Bgn_pcgn_alt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, s”yesh’ yeshchё etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "1",
			in:   "Россия, город Йошкар-Ола, улица Яна Крастыня",
			out:  "Rossiya, gorod Yoshkar-Ola, ulitsa Yana Krastynya",
		},
		{
			name: "2",
			in:   "Елизово",
			out:  "Yelizovo",
		},
		{
			name: "3",
			in:   "Чапаевск",
			out:  "Chapayevsk",
		},
		{
			name: "4",
			in:   "Мейеровка",
			out:  "Meyyerovka",
		},
		{
			name: "5",
			in:   "Юрьев объезд",
			out:  "Yur’yev ob”yezd",
		},
		{
			name: "6",
			in:   "Белкино",
			out:  "Belkino",
		},
		{
			name: "7",
			in:   "Ёдва",
			out:  "Yёdva",
		},
		{
			name: "8",
			in:   "Змииёвка",
			out:  "Zmiiyёvka",
		},
		{
			name: "9",
			in:   "Айёган",
			out:  "Ayyёgan",
		},
		{
			name: "10",
			in:   "Воробьёво",
			out:  "Vorob’yёvo",
		},
		{
			name: "11",
			in:   "Кебанъёль",
			out:  "Keban”yёl’",
		},
		{
			name: "12",
			in:   "Озёрный",
			out:  "Ozёrnyy",
		},
		{
			name: "13",
			in:   "Тыайа",
			out:  "Tyaya",
		},
		{
			name: "14",
			in:   "Сайылык",
			out:  "Sayylyk",
		},
		{
			name: "15",
			in:   "Ойусардах",
			out:  "Oyusardakh",
		},
		{
			name: "16",
			in:   "Йошкар-Ола",
			out:  "Yoshkar-Ola",
		},
		{
			name: "17",
			in:   "Бийск",
			out:  "Biysk",
		},
		{
			name: "18",
			in:   "Тыэкан",
			out:  "Tyekan",
		},
		{
			name: "19",
			in:   "Суык-Су",
			out:  "Suyk-Su",
		},
		{
			name: "20",
			in:   "Тында",
			out:  "Tynda",
		},
		{
			name: "21",
			in:   "Улан-Удэ",
			out:  "Ulan-Ude",
		},
		{
			name: "22",
			in:   "Электрогорск",
			out:  "Elektrogorsk",
		},
		{
			name: "23",
			in:   "Руэм",
			out:  "Ruem",
		},
		{
			name: "24",
			in:   "Вяртсиля",
			out:  "Vyartsilya",
		},
		{
			name: "25",
			in:   "Ташчишма",
			out:  "Tashchishma",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bgn_pcgn_alt.isBuilt = false
			got := Bgn_pcgn_alt.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Bgn_pcgn_alt)
				t.Errorf("Bgn_pcgn_alt got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Bs_2979 schema
func Test_Bs_2979(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, sʺeshʹ eshchё étikh myagkikh frantsuzskikh bulok iz Ĭoshkar-Olȳ, da vȳpeĭ altaĭskogo chayu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bs_2979.isBuilt = false
			got := Bs_2979.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Bs_2979)
				t.Errorf("Bs_2979 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Bs_2979_alt schema
func Test_Bs_2979_alt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, s\"esh' eshche etikh myagkikh frantsuzskikh bulok iz Ioshkar-Oly, da vypei altaiskogo chayu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bs_2979_alt.isBuilt = false
			got := Bs_2979_alt.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Bs_2979_alt)
				t.Errorf("Bs_2979_alt got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Gost_16876 schema
func Test_Gost_16876(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Ûliâ, sʺešʹ eŝё ètih mâgkih francuzskih bulok iz Joškar-Oly, da vypej altajskogo čaû",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gost_16876.isBuilt = false
			got := Gost_16876.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Gost_16876)
				t.Errorf("Gost_16876 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Gost_16876_alt schema
func Test_Gost_16876_alt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Julija, s\"esh' eshhjo ehtikh mjagkikh francuzskikh bulok iz Jjoshkar-Oly, da vypejj altajjskogo chaju",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gost_16876_alt.isBuilt = false
			got := Gost_16876_alt.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Gost_16876_alt)
				t.Errorf("Gost_16876_alt got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Gost_52290 schema
func Test_Gost_52290(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Россия, город Йошкар-Ола, улица Яна Крастыня",
			out:  "Rossiya, gorod Yoshkar-Ola, ulitsa Yana Krastynya",
		},
		{
			name: "1",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, syesh' eshche etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "2",
			in:   "Ё Крё Мякоё",
			out:  "Yo Krye Myakoyo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gost_52290.isBuilt = false
			got := Gost_52290.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Gost_52290)
				t.Errorf("Gost_52290 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Gost_52535 schema
func Test_Gost_52535(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия Щеглова",
			out:  "Iuliia Shcheglova",
		},
		{
			name: "1",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Iuliia, sesh eshche etikh miagkikh frantcuzskikh bulok iz Ioshkar-Oly, da vypei altaiskogo chaiu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gost_52535.isBuilt = false
			got := Gost_52535.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Gost_52535)
				t.Errorf("Gost_52535 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Gost_7034 schema
func Test_Gost_7034(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, s''esh' eshhyo etix myagkix francuzskix bulok iz Joshkar-Oly, da vypej altajskogo chayu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gost_7034.isBuilt = false
			got := Gost_7034.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Gost_7034)
				t.Errorf("Gost_7034 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Gost_779 schema
func Test_Gost_779(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Ûliâ, sʺešʹ eŝё ètih mâgkih francuzskih bulok iz Joškar-Oly, da vypej altajskogo čaû",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gost_779.isBuilt = false
			got := Gost_779.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Gost_779)
				t.Errorf("Gost_779 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Gost_779_alt schema
func Test_Gost_779_alt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, s``esh` eshhyo e`tix myagkix franczuzskix bulok iz Joshkar-Oly`, da vy`pej altajskogo chayu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gost_779_alt.isBuilt = false
			got := Gost_779_alt.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Gost_779_alt)
				t.Errorf("Gost_779_alt got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Icao_doc_9303 schema
func Test_Icao_doc_9303(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Iuliia, sieesh eshche etikh miagkikh frantsuzskikh bulok iz Ioshkar-Oly, da vypei altaiskogo chaiu",
		},
		{
			name: "1",
			in:   "Юлия Щеглова",
			out:  "Iuliia Shcheglova",
		},
		{
			name: "2",
			in:   "Гайа Васильева",
			out:  "Gaia Vasileva",
		},
		{
			name: "3",
			in:   "Андрей Видный",
			out:  "Andrei Vidnyi",
		},
		{
			name: "4",
			in:   "Артём Краевой",
			out:  "Artem Kraevoi",
		},
		{
			name: "5",
			in:   "Мадыр Чёткий",
			out:  "Madyr Chetkii",
		},
		{
			name: "6",
			in:   "Оксана Клеёнкина",
			out:  "Oksana Kleenkina",
		},
		{
			name: "7",
			in:   "Игорь Ильин",
			out:  "Igor Ilin",
		},
		{
			name: "8",
			in:   "Ян Разъездной",
			out:  "Ian Razieezdnoi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Icao_doc_9303.isBuilt = false
			got := Icao_doc_9303.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Icao_doc_9303)
				t.Errorf("Icao_doc_9303 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Iso_9_1954 schema
func Test_Iso_9_1954(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Julija, s\"ešʹ eščë ėtih mjagkih francuzskih bulok iz Joškar-Oly, da vypej altajskogo čaju",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Iso_9_1954.isBuilt = false
			got := Iso_9_1954.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Iso_9_1954)
				t.Errorf("Iso_9_1954 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Iso_9_1968 schema
func Test_Iso_9_1968(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Julija, sʺešʹ eščë ėtih mjagkih francuzskih bulok iz Joškar-Oly, da vypej altajskogo čaju",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Iso_9_1968.isBuilt = false
			got := Iso_9_1968.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Iso_9_1968)
				t.Errorf("Iso_9_1968 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Iso_9_1968_alt schema
func Test_Iso_9_1968_alt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yulyya, sʺeshʹ eshchë ėtykh myagkykh frantsuzskykh bulok yz Ĭoshkar-Oly, da vypeĭ altaĭskogo chayu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Iso_9_1968_alt.isBuilt = false
			got := Iso_9_1968_alt.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Iso_9_1968_alt)
				t.Errorf("Iso_9_1968_alt got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Mosmetro schema
func Test_Mosmetro(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, syesh esche etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "1",
			in:   "Битцевский парк",
			out:  "Bitsevsky park",
		},
		{
			name: "2",
			in:   "Верхние Лихоборы",
			out:  "Verkhnie Likhobory",
		},
		{
			name: "3",
			in:   "Воробьёвы горы",
			out:  "Vorobyovy gory",
		},
		{
			name: "4",
			in:   "Выхино",
			out:  "Vykhino",
		},
		{
			name: "5",
			in:   "Зябликово",
			out:  "Zyablikovo",
		},
		{
			name: "6",
			in:   "Измайловская",
			out:  "Izmaylovskaya",
		},
		{
			name: "7",
			in:   "Кожуховская",
			out:  "Kozhukhovskaya",
		},
		{
			name: "8",
			in:   "Крылатское",
			out:  "Krylatskoe",
		},
		{
			name: "9",
			in:   "Марьина Роща",
			out:  "Maryina Roscha",
		},
		{
			name: "10",
			in:   "Марьино",
			out:  "Maryino",
		},
		{
			name: "11",
			in:   "Молодёжная",
			out:  "Molodezhnaya",
		},
		{
			name: "12",
			in:   "Октябрьская",
			out:  "Oktyabrskaya",
		},
		{
			name: "13",
			in:   "Ольховая",
			out:  "Olkhovaya",
		},
		{
			name: "14",
			in:   "Парк Победы",
			out:  "Park Pobedy",
		},
		{
			name: "15",
			in:   "Площадь Ильича",
			out:  "Ploschad Ilyicha",
		},
		{
			name: "16",
			in:   "Площадь Революции",
			out:  "Ploschad Revolyutsii",
		},
		{
			name: "17",
			in:   "Пятницкое шоссе",
			out:  "Pyatnitskoe shosse",
		},
		{
			name: "18",
			in:   "Румянцево",
			out:  "Rumyantsevo",
		},
		{
			name: "19",
			in:   "Саларьево",
			out:  "Salaryevo",
		},
		{
			name: "20",
			in:   "Семёновская",
			out:  "Semenovskaya",
		},
		{
			name: "21",
			in:   "Сходненская",
			out:  "Skhodnenskaya",
		},
		{
			name: "22",
			in:   "Текстильщики",
			out:  "Tekstilschiki",
		},
		{
			name: "23",
			in:   "Тёплый стан",
			out:  "Teply stan",
		},
		{
			name: "24",
			in:   "Третьяковская",
			out:  "Tretyakovskaya",
		},
		{
			name: "25",
			in:   "Тропарёво",
			out:  "Troparevo",
		},
		{
			name: "26",
			in:   "Фонвизинская",
			out:  "Fonvizinskaya",
		},
		{
			name: "27",
			in:   "Чистые пруды",
			out:  "Chistye prudy",
		},
		{
			name: "28",
			in:   "Шоссе Энтузиастов",
			out:  "Shosse Entuziastov",
		},
		{
			name: "29",
			in:   "Щёлковская",
			out:  "Schelkovskaya",
		},
		{
			name: "30",
			in:   "Электрозаводская",
			out:  "Elektrozavodskaya",
		},
		{
			name: "31",
			in:   "Юго-Западная",
			out:  "Yugo-Zapadnaya",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mosmetro.isBuilt = false
			got := Mosmetro.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Mosmetro)
				t.Errorf("Mosmetro got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Mvd_310 schema
func Test_Mvd_310(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, syesh' eshche etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "1",
			in:   "Юлия Щеглова",
			out:  "Yuliya Shcheglova",
		},
		{
			name: "2",
			in:   "Гайа Васильева",
			out:  "Gaya Vasilyeva",
		},
		{
			name: "3",
			in:   "Андрей Видный",
			out:  "Andrey Vidnyy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mvd_310.isBuilt = false
			got := Mvd_310.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Mvd_310)
				t.Errorf("Mvd_310 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Mvd_310_fr schema
func Test_Mvd_310_fr(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Iouliia, sech echtche etikh miagkikh frantsouzskikh boulok iz Iochkar-Oly, da vypei altaiskogo tchaiou",
		},
		{
			name: "1",
			in:   "Юлия Щеглова",
			out:  "Iouliia Chtcheglova",
		},
		{
			name: "2",
			in:   "Гайа Васильева",
			out:  "Gaia Vasilieva",
		},
		{
			name: "3",
			in:   "Андрей Видный",
			out:  "Andrei Vidnyi",
		},
		{
			name: "4",
			in:   "Оксана Снегирёва",
			out:  "Oxana Sneguireva",
		},
		{
			name: "5",
			in:   "Юрий Васин",
			out:  "Iourii Vasine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mvd_310_fr.isBuilt = false
			got := Mvd_310_fr.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Mvd_310_fr)
				t.Errorf("Mvd_310_fr got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Mvd_782 schema
func Test_Mvd_782(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, syesh' eshche etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "1",
			in:   "Юлия Щеглова",
			out:  "Yuliya Shcheglova",
		},
		{
			name: "2",
			in:   "Гайа Васильева",
			out:  "Gaya Vasilyeva",
		},
		{
			name: "3",
			in:   "Андрей Видный",
			out:  "Andrey Vidnyy",
		},
		{
			name: "4",
			in:   "Артём Краевой",
			out:  "Artyem Krayevoy",
		},
		{
			name: "5",
			in:   "Мадыр Чёткий",
			out:  "Madyr Chetkiy",
		},
		{
			name: "6",
			in:   "Оксана Клеёнкина",
			out:  "Oksana Kleyonkina",
		},
		{
			name: "7",
			in:   "Игорь Ильин",
			out:  "Igor' Ilyin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mvd_782.isBuilt = false
			got := Mvd_782.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Mvd_782)
				t.Errorf("Mvd_782 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Scientific schema
func Test_Scientific(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Julija, sʺešʹ eščё ètix mjagkix francuzskix bulok iz Joškar-Oly, da vypej altajskogo čaju",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Scientific.isBuilt = false
			got := Scientific.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Scientific)
				t.Errorf("Scientific got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Telegram schema
func Test_Telegram(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Iuliia, sesh esce etih miagkih francuzskih bulok iz Ioshkar-Oly, da vypei altaiskogo chaiu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Telegram.isBuilt = false
			got := Telegram.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Telegram)
				t.Errorf("Telegram got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Ungegn_1987 schema
func Test_Ungegn_1987(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Julija, sʺešʹ eščё ètih mjagkih francuzskih bulok iz Joškar-Oly, da vypej altajskogo čaju",
		},
		{
			name: "1",
			in:   "Россия, город Йошкар-Ола, улица Яна Крастыня",
			out:  "Rossija, gorod Joškar-Ola, ulica Jana Krastynja",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ungegn_1987.isBuilt = false
			got := Ungegn_1987.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Ungegn_1987)
				t.Errorf("Ungegn_1987 got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Uz schema
func Test_Uz(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, Ёшкар-Оладан олинган бу юмшоқ франтсуз рулоларини кўпроқ истеъмол қилинг ва Олтой чойини ичинг",
			out:  "Yuliya, Yoshkar-Oladan olingan bu yumshoq frantsuz rulolarini koʻproq isteʼmol qiling va Oltoy choyini iching",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uz.isBuilt = false
			got := Uz.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Uz)
				t.Errorf("Uz got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Wikipedia schema
func Test_Wikipedia(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, syesh yeshchyo etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "1",
			in:   "Россия, город Йошкар-Ола, улица Яна Крастыня",
			out:  "Rossiya, gorod Yoshkar-Ola, ulitsa Yana Krastynya",
		},
		{
			name: "2",
			in:   "Ельцин",
			out:  "Yeltsin",
		},
		{
			name: "3",
			in:   "Раздольное",
			out:  "Razdolnoye",
		},
		{
			name: "4",
			in:   "Юрьев",
			out:  "Yuryev",
		},
		{
			name: "5",
			in:   "Белкин",
			out:  "Belkin",
		},
		{
			name: "6",
			in:   "Бийск",
			out:  "Biysk",
		},
		{
			name: "7",
			in:   "Подъярский",
			out:  "Podyarsky",
		},
		{
			name: "8",
			in:   "Мусийкъонгийкоте",
			out:  "Musiykyongiykote",
		},
		{
			name: "9",
			in:   "Давыдов",
			out:  "Davydov",
		},
		{
			name: "10",
			in:   "Усолье",
			out:  "Usolye",
		},
		{
			name: "11",
			in:   "Выхухоль",
			out:  "Vykhukhol",
		},
		{
			name: "12",
			in:   "Дальнегорск",
			out:  "Dalnegorsk",
		},
		{
			name: "13",
			in:   "Ильинский",
			out:  "Ilyinsky",
		},
		{
			name: "14",
			in:   "Красный",
			out:  "Krasny",
		},
		{
			name: "15",
			in:   "Великий",
			out:  "Veliky",
		},
		{
			name: "16",
			in:   "Набережные Челны",
			out:  "Naberezhnye Chelny",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Wikipedia.isBuilt = false
			got := Wikipedia.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Wikipedia)
				t.Errorf("Wikipedia got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Yandex_maps schema
func Test_Yandex_maps(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, syesh yeschyo etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name: "1",
			in:   "Россия, город Йошкар-Ола, улица Яна Крастыня",
			out:  "Rossiya, gorod Yoshkar-Ola, ulitsa Yana Krastynya",
		},
		{
			name: "2",
			in:   "Санкт-Петербург, Подъездной пер",
			out:  "Sankt-Peterburg, Podyezdnoy per",
		},
		{
			name: "3",
			in:   "Москва, ул Подъёмная",
			out:  "Moskva, ul Podyomnaya",
		},
		{
			name: "4",
			in:   "Астрахань, ул Подъяпольского",
			out:  "Astrakhan, ul Podyapolskogo",
		},
		{
			name: "5",
			in:   "Щегловитовка",
			out:  "Scheglovitovka",
		},
		{
			name: "6",
			in:   "Новый Уренгой",
			out:  "Noviy Urengoy",
		},
		{
			name: "7",
			in:   "Елабуга",
			out:  "Yelabuga",
		},
		{
			name: "8",
			in:   "Бабаево",
			out:  "Babayevo",
		},
		{
			name: "9",
			in:   "Белово",
			out:  "Belovo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Yandex_maps.isBuilt = false
			got := Yandex_maps.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Yandex_maps)
				t.Errorf("Yandex_maps got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}

// Yandex_money schema
func Test_Yandex_money(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "0",
			in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
			out:  "Yuliya, sesh esche etikh myagkikh frantsuzskikh bulok iz Ioshkar-Oly, da vypei altaiskogo chayu",
		},
		{
			name: "1",
			in:   "Юлия Щеглова",
			out:  "Yuliya Scheglova",
		},
		{
			name: "2",
			in:   "Иван Брызгальский",
			out:  "Ivan Bryzgalskii",
		},
		{
			name: "3",
			in:   "Ксения Стрый",
			out:  "Kseniya Stryi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Yandex_money.isBuilt = false
			got := Yandex_money.Translate(tt.in)
			if got != tt.out {
				fmt.Println(Yandex_money)
				t.Errorf("Yandex_money got:\n%v\nbut want:\n%v\n", got, tt.out)
			}
		})
	}
}
