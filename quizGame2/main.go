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
	// get name of the csv file from user
	fileName := flag.String("csv", "abc.csv", "csv file name")
	timeLimit := flag.Int("limit", 30, "time limit in seconds")
	flag.Parse()

	// read csv file
	file, err := os.Open(*fileName)
	if err != nil {
		exit(err.Error())
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(err.Error())
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(time.Duration(*timeLimit) * time.Second))

	correct := 0
problemsloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerch := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerch <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemsloop

		case answer := <-answerch:
			if answer == p.answer {
				correct++
			}

		}

	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func exit(msg string) {
	println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []Problem {

	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type Problem struct {
	question string
	answer   string
}
