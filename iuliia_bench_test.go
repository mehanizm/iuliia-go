package iuliia

import (
	"fmt"
	"math"
	"testing"
)

// benchmark split methods
var oneString = "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы, да выпей алтайского чаю"

func prepareData(size int) (res string) {
	for i := 0; i < size; i++ {
		res += oneString
	}
	return
}

func Benchmark_splitSentence(b *testing.B) {

	splitters := []struct {
		name string
		fun  func(source string) []string
	}{
		{"regex", splitSentenceRegex},
		{"unicode", splitSentenceUnicode},
		{"fields", splitSentenceFields},
		{"isCyrillic", splitSentence},
	}
	for _, splitter := range splitters {
		for k := 0.; k <= 12; k++ {
			n := int(math.Pow(2, k))
			testData := prepareData(n)
			b.Run(fmt.Sprintf("%v–n=%v", splitter.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					splitter.fun(testData)
				}
			})
		}
	}
}

// bechmark get pairs methods
var in = []rune{'a', 'b', 'c'}

func BenchmarkGetPairs(b *testing.B) {
	getters := []struct {
		name string
		fun  func(in []rune) (string, string, string)
	}{
		{"getPairs1", getPairs1},
		{"getPairs1", getPairs2},
	}
	for _, getter := range getters {
		b.Run(getter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getter.fun(in)
			}
		})
	}
}

// overall benchmark
var result string

func prepareTestData(size int) string {
	res := ""
	for i := 0; i < size; i++ {
		res += testString
	}
	return res
}
func BenchmarkIuliia(b *testing.B) {
	b.StopTimer()
	n := 256
	testData := prepareTestData(n)
	b.StartTimer()
	var r string
	for n := 0; n < b.N; n++ {
		r = Wikipedia.Translate(testData)
	}
	result = r
}

