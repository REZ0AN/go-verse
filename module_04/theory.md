# Module 04 — Collections, Pointers, and Methods

Go provides several built-in data structures for storing data. Understanding how they work—especially their value vs reference semantics—is essential for writing correct Go programs.

---

## Arrays

An array is a fixed-size collection of elements of the same type.

```go
var numbers [5]int
```

Array literals initialize an array with values.

```go
primes := [5]int{2, 3, 5, 7, 11}
```

The length of an array is part of its type.

```go
var a [3]int
var b [4]int
```

These are different types.

Arrays are **value types**.

```go
a := [3]int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

Copying an array creates an entirely new array.

---

## Slices

Slices provide a flexible view into an array.

```go
numbers := []int{1, 2, 3}
```

A slice contains:

- a pointer to an underlying array
- a length
- a capacity

Create slices with `make`.

```go
numbers := make([]int, 5)
```

Specify both length and capacity.

```go
buffer := make([]int, 2, 8)
```

Retrieve length and capacity.

```go
len(numbers)
cap(numbers)
```

Create a sub-slice.

```go
sub := numbers[1:4]
```

---

## Appending

Append elements to a slice.

```go
numbers = append(numbers, 4)
```

Append multiple values.

```go
numbers = append(numbers, 5, 6, 7)
```

Append another slice.

```go
more := []int{8, 9}

numbers = append(numbers, more...)
```

`append()` returns a new slice because it may allocate a new underlying array.

Always assign its result.

```go
numbers = append(numbers, 10)
```

---

## Maps

A map stores key-value pairs.

```go
ages := map[string]int{
    "Alice": 20,
}
```

Create an empty map.

```go
scores := make(map[string]int)
```

Insert or update values.

```go
scores["Math"] = 90
```

Read values.

```go
score := scores["Math"]
```

Check whether a key exists.

```go
score, ok := scores["Math"]
```

Delete a key.

```go
delete(scores, "Math")
```

Maps do **not** preserve insertion order.

The iteration order is intentionally unpredictable.

---

## Structs

A struct groups related fields together.

```go
type Person struct {
    Name string
    Age  int
}
```

Create a struct.

```go
person := Person{
    Name: "Alice",
    Age: 25,
}
```

Access fields.

```go
person.Name
person.Age
```

Update fields.

```go
person.Age = 26
```

Structs are **value types**.

Copying a struct copies all of its fields.

---

## Pointers

A pointer stores the address of another value.

```go
x := 10

p := &x
```

Dereference a pointer.

```go
fmt.Println(*p)
```

Modify through a pointer.

```go
*p = 20
```

Pointers may also point to structs.

```go
person := Person{}

ptr := &person
ptr.Age = 25
```

Go automatically dereferences pointers when accessing struct fields.

Pointers can also be `nil`.

```go
var p *int

if p == nil {
}
```

---

## Methods

A method is a function associated with a type.

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Call methods using dot notation.

```go
rect.Area()
```

Methods improve readability by attaching behavior directly to a type.

---

## Value Receivers

A value receiver receives a copy of the value.

```go
func (w Wallet) Add(amount int) {
    w.Balance += amount
}
```

Changes affect only the copy.

```go
wallet.Add(100)
```

The original wallet is unchanged.

---

## Pointer Receivers

A pointer receiver receives the address of the value.

```go
func (w *Wallet) Add(amount int) {
    w.Balance += amount
}
```

Changes affect the original value.

```go
wallet.Add(100)
```

Go automatically takes the address when calling pointer receiver methods on values.

Use pointer receivers when:

- the method modifies the receiver
- copying the receiver would be expensive

---

## Slice Memory

Although a slice is passed by value, it still points to the same underlying array.

```go
func modify(s []int) {
    s[0] = 999
}
```

The caller sees the modification because both slices share the same underlying array.

However, `append()` may allocate a new array.

```go
func modify(s []int) {
    s = append(s, 100)
}
```

The caller does not automatically see the appended element because only the local slice now points to the new array.

If the caller should observe the new slice, return it.

```go
func modify(s []int) []int {
    return append(s, 100)
}
```

---

## Map Values are Not Addressable

Map values do not have stable memory addresses.

This does **not** compile.

```go
m["count"].Value = 10
```

Instead, update the value using a read-modify-write sequence.

```go
counter := m["count"]
counter.Value++
m["count"] = counter
```

If direct mutation is required, store pointers.

```go
m := map[string]*Counter{}
```

Pointer values have stable addresses.

---

# Key Takeaways

- Arrays have a fixed size and are value types.
- Slices are lightweight descriptors over an underlying array.
- `append()` may allocate a new array, so always use its returned slice.
- Maps store key-value pairs with unpredictable iteration order.
- Structs group related data and are value types.
- Pointers allow multiple variables to refer to the same value.
- Methods attach behavior to types.
- Value receivers receive copies.
- Pointer receivers can modify the original value.
- Multiple slices can share the same underlying array.
- Map values are not addressable; use read-modify-write or store pointers.