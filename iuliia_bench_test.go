// Package iuliia benchmark tests
// go test -bench=. -benchmem
// go tool pprof mem.out
// go tool pprof cpu.out
// (pprof) web
package iuliia

import (
	"log"
	"testing"
)

// spell-checker: disable

const size = 10000

func prepareData(size int) (string, string) {
	test := struct {
		name string
		in   string
		out  string
	}{
		name: "1",
		in:   "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю",
		out:  "Yuliya, syesh yeshchyo etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
	}

	bigIn, bigOut := "", ""
	for i := 0; i <= size; i++ {
		bigIn += test.in + " "
		bigOut += test.out + " "
	}
	return bigIn, bigOut
}

var bigIn, _ = prepareData(size)

func Benchmark_Translate1(b *testing.B) {
	_, err := Wikipedia.Translate(bigIn)
	if err != nil {
		log.Println(err)
	}
}

func Benchmark_Translate2(b *testing.B) {
	_, err := Wikipedia.translateLargeText(bigIn)
	if err != nil {
		log.Println(err)
	}
}
