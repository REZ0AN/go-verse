# Module 03 — Functions

Functions are one of Go's fundamental building blocks. They let you organize code into reusable units, pass data through parameters, return results, and even treat functions themselves as values.

---

## Functions

A function is a reusable block of code that performs a specific task.

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

Call a function by using its name.

```go
greet("Go")
```

Functions can also return values.

```go
func add(a, b int) int {
    return a + b
}
```

---

## Parameters

Parameters allow functions to receive input.

```go
func rectangleArea(width, height int) int {
    return width * height
}
```

Arguments are supplied when calling the function.

```go
area := rectangleArea(10, 5)
```

Parameters of the same type can share the type declaration.

```go
func add(a, b int) int
```

instead of

```go
func add(a int, b int) int
```

---

## Multiple Return Values

Unlike many languages, Go functions can return multiple values.

```go
func divMod(a, b int) (int, int) {
    return a / b, a % b
}
```

Usage:

```go
quotient, remainder := divMod(17, 5)
```

You can ignore unwanted return values using the blank identifier.

```go
quotient, _ := divMod(20, 3)
```

This feature is heavily used throughout Go, especially for returning a value and an error.

```go
value, err := readFile(...)
```

---

## Named Return Values

Return values can be given names.

```go
func divide(a, b int) (quotient int, remainder int) {
    quotient = a / b
    remainder = a % b
    return
}
```

Named return values are simply local variables initialized to their zero values.

A bare `return` returns the current values of those variables.

Although supported, explicit return values are usually preferred unless named returns improve readability.

---

## Variadic Functions

Variadic functions accept zero or more arguments.

```go
func sum(numbers ...int) int {
}
```

Call it with any number of values.

```go
sum(1, 2, 3)
sum(10, 20)
sum()
```

Inside the function, the variadic parameter is just a slice.

```go
func sum(numbers ...int) {
    // numbers has type []int
}
```

You can also pass an existing slice.

```go
nums := []int{1, 2, 3}

sum(nums...)
```

A variadic parameter must always be the last parameter.

---

## Anonymous Functions

An anonymous function has no name.

```go
func(a, b int) int {
    return a + b
}
```

It can be called immediately.

```go
result := func(a, b int) int {
    return a + b
}(10, 20)
```

Or assigned to a variable.

```go
multiply := func(a, b int) int {
    return a * b
}
```

Functions are values in Go and can be stored in variables.

---

## Closures

A closure is a function that captures variables from its surrounding scope.

```go
func makeCounter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

Each call to `makeCounter()` creates a new independent `count` variable.

```go
c1 := makeCounter()
c2 := makeCounter()

fmt.Println(c1()) // 1
fmt.Println(c1()) // 2
fmt.Println(c2()) // 1
```

The captured variables continue to exist even after the outer function returns.

---

## Higher-Order Functions

A higher-order function either:

- accepts a function as a parameter, or
- returns a function.

Example:

```go
func apply(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}
```

Usage:

```go
result := apply(10, 20, add)
```

Anonymous functions can also be passed directly.

```go
apply(10, 20, func(a, b int) int {
    return a * b
})
```

Many Go standard library packages use higher-order functions.

---

## Recursion

A recursive function calls itself.

Every recursive function must have a base case.

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }

    return n * factorial(n-1)
}
```

Without a base case, recursion continues until the program overflows the stack.

Recursion is useful for naturally recursive problems such as trees, graphs, and divide-and-conquer algorithms.

---

## Closures in Practice

Closures are commonly used to preserve state.

Example:

```go
func makeRateLimiter(limit int) func() bool {
    count := 0

    return func() bool {
        if count < limit {
            count++
            return true
        }

        return false
    }
}
```

Each rate limiter has its own independent state.

```go
userA := makeRateLimiter(3)
userB := makeRateLimiter(3)
```

This pattern is frequently used for:

- Counters
- Rate limiters
- Middleware
- Caching
- Memoization

---

# Key Takeaways

- Functions organize reusable logic.
- Parameters allow functions to receive input.
- Go functions can return multiple values.
- Named returns are local variables initialized to zero values.
- Variadic functions accept zero or more arguments.
- Anonymous functions have no name and are values.
- Closures capture variables from their surrounding scope.
- Each closure has its own captured state.
- Higher-order functions accept or return functions.
- Recursive functions call themselves and require a base case.
- Closures enable practical patterns such as counters and rate limiters.