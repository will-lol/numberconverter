# numberconverter

⚠️ numberconverter is currently pre-release. The API is unstable and some methods may not be correct. 

[![Go Reference](https://pkg.go.dev/badge/github.com/will-lol/numberconverter.svg)](https://pkg.go.dev/github.com/will-lol/numberconverter)

---

numberconverter provides two simple methods for converting between English and Integers in Go.

```go
numberconverter.Etoi("One hundred thousand, three hundred and fifty two") // 100_352
numberconverter.Itoe(100352) // one hundred thousand three hundred fifty-two
```

## Style agnostic

It doesn't matter how you write your English numbers, they should parse in most cases. There is no prescribed style.

```go
numberconverter.Etoi("Three hundred and forty two million") // 342_000_000
numberconverter.Etoi("nineteen thirty six") // 1936
numberconverter.Etoi("one hundred two hundred") // 100200
numberconverter.Etoi("three hundred, fourty two million") // 342_000_000
```

## Batteries included

Methods for finding and replacing English numbers in your strings are provided.

```go 
numberconverter.FindAllEnglishNumber("Fifty five dogs. Three hundred and twenty three geese.") // {"Fifty five", "Three hundred and twenty three"}
numberconverter.EtoiReplaceAll("If we talk about dogs, I have three. Two of them live in a kennel") // "If we talk about dogs, I have 3. 2 of them live in a kennel"
```

## Generic

Want to parse int8? No problem!

```go
numberconverter.EtoiGeneric[int8]("Fifty five") // 55
```

Numbers above the given integer's maximum will produce unexpected results—be careful!

```go
numberconverter.EtoiGeneric[int8]("Fifty hundred and fifty two million") // 0
```

## Zero configuration

No need to pass around an instance of a struct to use the converter. The package provides two simple methods that may be used globally.

## Zero dependency

Nourish your codebase with pure Go goodness.
