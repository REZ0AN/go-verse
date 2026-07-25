# Theory — Writing Output

In most Go programs, you'll often see output written like this:

```go
fmt.Println("Hello, Go!")
```

In **Go Verse**, the examples use:

```go
fmt.Fprintln(w, "Hello, Go!")
```

At this stage, you don't need to understand what `w` is or how it works. Just think of it as **the destination where the output should be written**.

For now, both approaches produce the same result:

```text
Hello, Go!
```

The only difference is **how** the output is sent.

Later in the course (Module 5 — Interfaces), you'll learn why this approach is preferred in larger Go applications and how it makes programs more flexible and easier to test.

## Key Takeaways

* `fmt.Println()` prints directly to the terminal.
* `fmt.Fprintln()` writes output to a destination provided by the caller.
* For now, use `fmt.Fprintln()` without worrying about the details.
* You'll learn how `w` works when studying interfaces.