var testString = `
Юлия → Iuliia. Всё о транслитерации
Open source,
Программирование,
Алгоритмы,
Интерфейсы
Транслитерация


Транслитерация — это запись кириллических слов латиницей (Анна → Anna, Самара → Samara). Её используют в загранпаспортах, водительских удостоверениях, трансграничной доставке, библиотечных каталогах и множестве других международных процессов.


Так вышло, что я недавно окунулся в эту тему, а в Википедии она раскрыта слабо. Поэтому расскажу, что к чему (спойлер — если вы думаете, что с транслитерацией всё плохо, то на самом деле всё ещё хуже).


И конечно, поскольку это Хабр — предложу open-source библиотеки для решения проблемы.


Оглавление:


Кто виноват
Как выбрать схему (быстрый вариант)
Как выбрать схему (для дотошных)
Как транслитерировать

Кто виноват

Транслит — это хрестоматийная ситуация «у нас 14 плохих стандартов, давайте придумаем ещё один». Весь 20 век солидные, уважаемые люди придумывали всё новые и новые стандарты транслитерации.


Как приумножаются стандарты
Как приумножаются стандарты // xkcd


Получалось у них очень, очень плохо. Например, в загранпаспорте пишут Юлия → Iuliia не потому, что МИД хочет сделать вам больно, а потому что это международный стандарт ICAO Doc 9303 [1] — Machine Readable Travel Documents.


Такое ощущение, что все стандарты писались людьми, которые ненавидят русский язык. Если для англо-американского творчества это объяснимо, то что заставило советских учёных превратить Лёгкий в Ljogkijj (ГОСТ 16876-71[2]) — решительно непонятно.


Zato naši kosmičeskie korabli borozdili prostory vselennoj.
В 21 веке человечество оказалось с двумя наиболее распространёнными стандартами: ICO Doc 9303 (Юлия → Iuliia) и ISO 9:1995 (Юлия → Ûliâ), он же отечественный ГОСТ 7.79-2000[3]. Достойный результат для столетних усилий, ничего не скажешь.


Посмотрев на эту «красоту», ребята из Википедии взялись за голову и сделали нормальную схему транслитерации[4], благодаря которой у несчастной Юлии остаётся слабый шанс быть Yuliya. Международные и отечественные институты эту работу проигнорировали, к сожалению.


Yuliya, syesh yeshchyo etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu.
Конечно, нельзя было делать совсем уж хорошо (а то кто тогда станет придумывать новые стандарты). Поэтому у Википедии ещё превращается в yeshchyo. Схема хорошо передаёт фонетику, а вот выглядит иногда не очень — оцените E → YE, Щ → SHCH и Ё → YO в этом примере.


Не остался в стороне и Яндекс. У него две схемы — отдельно для ФИО[5], отдельно для адресов[6]. Здесь наконец-то сделали Щ → SCH. Но Юрий → Yurii, а Усолье → Usole, что понравится не всем. Не забываем оставлять пространство для новых стандартов!


И Студия Лебедева туда же (в рунете ничего без неё не обходится). Когда дизайнили схему московского метро, ребята отвергли стандарт ISO, а прочие, похоже, даже не смотрели. Ну и придумали свой вариант — Мосметро[7].


Чтобы вы представляли масштаб бедствия. Я насчитал 20 схем транслитерации, некоторые из которых предусматривают альтернативные наборы правил (например, с диакритикой и без). Из них 14 считаются действующими. Четырнадцать действующих «стандартов», прямо как в комиксе xkcd.


Кстати, на Хабре тоже есть транслитерация — например, для генерации id заголовков из текста. И угадайте что? Ну да: она не следует ни одному из официальных стандартов.
Habr, chto ty delaesh, ahaha, prekrati
В качестве вишенки на торте в рунете несметное количество сервисов типа «транслитерация онлайн», которые мало того что перевирают существующие схемы, так ещё и придумывают собственные. Исходники у них закрыты, разумеется.


Есть ещё такая штука — «обратная транслитерация», когда вы восстанавливаете справедливость и превращаете Iuliia → Юлия. Тут ситуация даже хуже, чем в прямой транслитерации, потому что при записи латиницей никто никаких стандартов не соблюдает, и встречаются жуткие монстры.

Обратная транслитерация — отдельная большая тема, здесь я её не рассматриваю. Но есть отличная статья DaryaRodionova — «Как написать свой транслитератор»[8].
Теперь несколько практических рекомендаций.


Как выбрать схему (быстрый вариант)

Загранпаспорт или в/у. По умолчанию используйте ICAO Doc 9303 — это требование закона. Впрочем, есть лайфхак: если написать отдельное заявление при подаче документов, сделают паспорт с нормальной транслитерацией. Тогда подойдёт старый стандарт МВД-310[9] или Мосметро.


import iuliia

source = "Юлия Щеглова"
iuliia.translate(source, schema=iuliia.ICAO_DOC_9303)
# Iuliia Shcheglova

Если нужно обратимое преобразование (cyr–lat). Используйте ГОСТ 7.79-2000 (aka ISO 9:1995). Там есть вариант с диакритикой и без.


import iuliia

source = "Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы"

iuliia.translate(source, schema=iuliia.GOST_779)
# Ûliâ, sʺešʹ eŝё ètih mâgkih francuzskih bulok iz Joškar-Oly

iuliia.translate(source, schema=iuliia.GOST_779_ALT)
# Yuliya, s""esh" eshhyo е"tix myagkix franczuzskix bulok iz Joshkar-Oly"

Если визуальная красота превыше всего. Используйте схему Мосметро, она сама лаконичная и приятная на вид.


import iuliia

source = "Юлия Щеглова"
iuliia.translate(source, schema=iuliia.MOSMETRO)
# Yuliya Scheglova

В остальных случаях. Используйте схему Википедии. Она лучше всех по фонетике и лишь немного уступает Мосметро визуально.


import iuliia

source = "Россия, город Йошкар-Ола, улица Яна Крастыня"
iuliia.translate(source, schema=iuliia.WIKIPEDIA)
# Rossiya, gorod Yoshkar-Ola, ulitsa Yana Krastynya

Как выбрать схему (для дотошных)

Чтобы не раздувать статью до неприличия, я сделал страницу со всеми схемами [10]. Там и сценарии использования, и фильтры, и подробные описания, и примеры. Читайте, выбирайте, что больше нравится. Все схемы уже реализованы на JavaScript, Python и Go, подключить библиотеку — минутное дело.


Здесь же перечислю вкратце только актуальные схемы с моими комментариями.


ГОСТ 7.79-2000

Универсальная схема транслитерации, международный стандарт ISO 9:1995. Громоздкая, зато обратимая. Есть вариант с диакритикой.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Ûliâ, sʺešʹ eŝё ètih mâgkih francuzskih bulok iz Joškar-Oly,
da vypej altajskogo čaû
или
Yuliya, s""esh" eshhyo е"tix myagkix franczuzskix bulok iz Joshkar-Oly",
da vy"pej altajskogo chayu

ГОСТ Р 52290-2004

Стандарт для транслитерации имён собственных на дорожных знаках. Неплохая, с одинарными апострофами. Много внимания написанию Е и Ё.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Yuliya, syesh' eshche etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly,
da vypey altayskogo chayu

ГОСТ Р 7.0.34-2014

Правила упрощенной транслитерации русского письма латинским алфавитом. Для библиотек и издательств. В меру приятная. Расслабленная: есть альтернативы для многих букв, можно без апострофов.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Yuliya, s''esh' eshhyo etix myagkix francuzskix bulok iz Joshkar-Oly,
da vypej altajskogo chayu
или
Yuliya, sesh eshhyo etikh myagkikh francuzskikh bulok iz Yoshkar-Oly,
da vypey altayskogo chayu

Телеграммы

Инструкция Минсвязи о порядке обработки международных телеграмм. Некрасивая, зато без апострофов.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Iuliia, sesh esce etih miagkih francuzskih bulok iz Ioshkar-Oly,
da vypei altaiskogo chaiu

ICAO DOC 9303

Стандарт Международной организации гражданской авиации. Используется МВД для ФИО в водительских удостоверениях, а МИД — в загранпаспортах. Подходит для гринкарты и вида на жительство. Используется некоторыми платёжными системами.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Iuliia, sieesh eshche etikh miagkikh frantsuzskikh bulok iz Ioshkar-Oly,
da vypei altaiskogo chaiu

UNGEGN 1987 V/18

Схема транслитерации ООН для географических названий. Основана на ГОСТ 16876-71. Такая же страшная, но до сих пор не отменена.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Julija, sʺešʹ eščё ètih mjagkih francuzskih bulok iz Joškar-Oly,
da vypej altajskogo čaju

BGN/PCGN

Древняя схема транслитерации ООН для географических названий. Почти без диакритики, но зачем-то оставила Ё с точками.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Yuliya, s”yesh’ yeshchё etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly,
da vypey altayskogo chayu

ALA-LC

Схема транслитерации американской Библиотеки Конгресса. Страшноватая. Есть варианты с диакритикой и без.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
I͡ulii͡a, sʺeshʹ eshchё ėtikh mi͡agkikh frant͡suzskikh bulok iz Ĭoshkar-Oly,
da vypeĭ altaĭskogo chai͡u
или
Iuliia, s"esh' eshche etikh miagkikh frantsuzskikh bulok iz Ioshkar-Oly,
da vypei altaiskogo chaiu

BS 2979:1958

Схема транслитерации Британской библиотеки. Используется издательством Oxford University Press. Изящно схлопывает окончания ИЙ и ЫЙ, в остальном так себе.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Yuliya, sʺeshʹ eshchё étikh myagkikh frantsuzskikh bulok iz Ĭoshkar-Olȳ,
da vȳpeĭ altaĭskogo chayu
или
Yuliya, s"esh' eshche etikh myagkikh frantsuzskikh bulok iz Ioshkar-Oly,
da vypei altaiskogo chayu

Научная

Великая праматерь всех схем. Используется в научных работах. Из неё вырос ГОСТ 16876-71. Достойна уважения за вклад в историю, но страшная.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Julija, sʺešʹ eščё ètix mjagkix francuzskix bulok iz Joškar-Oly,
da vypej altajskogo čaju

Википедия

Схема транслитерации, которую использует Википедия. Сделана на основе BGN/PCGN со значительными модификациями. Самая продуманная, звучит лучше всех и выглядит приятнее большинства прочих схем.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Yuliya, syesh yeshchyo etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly,
da vypey altayskogo chayu

Мосметро

Схема транслитерации, которую использует Московский метрополитен. Визуально самая приятная из всех, хотя уступает Википедии по фонетической точности. Придумана в Студии Лебедева, официально нигде не описана.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Yuliya, syesh esche etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly,
da vypey altayskogo chayu

Яндекс.Деньги

Правила Яндекса для банковских карт. Простая и удобная схема.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Yuliya, sesh esche etikh myagkikh frantsuzskikh bulok iz Ioshkar-Oly,
da vypei altaiskogo chayu

Яндекс.Карты

Правила Яндекса для адресов. Слегка улучшенная версия Яндекс.Денег.


Юлия, съешь ещё этих мягких французских булок из Йошкар-Олы,
да выпей алтайского чаю
↓ ↓ ↓
Yuliya, syesh eschyo etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly,
da vypey altayskogo chayu

Как транслитерировать

Не пишите логику транслитерации с нуля — велик шанс ошибиться и получить очередную (N+1) схему транслитерации, «спасибо» за которую вам не скажут.


Не берите библиотеки с гитхаба без проверки. Все, что я смотрел — реализуют стандарт некорректно, если он чуть сложнее таблицы с однозначным соответствием.


Мы с mehanizm сделали аккуратные библиотеки с нормальными тестами для Python, JavaScript и Go. Но лучше дополнительно проверьте на паре примеров, а то вы ведь знаете этих программистов ツ


UPD Больше библиотек!


C#, Андрей Белянин
Java, Антон Лаврентьев aka Homyakin
Java, Massita
Java, rrrad
PHP, Антон Перевощиков aka Fett
Ruby, Андрей Никифоров

Пример использования (Python):


import iuliia

# list all supported schemas
for schema_name in iuliia.Schemas.names():
    print(schema_name)

# transliterate using specified schema
source = "Юлия Щеглова"
iuliia.translate(source, schema=iuliia.ICAO_DOC_9303)
# "Iuliia Shcheglova"

# or pick schema by name
schema = iuliia.Schemas.get("wikipedia")
iuliia.translate(source, schema)
# "Yuliya Shcheglova"

Или Go:


import iuliia "github.com/mehanizm/iuliia-go"

func main() {
    translated, err := iuliia.Wikipedia.Translate("Юлия Щеглова")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(translated)
}

>> 'Yuliya Shcheglova'

Схемы транслитерации описаны декларативно в JSON, лежат в отдельном репозитории. Если я какую-то пропустил — вы знаете, что делать ツ


И поделитесь в комментариях — приходилось вам сталкиваться с транслитерацией по работе или в жизни? Какие впечатления?"
`
