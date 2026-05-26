package rule

import (
	"sync"
)

type Lowercase struct {
	name      string
	exclusive bool
	*sync.RWMutex
}

func (rule *Lowercase) Init() Rule { _ = "STUB: not implemented"; return *new(Rule) }

func (rule *Lowercase) GetName() string { _ = "STUB: not implemented"; return "" }

func (rule *Lowercase) SetParameters(params []string) error { _ = "STUB: not implemented"; return nil }

func (rule *Lowercase) GetParameters() []string { _ = "STUB: not implemented"; return nil }

func (rule *Lowercase) GetExclusive() bool { _ = "STUB: not implemented"; return false }

// Validate checks if every letter is lower
func (rule *Lowercase) Validate(value string, _ string, _ bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rule *Lowercase) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (rule *Lowercase) Copy() Rule { _ = "STUB: not implemented"; return *new(Rule) }
