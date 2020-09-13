package cmd

import (
	"QuizGame/internal"
	"flag"
	"os"
)

func Start() error {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		return err
	}
	err = internal.CSVRead(file, *timeLimit)
	if err != nil {
		return err
	}
	return nil
}
