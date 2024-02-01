package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {

	// declaring the flah -csv and talking default file as problems.csv
	// User can give there own file in format (Q,A)
	csvFileName := flag.String("csv", "problems.csv", "The CSV file Format is Question,Answers")
	flag.Parse()
	// Opening the File
	file, err := os.Open(*csvFileName)
	// checking for error is Any error Printing it out and exit
	if err != nil {
		exit(fmt.Sprintf("There is a error while Opeing the File %s Error: %s ", *csvFileName, err))
	}
	r := csv.NewReader(file)
	// lines return a 2D slices
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to Read the Lines in the CSV")
	}
	problems := parseLines(lines)
	// now printing it to user
	correct := Quiz(problems)

	fmt.Printf("You Scored %d out of %d.\n", correct, len(problems))
}

// Function for Parsing lines -> line as Return problem
func parseLines(lines [][]string) []problem {
	//creating a slice with len of lines
	re := make([]problem, len(lines))

	for i, line := range lines {
		re[i] = problem{
			q: line[0],
			// trims the sapces in the String
			a: strings.TrimSpace(line[1]),
		}

	}

	return re
}

// creating a struc for Q,A

// func for checking the ans is correct

func Quiz(problems []problem) int {
	var correct int

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		var answer string
		fmt.Scanf("%s", &answer)

		if answer == p.a {
			correct++
		}
	}

	return correct
}

// Making the Exit function
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
