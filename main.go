package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"log"
	"time"
	"context"
	"flag"
	"math/rand"
)

type question struct {
	question string
	answer string
}

var timeout time.Duration
var questions_correct int

func formatData(data [][]string) []question {
	var questions []question
	for _, line := range data {
		var new_question question
		new_question.question = line[0]
		new_question.answer = line[1]
		questions = append(questions, new_question)
	}
	return questions
}

func setupQuestions(questions []question) {
	channel := make(chan int)
	ctx, cancelFunction := context.WithTimeout(context.Background(), timeout)
	
	go startAskingQuestions(questions, channel)
	
	select {
	case <-ctx.Done():
		fmt.Println("Time is up!")
		cancelFunction()
		return
	case <-channel:
		fmt.Println("That was the last question!")
		cancelFunction()
		return
	}
}

func startAskingQuestions(questions []question, channel chan int) {
	var user_answer string
	
	for i, question := range questions {
		fmt.Printf("Question %d: %s?\n", i + 1, question.question)
		fmt.Scanln(&user_answer)
		if user_answer == question.answer {
			questions_correct++
		}
	}
	channel <- 1
}

func main() {
    
	// Check for timeout flag
	var timeout_flag = flag.Int("timeout", 30, "Total amount of time for quiz")
	var shuffle_flag = flag.Bool("shuffle", false, "If true, the questions shuffle before being asked")
	flag.Parse()
	if *timeout_flag > 0 {
		timeout = time.Duration(*timeout_flag) * time.Second
	}

	// Open problems.csv
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Close file after main function is done
	defer file.Close()

	// Read csv file
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Shuffle questions
	if *shuffle_flag == true {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
	}

	// Parse user data into usable questions
	questions := formatData(data)

	// Explain information to user
	var start string
	fmt.Printf("Welcome to the Quiz Game! Can you answer all %d questions correctly in %d seconds?\n", len(questions), timeout / time.Second)
	fmt.Println("Type [start] to continue")
	for start != "start" {
		fmt.Scanln(&start)
	}

	// Setup and start asking questions from the user
	setupQuestions(questions)

	// Print results
	fmt.Printf("You answered %d of %d questions correctly\n", questions_correct, len(questions))
}
