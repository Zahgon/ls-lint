package rule

import (
	"sync"
)

type KebabCase struct {
	name      string
	exclusive bool
	*sync.RWMutex
}

func (rule *KebabCase) Init() Rule { _ = "STUB: not implemented"; return *new(Rule) }

func (rule *KebabCase) GetName() string { _ = "STUB: not implemented"; return "" }

func (rule *KebabCase) SetParameters(params []string) error { _ = "STUB: not implemented"; return nil }

func (rule *KebabCase) GetParameters() []string { _ = "STUB: not implemented"; return nil }

func (rule *KebabCase) GetExclusive() bool { _ = "STUB: not implemented"; return false }

// Validate checks if string is kebab case
// false if rune is no lowercase letter, digit or -
func (rule *KebabCase) Validate(value string, _ string, _ bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// 45 => -

func (rule *KebabCase) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (rule *KebabCase) Copy() Rule { _ = "STUB: not implemented"; return *new(Rule) }
