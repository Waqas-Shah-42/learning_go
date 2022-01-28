package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "A csv file in format: question,answer")
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
	runGame(problems)
}

// Takes a struct containing problems and answers and runs the game
func runGame(problems []problem) (error){
	correct_ans_count := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == problem.a {
			fmt.Printf("Correct\n")
			correct_ans_count++
		}
	}
	fmt.Printf("You got %d out of %d questions Correct.\n", correct_ans_count, len(problems))
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
