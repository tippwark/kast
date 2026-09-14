package main

import (
	"os"
)

var (
	compileDate = "unknown"
	kastVersion = "vX.X.X"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
