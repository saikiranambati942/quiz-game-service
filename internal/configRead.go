package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func CSVRead(f *os.File, timeLimit int) error {
	correct := 0
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
problemLoop:
	for i, problem := range problems {
		fmt.Printf("#Problem%d: %s = ", i+1, problem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("\nQuizGame time expired")
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.a {
				correct++
			}
		}
	}
	fmt.Printf("No. of correct answers: %d out of %d\n", correct, len(lines))
	return nil
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{q: line[0], a: strings.TrimSpace(line[1])}
	}
	return res
}
