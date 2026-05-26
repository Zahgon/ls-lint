package rule

import "sync"

type Error struct {
	Path  string
	Dir   bool
	Ext   string
	Rules []Rule
	*sync.RWMutex
}

func (err *Error) GetPath() string { _ = "STUB: not implemented"; return "" }

func (err *Error) IsDir() bool { _ = "STUB: not implemented"; return false }

func (err *Error) GetExt() string { _ = "STUB: not implemented"; return "" }

func (err *Error) GetRules() []Rule { _ = "STUB: not implemented"; return nil }
