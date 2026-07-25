# Module 01 — Go Fundamentals

This module introduces the core building blocks of the Go language. By the end of this module, you will be comfortable declaring variables, working with built-in data types, performing type conversions, formatting output, manipulating strings, understanding Unicode, using operators, and reading user input.

---

# Variables

Variables store values that can change during program execution.

Go provides three common ways to declare variables.

```go
var language string = "Go" // explicit type
var version = 1            // type inference
isOpenSource := true       // short declaration
```

### Notes

- `var` can be used inside and outside functions.
- `:=` can only be used inside functions.
- The compiler infers the type whenever possible.

---

# Constants

Constants represent values that never change.

```go
const Pi = 3.14159
const Language = "Go"
```

Constants are evaluated at compile time.

### Typed vs Untyped Constants

```go
const Pi = 3.14159      // untyped
const MaxUsers int = 10 // typed
```

Untyped constants automatically adapt to the required type.

### iota

`iota` automatically generates increasing integer constants.

```go
const (
    Sunday = iota
    Monday
    Tuesday
)
```

Output

```
Sunday  = 0
Monday  = 1
Tuesday = 2
```

---

# Zero Values

Every variable in Go is automatically initialized.

| Type | Zero Value |
|------|------------|
| int | 0 |
| float64 | 0 |
| bool | false |
| string | "" |
| pointer | nil |
| slice | nil |
| map | nil |
| channel | nil |
| interface | nil |
| function | nil |

Unlike C/C++, Go never leaves variables uninitialized.

---

# Built-in Data Types

Go's built-in types include

## Numeric

- int
- int8
- int16
- int32
- int64

- uint
- uint8 (byte)
- uint16
- uint32
- uint64

- float32
- float64

- complex64
- complex128

## Other

- bool
- string
- rune (alias of int32)
- byte (alias of uint8)

Use `%T` with `fmt.Printf` or `fmt.Fprintf` to inspect a value's type.

---

# Type Conversions

Go does **not** perform implicit type conversions.

```go
var x int = 10

var y float64 = float64(x)
```

Every conversion must be explicit.

### Integer → String

```go
strconv.Itoa(123)
```

or

```go
fmt.Sprintf("%d", 123)
```

### String → Integer

```go
value, err := strconv.Atoi("123")
```

### Common Mistake

```go
string(123)
```

does **not** produce

```
"123"
```

It converts `123` into the Unicode character whose code point is `123`.

---

# Formatting Output

The `fmt` package provides several formatting functions.

| Function | Purpose |
|----------|---------|
| Fprint | Write without newline |
| Fprintln | Write with newline |
| Fprintf | Formatted output |
| Sprintf | Format into a string |

Common formatting verbs

| Verb | Meaning |
|------|---------|
| `%d` | Integer |
| `%f` | Floating point |
| `%s` | String |
| `%t` | Boolean |
| `%v` | Default format |
| `%T` | Type |
| `%q` | Quoted string |

---

# Strings

Strings are immutable sequences of UTF-8 encoded bytes.

```go
greeting := "Hello"
```

Common operations

```go
len(greeting)
greeting[0]
greeting[1:4]

greeting + " World"
```

Strings cannot be modified directly.

```go
// greeting[0] = 'J' // compile error
```

Instead, create a new string.

---

# Unicode, Bytes and Runes

Go stores strings using **UTF-8**.

```go
text := "Hello, 世界"
```

### Bytes

`len()` returns the number of bytes.

```go
len(text)
```

### Runes

A rune represents one Unicode code point.

```go
utf8.RuneCountInString(text)
```

### Iterating

```go
for i, r := range text {
    fmt.Println(i, r)
}
```

- `i` is the byte offset.
- `r` is the decoded rune.

Convert to a rune slice when indexing by character is required.

```go
runes := []rune(text)
```

---

# Operators

Go supports the following operator categories.

### Arithmetic

```
+
-
*
/
%
```

### Comparison

```
==
!=
<
<=
>
>=
```

### Logical

```
&&
||
!
```

### Assignment

```
=
+=
-=
*=
/=
%=
```

### Bitwise

```
&
|
^
&^
<<
>>
```

---

# Reading Input

Examples in this repository receive input through an `io.Reader`.

Simple input

```go
fmt.Fscan(reader, &name)
```

Reading an entire line

```go
bufio.NewReader(reader)
```

Using `io.Reader` instead of `os.Stdin` makes code reusable and testable.

---

# Key Takeaways

- Go is statically typed with compile-time type checking.
- Variables can be declared using `var` or `:=`.
- Constants are immutable and evaluated at compile time.
- Every variable has a zero value.
- Go never performs implicit type conversions.
- Strings are immutable UTF-8 encoded byte sequences.
- `byte` represents raw bytes; `rune` represents Unicode code points.
- The `fmt` package provides formatted output functions.
- Input should be read through an `io.Reader` for flexibility and testability.