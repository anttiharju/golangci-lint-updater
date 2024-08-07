package debug

import (
	"os"

	"github.com/anttiharju/vmatch-golangci-lint/src/pathfinder"
)

func getFilePath() string {
	return pathfinder.GetBinDir() + string(os.PathSeparator) + "debug.txt"
}

func WriteToFile(s string) {
	d1 := []byte(s)
	_ = os.WriteFile(getFilePath(), d1, 0o600)
}
