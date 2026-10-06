package prefixcheck

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadModel(baseDir, name string) ([]byte, error) {
	full := filepath.Join(baseDir, name)
	if !strings.HasPrefix(full, baseDir) {
		return nil, os.ErrPermission
	}
	return os.ReadFile(full)
}
