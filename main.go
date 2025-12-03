package main

import (
	"os"

	"github.com/pmady/kube-dependency-checker/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
