package main

import (
	"flag"
	"time"
)

func main() {
	csvFilename := flag.String("csvfile", "problems.csv", "A CSV file format contains 'question,aswer' lines")
	gameDuration := flag.Int("duration", 30, "The game duration")
	flag.Parse()

	problems := loadProblems(*csvFilename)
	startGame(problems, time.Duration(*gameDuration))
}
