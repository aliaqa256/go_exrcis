package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// get name of the csv file from user
	fileName := flag.String("csv", "abc.csv", "csv file name")
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
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
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
