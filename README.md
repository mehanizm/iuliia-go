# `Iuliia`
> Transliterate Cyrillic → Latin in every possible way

> This is the port of the incredible python library [iuliia](https://github.com/nalgeon/iuliia-py) made by @nalgeon

[![GoDoc](https://godoc.org/github.com/mehanizm/iuliia-go?status.svg)](https://pkg.go.dev/github.com/mehanizm/iuliia-go)
![Go](https://github.com/mehanizm/iuliia-go/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/mehanizm/iuliia-go/branch/master/graph/badge.svg)](https://codecov.io/gh/mehanizm/iuliia-go)
[![Go Report](https://goreportcard.com/badge/github.com/mehanizm/iuliia-go)](https://goreportcard.com/report/github.com/mehanizm/iuliia-go)

Transliteration means representing Cyrillic data (mainly names and geographic locations) with Latin letters. It is used for international passports, visas, green cards, driving licenses, mail and goods delivery etc.

`Iuliia` makes transliteration as easy as:

```go
import iuliia "github.com/mehanizm/iuliia-go"

func main() {
    translated := iuliia.Wikipedia.Translate("Юлия Щеглова")
    fmt.Println(translated)
}

>> 'Yuliya Shcheglova'
```

## Why use `Iuliia`

- [20 transliteration schemas](https://github.com/nalgeon/iuliia) (rule sets), including all main international and Russian standards.
- Correctly implements not only the base mapping, but all the special rules for letter combinations and word endings (AFAIK, Iuliia is the only library which does so).
- Simple API and zero third-party dependencies.

## Installation

```sh
go get github.com/mehanizm/iuliia-go
```

`schemas` folder is the git submodule from [general repository](https://github.com/nalgeon/iuliia). You can add schemes manually and use generation to generate all code and tests:

```sh
go generate
```

## CLI interface

In the folder `cmd` you can find CLI app to play with the transliteration.

```sh
# run the app with help
./iuliia -h
> Usage of ./iuliia:
>   -schema string
>         choose schema name (default "wikipedia")
>   -show
>         showing list of the schemas

# simple run the program
./iuliia
> type phrase to translate with schema wikipedia:
> 'quit' to exit or 'help' to get help
> > 

# show help
> help
>> * show - to show all schemas
>> * change schema_name - to change schema
>> * quit - to quit the program
>> type phrase to translate with schema wikipedia:
>> 'quit' to exit or 'help' to get help

# translate
> Привет. Как у тебя дела?
>> Privet. Kak u tebya dela?

# show all schemas
> show
>> ala_lc:              ALA-LC transliteration schema.
>> ala_lc_alt:          ALA-LC transliteration schema.
>> bgn_pcgn:            BGN/PCGN transliteration schema
>> bgn_pcgn_alt:        BGN/PCGN transliteration schema
>> bs_2979:             British Standard 2979:1958 transliteration schema
>> bs_2979_alt:         British Standard 2979:1958 transliteration schema
>> gost_16876:          GOST 16876-71 (aka GOST 1983) transliteration schema
>> gost_16876_alt:      GOST 16876-71 (aka GOST 1983) transliteration schema
>> gost_52290:          GOST R 52290-2004 transliteration schema
>> gost_52535:          GOST R 52535.1-2006 transliteration schema
>> gost_7034:           GOST R 7.0.34-2014 transliteration schema
>> gost_779:            GOST 7.79-2000 (aka ISO 9:1995) transliteration schema
>> gost_779_alt:        GOST 7.79-2000 (aka ISO 9:1995) transliteration schema
>> icao_doc_9303:       ICAO DOC 9303 transliteration schema
>> iso_9_1954:          ISO/R 9:1954 transliteration schema
>> iso_9_1968:          ISO/R 9:1968 transliteration schema
>> iso_9_1968_alt:      ISO/R 9:1968 transliteration schema
>> mosmetro:            Moscow Metro map transliteration schema
>> mvd_310:             MVD 310-1997 transliteration schema
>> mvd_310_fr:          MVD 310-1997 transliteration schema
>> mvd_782:             MVD 782-2000 transliteration schema
>> scientific:          Scientific transliteration schema
>> telegram:            Telegram transliteration schema
>> ungegn_1987:         UNGEGN 1987 V/18 transliteration schema
>> wikipedia:           Wikipedia transliteration schema
>> yandex_maps:         Yandex.Maps transliteration schema
>> yandex_money:        Yandex.Money transliteration schema

# change schema
> change telegram
>> schema was changed to telegram
>> 
>> type phrase to translate with schema telegram:
>> 'quit' to exit or 'help' to get help
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Make sure to add or update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
