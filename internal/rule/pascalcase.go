package rule

import (
	"sync"
)

type PascalCase struct {
	name      string
	exclusive bool
	*sync.RWMutex
}

func (rule *PascalCase) Init() Rule { _ = "STUB: not implemented"; return *new(Rule) }

func (rule *PascalCase) GetName() string { _ = "STUB: not implemented"; return "" }

func (rule *PascalCase) SetParameters(params []string) error { _ = "STUB: not implemented"; return nil }

func (rule *PascalCase) GetParameters() []string { _ = "STUB: not implemented"; return nil }

func (rule *PascalCase) GetExclusive() bool { _ = "STUB: not implemented"; return false }

// Validate checks if string is pascal case
// false if rune is no letter and no digit
// false if first rune is not upper
func (rule *PascalCase) Validate(value string, _ string, _ bool) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// must be letter or digit
		nil
}

// first rune must be upper

// rune -1 can be digit

// allow cases like SsrVFor.ts

// rune -1 must be lower

func (rule *PascalCase) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (rule *PascalCase) Copy() Rule { _ = "STUB: not implemented"; return *new(Rule) }
