# How Go represents data in memory, and how it forces you to be explicit about types 


## Why Go is statically typed AND why it still infers types

Go is statically typed, every variable has a fixed type, checked at `compile time`, not **runtime**. This is a deliberate performance + safety choice:

- **Performance**: the compiler knows exactly how many bytes each variable needs and generates direct machine code for operations, no runtime type-checking overhead like Python's dynamic typing.
- **Safety**: an entire category of bugs ("5" + 5 in JavaScript silently becoming "55") simply cannot compile in Go.

But Go also has **type inference** via `:=`, so you get the safety of static typing without typing var x int = 5 everywhere.This feel the statically typing almost as lightweight to write as a dynamic language.

## Zero values — Go's answer to "uninitialized memory"

This is one of Go's most underrated design decisions. In **C**, an uninitialized variable contains `garbage`, whatever bits happened to be sitting in that memory location. This is a classic source of undefined-behavior bugs.

Go instead guarantees: every variable, if not explicitly initialized, is **automatically** set to its type's **"zero value"** No garbage, ever.

| Type | Zero Value |
| int, float64 | 0, 0.0 |
| bool | false |
| string | "" (empty string, not nil!) |
| pointer, slice, map, channel, function, interface | nil |

This matters at the `memory level`: when Go allocates space for a variable (stack or heap), it zeroes that memory before your code ever touches it. That zeroing is a real, **measurable** operation the runtime performs.

## Why Go has no implicit type conversion

In C: `int x = 5; double y = x;`, silently converts. In Go, this does not compile:
```go 
var x int = 5
var y float64 = x // ❌ compile error: cannot use x (type int) as type float64
```
You must write `float64(x)` explicitly. 

**Why?** Go's designers considered implicit numeric conversion a hidden source of bugs and **precision loss**, e.g., silently truncating an `int64` into an `int32` and losing data without any visible signal in the code. Go forces every conversion to be visible in the source, a reader can grep for every place **truncation/widening** happens.

## Constants — compile-time, not runtime, values

`const` in Go isn't just "a variable you can't change." Go constants are untyped by default and resolved at compile time, which enables a neat trick:

```go
const Pi = 3.14159
var radius float32 = 2
var area = Pi * radius // works! Pi adapts to float32 here
```

If Pi were declared var Pi float64 = 3.14159 instead, this would be a compile error (mixing float64 and float32). Untyped constants act like **literals**, they take on whatever type context they're used in, up until you explicitly give them a type.


## byte vs rune — why this distinction exists
```go
s := "héllo"
fmt.Println(len(s))          // 6, not 5!
for i, r := range s {
    fmt.Println(i, r, string(r))
}
```
Go strings are **UTF-8** encoded byte sequences, not arrays of characters. `é` takes 2 bytes in **UTF-8**, so `len(s)` counts bytes, giving 6 for a 5-character string.
This is a very common early bug: assuming `len(string)` means "number of characters."

`byte (alias uint8)` = one raw byte of that **UTF-8** encoding
`rune (alias int32)` = one decoded Unicode code point (an actual "character")

`for i, r := range s` specifically decodes **UTF-8** for you and gives you runes, with `i` being the **byte** offset (which can skip numbers, e.g., 0, 1, 3, 4, 5 — skipping 2 because é consumed 2 bytes).