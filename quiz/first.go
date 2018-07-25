package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFile := flag.String("csv", "data.csv", "CSV description...")
	timeLimit := flag.Int("timer", 30, "TIMER description...")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Println("Can't open ", &csvFile)
		os.Exit(1)
	}
	fileBody := csv.NewReader(file)
	lines, err := fileBody.ReadAll()
	if err != nil {
		fmt.Println("Can't read ", &csvFile)
		os.Exit(1)
	}
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	problems := parseLines(lines)
	countCorrect := 0
	i := 0
loop:
	for key, value := range problems {
		fmt.Printf("#%d: %s = \n", i+1, key)
		done := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			done <- answer
		}()

		select {
		case answer := <-done:
			if answer == value {
				countCorrect++
			}
		case <-timer.C:
			fmt.Println()
			break loop
		}
		i++
	}
	fmt.Println("Score ", countCorrect, "/", len(problems))
}

func parseLines(lines [][]string) map[string]string {
	problems := make(map[string]string)
	for _, line := range lines {
		problems[line[0]] = line[1]
	}
	return problems
}
