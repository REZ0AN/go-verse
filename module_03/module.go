package module_03

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_03",
	Description: "Functions i.e multiple/named returns, variadics, closures, and higher-order functions",
	Examples: map[string]cli.Example{
		"functions": {
			Name:        "functions",
			Description: "Declare and call functions",
			Run:         Functions,
		},
		"parameters": {
			Name:        "parameters",
			Description: "Function parameters",
			Run:         Parameters,
		},
		"multiple_returns": {
			Name:        "multiple_returns",
			Description: "Return multiple values",
			Run:         MultipleReturns,
		},
		"named_returns": {
			Name:        "named_returns",
			Description: "Named return values",
			Run:         NamedReturns,
		},
		"variadic": {
			Name:        "variadic",
			Description: "Variadic functions",
			Run:         Variadic,
		},
		"anonymous": {
			Name:        "anonymous",
			Description: "Anonymous functions",
			Run:         Anonymous,
		},
		"closures": {
			Name:        "closures",
			Description: "Closures and captured variables",
			Run:         Closures,
		},
		"higher_order": {
			Name:        "higher_order",
			Description: "Higher-order functions",
			Run:         HigherOrder,
		},
		"recursion": {
			Name:        "recursion",
			Description: "Recursive functions",
			Run:         Recursion,
		},
		"rate_limiter": {
			Name:        "rate_limiter",
			Description: "Build a simple rate limiter with closures",
			Run:         RateLimiter,
		},
	},
}
