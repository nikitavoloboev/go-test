package main

import (
	"fmt"
	"os"

	"github.com/nikitavoloboev/go-test/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
