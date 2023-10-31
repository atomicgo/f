<h1 align="center">AtomicGo | f</h1>

<p align="center">
<img src="https://img.shields.io/endpoint?url=https%3A%2F%2Fatomicgo.dev%2Fapi%2Fshields%2Ff&style=flat-square" alt="Downloads">

<a href="https://github.com/atomicgo/f/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/f?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/f" target="_blank">
<img src="https://img.shields.io/github/actions/workflow/status/atomicgo/f/go.yml?style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/f" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/f?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/f">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-14-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>
  
<a href="https://goreportcard.com/report/github.com/atomicgo/f" target="_blank">
<img src="https://goreportcard.com/badge/github.com/atomicgo/f?style=flat-square" alt="Go report">
</a>   

</p>

---

<p align="center">
<strong><a href="https://pkg.go.dev/atomicgo.dev/f#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/f</pre></h3>
<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# f

```go
import "atomicgo.dev/f"
```

F is the closest thing to template literals in Go.

F is a simple, fast and safe way to format strings in Go, with a familiar syntax. It evaluates expressions inside \`$\{\}\` and replaces them with their values.

The expressions support many operators, including ternary operator, and function calls. See the syntax here: https://expr.medv.io/docs/Language-Definition

<details><summary>Example (Demo)</summary>
<p>



```go
package main

import (
	"fmt"

	"atomicgo.dev/f"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Format a string with a struct
	john := Person{Name: "Bob", Age: 22}
	fmt.Println(f.Format("${Name} is ${Age} years old", john))

	// Format a string with a map
	alice := map[string]any{"Name": "Alice", "Age": 27}
	fmt.Println(f.Format("${Name} is ${Age} years old", alice))

	// Evaluate an expression
	fmt.Println(f.Format("John is 22 years old: ${Age == 22}", john))

	// Ternary operator
	fmt.Println(f.Format("John is 22 years old: ${Age == 22 ? 'yes' : 'no'}", john))

}
```

#### Output

```
Bob is 22 years old
Alice is 27 years old
John is 22 years old: true
John is 22 years old: yes
```

</p>
</details>

## Index

- [func Format\(template string, data ...any\) string](<#Format>)
- [func FormatSafe\(template string, data ...any\) \(string, error\)](<#FormatSafe>)
- [type Parsed](<#Parsed>)
  - [func Parse\(template string\) Parsed](<#Parse>)
  - [func \(parsed Parsed\) Eval\(data any\) \(string, error\)](<#Parsed.Eval>)
  - [func \(parsed Parsed\) String\(\) string](<#Parsed.String>)
- [type Part](<#Part>)


<a name="Format"></a>
## func [Format](<https://github.com/atomicgo/f/blob/main/f.go#L4>)

```go
func Format(template string, data ...any) string
```

Format formats the template string.

<a name="FormatSafe"></a>
## func [FormatSafe](<https://github.com/atomicgo/f/blob/main/f.go#L10>)

```go
func FormatSafe(template string, data ...any) (string, error)
```

FormatSafe formats the template string and returns an additional, optional error, if something goes wrong.

<a name="Parsed"></a>
## type [Parsed](<https://github.com/atomicgo/f/blob/main/parser.go#L12-L15>)

Parsed contains a parsed template string, ready to be evaluated.

```go
type Parsed struct {
    Template string
    Parts    []Part
}
```

<a name="Parse"></a>
### func [Parse](<https://github.com/atomicgo/f/blob/main/parser.go#L18>)

```go
func Parse(template string) Parsed
```

Parse parses a template string into a Parsed struct.

<a name="Parsed.Eval"></a>
### func \(Parsed\) [Eval](<https://github.com/atomicgo/f/blob/main/parser.go#L36>)

```go
func (parsed Parsed) Eval(data any) (string, error)
```

Eval evaluated expressions in the parsed template string.

<a name="Parsed.String"></a>
### func \(Parsed\) [String](<https://github.com/atomicgo/f/blob/main/parser.go#L26>)

```go
func (parsed Parsed) String() string
```

String returns the parsed template parts as a single string.

<a name="Part"></a>
## type [Part](<https://github.com/atomicgo/f/blob/main/parser.go#L60-L63>)

Part is a single part of a template string. Can either be a raw string, or an expression.

```go
type Part struct {
    Value  string
    Parsed bool
}
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
