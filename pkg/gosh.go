package gosh

import (
	"bufio"
	"fmt"
	"os"
)

func Gosh() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = execCommand(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
