package rule

import (
	"sync"
)

type Exists struct {
	name      string
	exclusive bool
	min       uint16
	max       uint16
	count     uint16
	test      int
	*sync.RWMutex
}

func (rule *Exists) Init() Rule { _ = "STUB: not implemented"; return *new(Rule) }

func (rule *Exists) GetName() string { _ = "STUB: not implemented"; return "" }

// 0 = regex pattern
func (rule *Exists) SetParameters(params []string) error { _ = "STUB: not implemented"; return nil }

// exists

// exists:

// exists:1

// exists:1-4

func (rule *Exists) GetParameters() []string { _ = "STUB: not implemented"; return nil }

func (rule *Exists) GetExclusive() bool { _ = "STUB: not implemented"; return false }

func (rule *Exists) Validate(value string, _ string, fail bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rule *Exists) getMin() uint16 { _ = "STUB: not implemented"; return 0 }

func (rule *Exists) getMax() uint16 { _ = "STUB: not implemented"; return 0 }

func (rule *Exists) getCount() uint16 { _ = "STUB: not implemented"; return 0 }

func (rule *Exists) incrementCount() { _ = "STUB: not implemented"; return }

func (rule *Exists) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (rule *Exists) Copy() Rule { _ = "STUB: not implemented"; return *new(Rule) }
