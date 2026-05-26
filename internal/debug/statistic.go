package debug

import (
	"sync"
	"time"
)

type Statistic struct {
	Start     time.Time
	Files     int64
	FileSkips int64
	Dirs      int64
	DirSkips  int64
	*sync.RWMutex
}

func NewStatistic() *Statistic { _ = "STUB: not implemented"; return nil }

func (statistic *Statistic) AddFile() { _ = "STUB: not implemented"; return }

func (statistic *Statistic) AddFileSkip() { _ = "STUB: not implemented"; return }

func (statistic *Statistic) AddDir() { _ = "STUB: not implemented"; return }

func (statistic *Statistic) AddDirSkip() { _ = "STUB: not implemented"; return }
