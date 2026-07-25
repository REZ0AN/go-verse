# Module 02 — Control Flow

Control flow determines the order in which a program executes statements. Go provides simple but powerful constructs for making decisions, repeating work, postponing work, and handling exceptional situations.

---

## if

Use `if` to execute code only when a condition is true.

```go
if age >= 18 {
    fmt.Println("Adult")
}
```

An optional initializer can be declared before the condition.

```go
if n := len(items); n == 0 {
    fmt.Println("Empty")
}
```

Use `if` when there are only a few conditions to check.

---

## switch

`switch` is often cleaner than a long chain of `if-else` statements.

```go
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Unknown")
}
```

A switch can also omit the expression.

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
}
```

Unlike C/C++, Go automatically breaks after each matching case.

---

## for

Go has only one looping construct: `for`.

Classic loop:

```go
for i := 0; i < 5; i++ {
}
```

While-style loop:

```go
for condition {
}
```

Infinite loop:

```go
for {
}
```

---

## break

`break` immediately exits the nearest loop or `switch`.

```go
for {
    if done {
        break
    }
}
```

A labeled `break` can exit an outer loop.

```go
Outer:
for {
    for {
        break Outer
    }
}
```

---

## continue

`continue` skips the current iteration and starts the next one.

```go
for i := 1; i <= 5; i++ {
    if i%2 == 0 {
        continue
    }

    fmt.Println(i)
}
```

A labeled `continue` jumps to the next iteration of an outer loop.

---

## range

`range` simplifies iteration over collections.

```go
for index, value := range slice {
}
```

It works with:

- Arrays
- Slices
- Maps
- Strings
- Channels (covered later)

When ranging over a slice, the value is a copy.

```go
for _, value := range numbers {
    value *= 10 // numbers is unchanged
}
```

To modify the original slice, use the index.

```go
for i := range numbers {
    numbers[i] *= 10
}
```

---

## defer

`defer` postpones a function call until the surrounding function returns.

```go
defer cleanup()
```

Deferred calls execute in **Last-In, First-Out (LIFO)** order.

```go
defer fmt.Println("First")
defer fmt.Println("Second")
defer fmt.Println("Third")
```

Output:

```
Third
Second
First
```

Arguments to deferred functions are evaluated immediately, even though the function executes later.

```go
x := 5
defer fmt.Println(x)

x = 10
```

Output:

```
5
```

---

## defer in Loops

A common beginner mistake is using `defer` inside a loop.

```go
for i := 1; i <= 3; i++ {
    defer fmt.Println(i)
}
```

Output:

```
3
2
1
```

The deferred calls do **not** execute after each iteration. They all execute only when the surrounding function returns.

This behavior is important when deferring resource cleanup such as closing files or network connections inside loops.

---

## panic

`panic` immediately stops normal execution of the current function.

```go
panic("something went wrong")
```

As the panic propagates up the call stack:

- Normal execution stops.
- Deferred functions still execute.
- If the panic is not recovered, the program terminates.

Use `panic` only for unrecoverable programming errors or exceptional situations.

---

## recover

`recover` stops a panic, but only when called inside a deferred function.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println(r)
    }
}()
```

If a panic occurs:

1. Deferred functions execute.
2. `recover()` captures the panic.
3. The panicking function returns normally.
4. Program execution continues.

Without `recover`, the program exits with a panic.

---

## Labels

A label gives a name to a statement.

Labels are commonly used with `break` and `continue` to control nested loops.

```go
Outer:
for {
    for {
        break Outer
    }
}
```

Labels improve readability when exiting multiple nested loops.

---

## goto

`goto` jumps directly to a labeled statement.

```go
goto End

fmt.Println("Skipped")

End:
fmt.Println("Done")
```

Although Go supports `goto`, it is rarely needed and is generally discouraged because it can make code harder to understand.

---

## fallthrough

Unlike C/C++, Go automatically stops after a matching `case`.

Use `fallthrough` to continue to the next case.

```go
switch n {
case 2:
    fallthrough
case 3:
    fmt.Println("Executed")
}
```

`fallthrough` ignores the next case's condition and always executes the next case.

Use it sparingly, as it can make control flow less obvious.

---

# Key Takeaways

- Use `if` for simple decisions.
- Use `switch` for multiple branches.
- Go has only one looping construct: `for`.
- Use `range` to iterate over collections.
- `break` exits a loop or switch.
- `continue` skips the current iteration.
- `defer` postpones work until the surrounding function returns.
- Deferred calls execute in **LIFO** order.
- Deferred function arguments are evaluated immediately.
- Avoid using `defer` inside long-running loops for resource cleanup.
- `panic` stops normal execution and unwinds the call stack.
- `recover` catches a panic inside a deferred function.
- Labels help control nested loops.
- `goto` exists but should rarely be used.
- `fallthrough` explicitly continues to the next `switch` case.