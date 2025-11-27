package main

import (
	"agit/cmd"

	"github.com/pykelysia/pyketools"
)

func main() {
	ps := cmd.NewCmd()
	for {
		if ok, err := ps.Run(); ok {
			pyketools.Infof("Everything is up-to-data.")
			break
		} else if err != nil {
			pyketools.Fatalf("Error occurred: %v", err)
		}
	}
}
