# `Iuliia`
> Transliterate Cyrillic → Latin in every possible way

[![GoDoc](https://godoc.org/github.com/mehanizm/iuliia-go?status.svg)](https://pkg.go.dev/github.com/mehanizm/airtable)
![Go](https://github.com/mehanizm/iuliia-go/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/mehanizm/iuliia-go/branch/master/graph/badge.svg)](https://codecov.io/gh/mehanizm/iuliia-go)
[![Go Report](https://goreportcard.com/badge/github.com/mehanizm/iuliia-go)](https://goreportcard.com/report/github.com/mehanizm/iuliia-go)

Transliteration means representing Cyrillic data (mainly names and geographic locations) with Latin letters. It is used for international passports, visas, green cards, driving licenses, mail and goods delivery etc.

This is port of the incredible python library [iuliia](https://github.com/nalgeon/iuliia-py).

`Iuliia` makes transliteration as easy as:

```go
import iuliia "github.com/mehanizm/iuliia-go"

func main() {
    translated, err := iuliia.Wikipedia.Translate("Юлия Щеглова")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(translated)
}

>> 'Yuliya Shcheglova'
```

## Why use `Iuliia`

- 19 transliteration schemas (rule sets), including all main international and Russian standards.
- Correctly implements not only the base mapping, but all the special rules for letter combinations and word endings (AFAIK, Iuliia is the only library which does so).
- Simple API and zero third-party dependencies.

## Installation

```sh
go get github.com/mehanizm/iuliia-go
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Make sure to add or update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
