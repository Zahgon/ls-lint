package rule

import (
	"sync"
)

type SnakeCase struct {
	name      string
	exclusive bool
	*sync.RWMutex
}

func (rule *SnakeCase) Init() Rule { _ = "STUB: not implemented"; return *new(Rule) }

func (rule *SnakeCase) GetName() string { _ = "STUB: not implemented"; return "" }

func (rule *SnakeCase) SetParameters(params []string) error { _ = "STUB: not implemented"; return nil }

func (rule *SnakeCase) GetParameters() []string { _ = "STUB: not implemented"; return nil }

func (rule *SnakeCase) GetExclusive() bool { _ = "STUB: not implemented"; return false }

// Validate checks if string is sneak case
// false if rune is no lowercase letter, digit or _
func (rule *SnakeCase) Validate(value string, _ string, _ bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// 95 => _

func (rule *SnakeCase) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (rule *SnakeCase) Copy() Rule { _ = "STUB: not implemented"; return *new(Rule) }
