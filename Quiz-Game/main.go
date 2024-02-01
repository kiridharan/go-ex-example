package main

// importing packages
import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creating a struc for Q,A
type problem struct {
	q string
	a string
}

// main func
func main() {

	// declaring the flah -csv and talking default file as problems.csv
	// User can give there own file in format (Q,A)
	csvFileName := flag.String("csv", "problems.csv", "The CSV file Format is Question,Answers")

	// Setting Flag for Randomising the
	isRandom := flag.Bool("r", true, "Randomise the Problems")

	timeLimit := flag.Int("tl", 30, "Time Limit for attending the quiz in seconds")

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

	// This is randomising the Problems

	if *isRandom {
		problems = ShuffleProblems(problems)
	}

	// now printing it to user

	Quiz(problems, time.Duration(*timeLimit))

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

// func for checking the ans is correct

func Quiz(problems []problem, timeLimit time.Duration) {
	var correct int

	timer := time.NewTimer(timeLimit * time.Second)
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		// allocates/definig ansChannel to listen the user input
		ansCh := make(chan string)
		// Anonymous Function to check the Go routine , Send
		// answer to channel ansCH
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			ansCh <- answer
		}()

		select {
		// listens to te channel timer and if the timer trigger time exp it'll stop it
		case <-timer.C:
			fmt.Println()
			break problemLoop
		// it'll listens to goroutine  func() 	 at the top if there is change it'll be exc and check
		// ans and update the correct counter
		case answer := <-ansCh:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("You Scored %d out of %d.\n", correct, len(problems))

}

func ShuffleProblems(problems []problem) []problem {
	shuffledProblems := make([]problem, len(problems))
	copy(shuffledProblems, problems)
	for i := len(shuffledProblems) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffledProblems[i], shuffledProblems[j] = shuffledProblems[j], shuffledProblems[i]
	}
	return shuffledProblems
}

// Making the Exit function
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
