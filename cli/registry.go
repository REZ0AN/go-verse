package cli

type Example struct {
	Name        string
	Description string
	Run         func()
}

type Module struct {
	Name        string
	Description string
	Examples    map[string]Example
}

var Modules = map[string]Module{}

func RegisterModule(m Module) {
	Modules[m.Name] = m
}
