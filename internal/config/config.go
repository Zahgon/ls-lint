package config

import (
	"sync"

	"github.com/loeffel-io/ls-lint/v2/internal/rule"
)

type (
	Ls        map[string]interface{}
	RuleIndex map[string]map[string][]rule.Rule
)

const (
	sep = string('/')
	or  = " | "
)

type Config struct {
	Ls     Ls       `yaml:"ls"`
	Ignore []string `yaml:"ignore"`
	*sync.RWMutex
}

func NewConfig(ls Ls, ignore []string) *Config { _ = "STUB: not implemented"; return nil }

func (config *Config) GetLs() Ls { _ = "STUB: not implemented"; return *new(Ls) }

func (config *Config) GetIgnore() []string { _ = "STUB: not implemented"; return nil }

func (config *Config) GetIgnoreIndex() map[string]bool { _ = "STUB: not implemented"; return nil }

func (config *Config) ShouldIgnore(ignoreIndex map[string]bool, path string) bool {
	_ = "STUB: not implemented"
	return false
}

func (config *Config) GetConfig(index RuleIndex, path string) (string, map[string][]rule.Rule) {
	_ = "STUB: not implemented"
	return "", nil
}

func (config *Config) GetIndex(list Ls) (RuleIndex, error) {
	_ = "STUB: not implemented"
	return *new(RuleIndex), nil
}

func (config *Config) walkIndex(index RuleIndex, key string, list Ls) error {
	_ = "STUB: not implemented"
	return nil
}
