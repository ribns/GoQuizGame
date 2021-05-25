package main

import (
	"fmt"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func (p *problem) printProblem() {
	fmt.Printf("What is %s: ", p.question)
}

func (p *problem) checkAnswer(answer string) bool {
	return strings.TrimSpace(answer) == strings.TrimSpace(p.answer)
}

func createProblems(csvLines *[][]string) (problems []problem) {
	problems = make([]problem, len(*csvLines))

	for i, v := range *csvLines {
		problems[i] = problem{
			question: v[0],
			answer:   v[1],
		}
	}

	return
}

func loadProblems(filename string) []problem {
	/* Load problems from a CSV file */
	file := openFile(filename)
	csvLines := parseCSV(file)
	problems := createProblems(&csvLines)

	return problems
}
