# Control Flow

## Why Go's if has no parentheses, but requires braces

Compare to **C/Java**: `if (x > 5) { ... }`. Go removes the parentheses around the condition (they're simply not allowed, a stylistic choice paired with `gofmt`, since Go decided the "**prefix-lisp-condition**" parens added visual noise without adding meaning). But Go requires braces even for single-line bodies.

```go
if x > 5
    doSomething(); // ❌ illegal
```
This is a deliberate defensive design decision. It closes off an entire, historically real class of bugs, Apple's infamous 2014 goto fail; SSL bug happened partly because C allows single-statement if bodies without braces, so an extra unconditional line silently fell outside the intended if. Go's compiler simply refuses to let that ambiguity exist.

## The if statement's "hidden" initializer

```go
if err := doSomething(); err != nil {
    // handle err
}
```
This is huge and you'll see it everywhere in real Go code. The if statement supports an optional init statement before the condition, separated by ;. The variable declared there (err) is scoped only to the if/else blocks, it doesn't leak into the enclosing function scope.

## switch 

Two **critical** differences from **C/Java** `switch`:

- No **fallthrough** by default. In C, cases fall through to the next case unless you `break`. Go does the opposite, each case **automatically** breaks, and if you actually want **fallthrough** behavior, you must write the `fallthrough` keyword explicitly. Go's designers considered `C`'s **implicit-fallthrough-by-default** to be a major **footgun** (forgetting a break is one of the most common C bugs ever), so they flipped the **default**.
- Switch doesn't require a value at all. `switch { case x > 10: ... }`, this is just a cleaner if/else if chain. This is genuinely idiomatic Go, not a hack.

## for

Go has no `while`, no `do-while` — just `for`, which subsumes all three shapes:

```go
for i := 0; i < 10; i++ { }   // classic C-style for
for x < 10 { }                  // while-style
for { }                          // infinite loop (while(true) equivalent)
for i, v := range collection { } // range-based iteration
```

## range

| Collection | for k, v := range x gives you |
| slice/array | index (int), element value (copy) |
| map | key, value |
| string | byte index, rune (decoded UTF-8 code point) |
| channel | value only (single variable form) |

Critical internal detail: for `slices/arrays`, `v` is a **copy** of each element, not a reference into the original collection. Mutating `v` inside the loop does not mutate the original slice. This is a very common bug source, but flagging it now since it's a range behavior, not a slice behavior.

## What defer actually does
```go
func readFile() {
    f, _ := os.Open("data.txt")
    defer f.Close() // scheduled, NOT executed now
    // ... work with f ...
} // f.Close() runs HERE, right before the function returns
```
**defer** -> schedules a function call to run when the surrounding function returns, not immediately, and not at the end of the current block, but at the very end of the function's execution, regardless of how it returns (normal return, or a panic unwinding through it).

**Why this exists:** in C, resource cleanup (fclose, free, mutex unlock) has to be manually placed before every return path, including early returns, error paths, and paths through nested ifs. Miss one, and you leak a file handle or leave a mutex locked forever. defer lets you write the cleanup once, right next to the acquisition, and the runtime guarantees it fires no matter which exit path the function takes.
Deferred calls run in `LIFO` order — Last In, First Out, like a stack:
```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
// prints: 3, 2, 1
```

Critical, commonly-misunderstood detail: `arguments` to a deferred call are evaluated immediately, at the point of the defer statement, only the call itself is delayed.
```go
func example() {
    x := 1
    defer fmt.Println("x was:", x) // "x" is captured as 1 RIGHT NOW
    x = 2
    fmt.Println("x is now:", x)
}
// prints:
// x is now: 2
// x was: 1
```
This surprises almost everyone the first time. `x` is evaluated and copied into the deferred call's argument list at the moment defer executes, not at the moment the deferred call actually runs.

**panic** —> Go's mechanism for **unrecoverable-by-default** errors. Calling `panic(v)` immediately stops normal execution of the current function, runs any deferred calls in that function (and then the caller's, and the caller's caller's...) as the stack unwinds, and if nothing intercepts it, crashes the whole program with a stack trace.
**recover** —> can only usefully be called inside a deferred function, and it stops a panic from propagating further, letting the program continue. It's Go's closest analog to try/catch, but deliberately harder to reach for casually.

## Why panic/recover isn't "Go's exceptions"

This is an important philosophical point, not just trivia. Go's designers wanted error return values to be the `default`, expected path for anything that can normally fail (file not found, network timeout, invalid input). `panic`/`recover` is reserved for truly exceptional, **programmer-error-level** situations, `nil` pointer dereference, index out of bounds, or deliberately signaling "this invariant should never be violated." Idiomatic Go code rarely calls `panic()` directly outside of package-init-time sanity checks; it's not meant to replace `if err != nil` as a general control-flow tool.

```go
defer functionCall(args)         // schedules functionCall(args) for function-return time

func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    result = a / b // panics if b == 0 (division by zero)
    return
}
```
`defer func() { ... }()` —> deferring an anonymous function, immediately-invoked-inside-the-defer. This is the standard pattern for recover, since `recover()` only has effect when called directly inside a deferred function.
`recover() ` -> returns the value passed to `panic()`, or `nil` if there was no panic. Calling `recover()` outside a deferred function always returns `nil` and does nothing.
**Named return values** `(result int, err error)` —>  The deferred function can modify `err` even after return was hit, because return on a named-return function doesn't **immediately** exit, deferred calls run first and can still **mutate** those named return variables.

## `labels` & `goto` in GO

**Labels** exist almost exclusively to solve one specific problem: break and continue normally only affect the innermost loop. If you have nested loops and need to break/continue an outer loop from inside an inner one, a label lets you target it explicitly.

```go
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue outer // skips to next i, not next j
            }
            fmt.Println(i, j)
        }
    }
```

**goto** is a direct, unconditional jump to a labeled statement. Go kept it (many modern languages dropped it entirely) but with strict rules: you can't **jump into a block**, or **jump over a variable** declaration in a way that would leave it uninitialized when reached. It's rarely used in practice, mostly you'll see it in generated code or very specific error-cleanup patterns in older/lower-level code.

```go
func main() {
    i := 0
loop:
    if i < 3 {
        fmt.Println(i)
        i++
        goto loop
    }
}
```