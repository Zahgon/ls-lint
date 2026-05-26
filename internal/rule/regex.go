package rule

import (
	"sync"
)

const negate = '!'

type Regex struct {
	name         string
	exclusive    bool
	regexPattern string
	negate       bool
	*sync.RWMutex
}

func (rule *Regex) Init() Rule { _ = "STUB: not implemented"; return *new(Rule) }

func (rule *Regex) GetName() string { _ = "STUB: not implemented"; return "" }

// 0 = regex pattern
func (rule *Regex) SetParameters(params []string) error { _ = "STUB: not implemented"; return nil }

func (rule *Regex) GetParameters() []string { _ = "STUB: not implemented"; return nil }

func (rule *Regex) GetExclusive() bool { _ = "STUB: not implemented"; return false }

// Validate checks if full string matches regex
func (rule *Regex) Validate(value string, path string, _ bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rule *Regex) getRegexPattern() string { _ = "STUB: not implemented"; return "" }

func (rule *Regex) GetErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (rule *Regex) Copy() Rule { _ = "STUB: not implemented"; return *new(Rule) }
