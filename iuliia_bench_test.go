package iuliia

import (
	"fmt"
	"math"
	"testing"
)

var oneString = "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю"

func prepareData(size int) (res string) {
	for i := 0; i < size; i++ {
		res += oneString
	}
	return
}

func Benchmark_splitSentence(b *testing.B) {

	funcs := []struct {
		name string
		fun  func(source string) []string
	}{
		{"regex", splitSentenceRegex},
		{"unicode", splitSentenceUnicode},
		{"fields", splitSentenceFields},
		{"isCyrillic", splitSentence},
	}
	for _, funcs := range funcs {
		for k := 0.; k <= 12; k++ {
			n := int(math.Pow(2, k))
			testData := prepareData(n)
			b.ResetTimer()
			b.Run(fmt.Sprintf("%v–n=%v", funcs.name, n), func(b *testing.B) {
				funcs.fun(testData)
			})
		}
	}
}
