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

func startAskingQuestions(questions []question) {
	var user_answer string
	for i, question := range questions {
		fmt.Printf("Question %d: %s?\n", i + 1, question.question)
		fmt.Scanln(&user_answer)
		if user_answer != question.answer {
			break
		}
	}
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

	startAskingQuestions(questions)

	fmt.Println("Finish!")
}