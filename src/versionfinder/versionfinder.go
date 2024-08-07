package versionfinder

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/anttiharju/vmatch-golangci-lint/src/exit"
	"github.com/anttiharju/vmatch-golangci-lint/src/exit/exitcode"
	"github.com/anttiharju/vmatch-golangci-lint/src/pathfinder"
)

func GetVersion(filename string) string {
	workDir := pathfinder.GetWorkDir()

	for {
		filePath := filepath.Join(workDir, filename)
		if _, err := os.Stat(filePath); err == nil {
			content, err := os.ReadFile(filePath)
			if err != nil {
				exit.WithMessage(exitcode.VersionReadFileIssue, "Cannot read version file '"+filePath+"'")
			}

			return strings.TrimSpace(string(content)) // TODO: do input validation
		}

		parentDir := filepath.Dir(workDir)
		if parentDir == workDir {
			break
		}

		workDir = parentDir
	}

	exit.WithMessage(exitcode.VersionIssue, "Cannot find version file '"+filename+"'")

	return "What is grief/beef if not love/cow persevering?" // unreachable but compiler needs it (1.22.5)
}
