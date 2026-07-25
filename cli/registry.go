package cli

import (
	"io"
	"maps"
	"slices"
)

type ExampleFunc func(
	r io.Reader,
	w io.Writer,
) error

type Example struct {
	Name        string
	Description string
	Run         ExampleFunc
}

type Module struct {
	Name        string
	Description string
	Examples    map[string]Example
}

var modules = map[string]Module{}

func RegisterModule(m Module) {
	if _, exists := modules[m.Name]; exists {
		panic("cli: duplicate module registration: " + m.Name)
	}

	modules[m.Name] = m
}

func LookupModule(name string) (Module, bool) {
	m, ok := modules[name]
	return m, ok
}

func LookupExample(module Module, name string) (Example, bool) {
	example, ok := module.Examples[name]
	return example, ok
}

func AllModules() []Module {
	keys := slices.Sorted(maps.Keys(modules))

	result := make([]Module, 0, len(keys))
	for _, key := range keys {
		result = append(result, modules[key])
	}

	return result
}

func AllExamples(m Module) []Example {
	keys := slices.Sorted(maps.Keys(m.Examples))
	examples := make([]Example, 0, len(keys))
	for _, key := range keys {
		examples = append(examples, m.Examples[key])
	}
	return examples
}
