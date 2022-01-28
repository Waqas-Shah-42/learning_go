package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "A csv file in format: question,answer")
	timeLimit := flag.Int("time_limit", 15, "This is the time lint for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", *csvFile)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read the provided file \n\n%s\n", err)
		os.Exit(1)
	}

	problems := parseLines(lines)

	runGame(problems, *timeLimit)
}

// Takes a struct containing problems and answers and runs the game
func runGame(problems []problem, timeLimit int) error {
	correct_ans_count := 0
	wrong_ans_count := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, problem := range problems {
		answerCh := make(chan string)
		go func() {
			fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\n\nCorrect answers: %d\nIncorrect answers: %d\nQuestions not attempted: %d\nTotal questions: %d\n", correct_ans_count, wrong_ans_count, len(problems)-correct_ans_count-wrong_ans_count, len(problems))
			return nil
		case answer := <- answerCh:
			if answer == problem.a {
				fmt.Printf("Correct Answer :-)\n\n")
				correct_ans_count++
			} else {
				fmt.Printf("Wrong Answer :-(\n\n")
				wrong_ans_count++
			}
		}
	}
	fmt.Printf("\n\nCorrect answers: %d\nIncorrect answers: %d\nQuestions not attempted: %d\nTotal questions: %d\n", correct_ans_count, wrong_ans_count, len(problems)-correct_ans_count-wrong_ans_count, len(problems))
	return nil
}

// Parses lines and returns array of stuct containins ques and ans
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// Is a struct which contains both the ques and ans
type problem struct {
	q string
	a string
}
