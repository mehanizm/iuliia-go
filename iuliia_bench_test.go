package iuliia

import "testing"

var oneString = "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю"

func prepareData(size int) (res string) {
	for i := 0; i < size; i++ {
		res += oneString
	}
	return
}

func Benchmark_splitSentence(b *testing.B) {
	testData := prepareData(100)
	funcs := []struct {
		name string
		fun  func(source string) []string
	}{
		{"regex", splitSentenceRegex},
		{"unicode", splitSentenceUnicode},
		{"fields", splitSentenceFields},
	}
	for _, funcs := range funcs {
		b.Run(funcs.name, func(b *testing.B) {
			funcs.fun(testData)
		})
	}
}
