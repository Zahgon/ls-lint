package linter

import (
	"io/fs"
	"sync"

	"github.com/loeffel-io/ls-lint/v2/internal/config"
	"github.com/loeffel-io/ls-lint/v2/internal/debug"
	"github.com/loeffel-io/ls-lint/v2/internal/rule"
)

const (
	extSep = "."
	dir    = ".dir"
)

type Linter struct {
	root      string
	config    *config.Config
	statistic *debug.Statistic
	errors    []*rule.Error
	*sync.RWMutex
}

func NewLinter(root string, config *config.Config, statistic *debug.Statistic, errors []*rule.Error) *Linter {
	_ = "STUB: not implemented"
	return nil
}

func (linter *Linter) GetStatistics() *debug.Statistic { _ = "STUB: not implemented"; return nil }

func (linter *Linter) GetErrors() []*rule.Error { _ = "STUB: not implemented"; return nil }

func (linter *Linter) AddError(error *rule.Error) { _ = "STUB: not implemented"; return }

func (linter *Linter) validateDir(index config.RuleIndex, path string, validate bool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (linter *Linter) validateFile(index config.RuleIndex, path string, validate bool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// compatibility with windows

// 2^N combinations

// from left to right; right to left: i&(1<<j)
// Keep original

// Replace with "*"

func (linter *Linter) Run(filesystem fs.FS, paths map[string]struct{}, debug bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// create index

// glob index

// glob ignore index

// validate exists
