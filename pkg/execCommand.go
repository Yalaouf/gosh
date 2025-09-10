package gosh

import (
	"errors"
	"gosh/pkg/internal"
	"gosh/pkg/utils"
	"os"
	"os/exec"
	"strings"
)

func execCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := utils.SanitizeInput(input)

	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "exit":
		os.Exit(0)

	case "cd":
		if len(args) > 1 {
			return os.Chdir(args[1])
		} else {
			return errors.New(internal.PATH_REQUIRED)
		}

	default:
		break
	}

	res, err := searchPath(args[0])
	if err != nil {
		return err
	}

	cmd := exec.Command(res, args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
