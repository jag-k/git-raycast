package main

import (
	"git-raycast/git-raycast/cmd"
	"log"
)

var version = "dev"

func main() {
	cmd.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
