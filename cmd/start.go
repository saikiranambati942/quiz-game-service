package cmd

import (
	"QuizGame/internal"
	"flag"
	"os"
)

func Start() error {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		return err
	}
	err = internal.CSVRead(file)
	if err != nil {
		return err
	}
	return nil
}
