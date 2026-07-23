package module_03

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_03",
	Description: "Functions i.e multiple/named returns, variadics, closures, and higher-order functions",
	Examples: map[string]cli.Example{
		"functions": {
			Name:        "functions",
			Description: "Multiple/named returns, variadics, closures, and HOC",
			Run:         Function_Examples,
		},

		"rateLimiter_closure": {
			Name:        "rateLimiter_closure",
			Description: "A closure-based rate limiter",
			Run:         RateLimiterClosure_Examples,
		},
	},
}
