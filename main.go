package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/chrfrasco/git-open/git-remote"
)

func main() {
	remotes, err := gitremote.Parse()
	if err != nil {
		log.Fatal(err.Error())
	}

	target := "origin"
	if len(os.Args) > 1 {
		target = os.Args[1]
	}

	for _, remote := range remotes {
		if remote.Name == target {
			cmd := exec.Command("open", remote.HTTP())
			if err := cmd.Run(); err != nil {
				output, _ := cmd.CombinedOutput()
				log.Fatalf("%s %v", output, err)
			}
		}
	}
}
