package rule

import (
	"sync"
)

type CamelCase struct {
	name      string
	exclusive bool
	*sync.RWMutex
}

func (rule *CamelCase) Init() Rule { _ = "STUB: not implemented"; return *new(Rule) }

func (rule *CamelCase) GetName() string { _ = "STUB: not implemented"; return "" }

func (rule *CamelCase) SetParameters(params []string) error { _ = "STUB: not implemented"; return nil }

func (rule *CamelCase) GetParameters() []string { _ = "STUB: not implemented"; return nil }

func (rule *CamelCase) GetExclusive() bool { _ = "STUB: not implemented"; return false }

// Validate checks if string is camel case
// false if rune is no letter and no digit
func (rule *CamelCase) Validate(value string, _ string, _ bool) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// must be letter or digit
		nil
}

// first rune cannot be upper

// rune -1 can be digit

// allow cases like ssrVFor.ts

// rune -1 must be lower

func (rule *CamelCase) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (rule *CamelCase) Copy() Rule { _ = "STUB: not implemented"; return *new(Rule) }
