package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"log"
)

type question struct {
	question string
	answer string
}

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

func startAskingQuestions(questions []question) (questions_correct int) {
	var user_answer string
	for i, question := range questions {
		fmt.Printf("Question %d: %s?\n", i + 1, question.question)
		fmt.Scanln(&user_answer)
		if user_answer == question.answer {
			questions_correct++
		}
	}
	return
}

func main() {
    
	// Open problems.csv
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Close file after main function is done
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Parse user data into usable questions
	questions := formatData(data)

	// Start asking user questions and keep track of the amount of correct answers
	questions_correct := startAskingQuestions(questions)

	// Print results
	fmt.Println("Done!")
	fmt.Printf("You answered %d of %d questions correctly\n", questions_correct, len(questions))
}