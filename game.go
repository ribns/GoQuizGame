package main

import (
	"fmt"
	"time"
)

func theStartLine() {
	fmt.Print("Ready? press the enter key to start...")
	fmt.Scanln()
}

func initTimer(duration time.Duration) *time.Timer {
	gameDuration := duration * time.Second
	timer := time.NewTimer(gameDuration)
	fmt.Println("Timer started:", gameDuration)

	return timer
}

func endGame(userScore, maxScore int, timeIsUp bool) {
	if timeIsUp {
		fmt.Println("\nTime's up. 😐")
	}
	fmt.Println("\nYou scored", userScore, "out of", maxScore)
}

func checkUserAnswer(p problem, userAnswer string) bool {
	return p.checkAnswer(userAnswer)
}

func takeUserInput(c chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	c <- answer
}

func printWrongAnswer() {
	fmt.Println("Wrong 🥲")
}

func startGame(problems []problem, duration time.Duration) {
	var answerChannel = make(chan string, 1)
	maxScore := len(problems)
	userScore := 0

	theStartLine()

	timer := initTimer(duration)

	for _, p := range problems {
		p.printProblem()
		go takeUserInput(answerChannel)

		select {
		case <-timer.C:
			endGame(userScore, maxScore, true)
			return

		case answer := <-answerChannel:
			if checkUserAnswer(p, answer) {
				userScore++
			} else {
				printWrongAnswer()
			}
		}
	}

	endGame(userScore, maxScore, false)
}
