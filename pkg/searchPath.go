package gosh

import (
	"gosh/pkg/internal"
	"os"
	"strings"
)

func searchPath(command string) (string, error) {
	env := os.Getenv("PATH")
	envDirs := strings.SplitSeq(env, ":")

	for dir := range envDirs {
		elements, err := os.ReadDir(dir)
		if err != nil {
			if strings.Contains(err.Error(), internal.NOT_FOUND) ||
				strings.Contains(err.Error(), internal.NOT_DIRECTORY) {
				continue
			}

			return "", err
		}

		for _, element := range elements {
			if element.IsDir() {
				continue
			}

			if element.Name() == command {
				return dir + "/" + command, nil
			}
		}
	}

	return "", nil
}
