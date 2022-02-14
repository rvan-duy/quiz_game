# Quiz Game
Run timed quizzes via the command line.
The program will read in a quiz provided via a CSV file (problems.csv) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question will be asked immediately afterwards. At the end of the quiz the program will output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

## Requirements
- A version of Go (https://go.dev/doc/install)

## Clone and build the game

```
$ git clone https://github.com/rvan-duy/quiz_game
$ cd quiz_game
$ go build
```

## To run the game
```
$ ./quiz_game
```
Quiz game also supports some flags, you can display them as follows:
```
$ ./quiz_game -h
```
Then type for example:
```
$ ./quiz_game -shuffle -timeout=10
```

## Things I learned while working on this project:
- Basic Go syntax
- How to build, run and test Go projects
- Goroutines
- Git branches
- Merging git branches and solving merge conflicts
- Setting up git actions
