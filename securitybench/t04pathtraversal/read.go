package pathtraversal

import (
	"os"
	"path/filepath"
)

func ReadUserFile(baseDir, name string) ([]byte, error) {
	path := filepath.Join(baseDir, name)
	return os.ReadFile(path)
}
