# Go-Verse [Learning by Doing]

Go-Verse is a collection of executable examples, here to help you learn Go through experimentation. Instead of following long tutorials, each concept is implemented as a small, focused program that can be explored independently.

The project is organized into modules covering different areas of the language, with a simple command-line interface for discovering and running examples.

## Getting Started

Clone the repository and run the CLI.

```bash
go run .
```

This displays all available learning modules.

Example:

```text
Go Verse - A Go Learning Playground

Available Modules

module_01     Go Fundamentals
...
...
...
```

## Viewing Module Documentation

To view all available examples within a module:

```bash
go run . module_01
```

or

```bash
go run . module_01 --help
```

Example output:

```text
module_01

Go Fundamentals

Available Examples

variables
datatypes
operators
formatting
io
strings
```

## Running Examples

Each example can be executed independently.

```bash
go run . module_01 variables
```

Other examples:

```bash
go run . module_01 datatypes
go run . module_01 operators
go run . module_01 formatting
go run . module_01 io
go run . module_01 strings
```

## Adding a New Example

Create a new Go source file inside the appropriate module.

For example:

```text
module_01/
├── arrays.go
```

Implement the example function and register it in the module registry.

Once registered, it becomes available through the CLI.

```bash
go run . module_01 arrays
```

## Adding a New Module

Create a new module directory containing its examples and registry.

Register the module using its `init()` function, then import the package in `main.go` using a blank import.

```go
import (
    _ "github.com/REZ0AN/go-verse/module_06"
)
```

After that, the module will automatically appear in the CLI without any additional changes.

## Philosophy

Go-Verse is built around a simple principle:

> Learn by writing code, running code, and modifying code.

Each example is intentionally small and isolated so concepts can be explored without unnecessary boilerplate. The repository is intended to grow alongside the language learning journey, from Go fundamentals to concurrency, networking, systems programming, and beyond.