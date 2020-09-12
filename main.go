package main

import (
	"QuizGame/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
