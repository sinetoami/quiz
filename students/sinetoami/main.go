package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Problem struct {
	question, answer string
}

func main() {
	filename := "problems.csv"
	file, err := os.Open(filename)

	if err != nil {
		fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Errorf("could not read file: %v", err)
	}

	problems := []Problem{}
	for _, line := range lines {
		problems = append(problems, Problem{line[0], line[1]})
	}

	score := 0
	for i, problem := range problems {
		userAnswer := ""
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		fmt.Scan(&userAnswer)
		if userAnswer == problem.answer {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d.", score, len(problems))
}
