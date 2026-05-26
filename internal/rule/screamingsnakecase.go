package rule

import (
	"sync"
)

type ScreamingSnakeCase struct {
	name      string
	exclusive bool
	*sync.RWMutex
}

func (rule *ScreamingSnakeCase) Init() Rule { _ = "STUB: not implemented"; return *new(Rule) }

func (rule *ScreamingSnakeCase) GetName() string { _ = "STUB: not implemented"; return "" }

func (rule *ScreamingSnakeCase) SetParameters(params []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rule *ScreamingSnakeCase) GetParameters() []string { _ = "STUB: not implemented"; return nil }

func (rule *ScreamingSnakeCase) GetExclusive() bool { _ = "STUB: not implemented"; return false }

// Validate checks if string is screaming sneak case
// false if rune is no uppercase letter, digit or _
func (rule *ScreamingSnakeCase) Validate(value string, _ string, _ bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// 95 => _

func (rule *ScreamingSnakeCase) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (rule *ScreamingSnakeCase) Copy() Rule { _ = "STUB: not implemented"; return *new(Rule) }
