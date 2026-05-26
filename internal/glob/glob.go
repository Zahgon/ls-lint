package glob

import (
	"io/fs"

	"github.com/loeffel-io/ls-lint/v2/internal/config"
)

func Index(filesystem fs.FS, index config.RuleIndex, files bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// delete(index, key) // https://github.com/loeffel-io/ls-lint/issues/249

func IgnoreIndex(filesystem fs.FS, index map[string]bool, files bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}
