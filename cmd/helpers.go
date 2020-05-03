package cmd

import (
	"fmt"
	"os"
)

func handleError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, "Error: ", err)
	os.Exit(1)
}
