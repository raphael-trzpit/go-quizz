package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var file = flag.String("file", "quizz.json", "Filepath to the json file defining the quizz.")
	flag.Parse()
	raw, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatalf("Can't read file %s : %s", *file, err.Error())
	}
	var quizzConfig QuizzConfig
	err = json.Unmarshal(raw, &quizzConfig)
	if err != nil {
		log.Fatalf("Invalid json file %s : %s", *file, err.Error())
	}
	quizz := quizzConfig.Quizz

	// Questions.
	for i, question := range quizz.Questions {
		fmt.Printf("Question %d:\n", i)
		fmt.Println(question.Ask())
		for _, choice := range question.GiveChoices() {
			fmt.Println(choice)
		}

		var response string
		i, err := fmt.Fscanln(os.Stdin, &response)
		switch {
		case 0 == i:
			fmt.Print("You must answer the question !")
		case err != nil:
			log.Fatal(err)
		default:
			valid := question.Answer(response)
			if !valid {
				fmt.Println("Wrong !")
				fmt.Println("Good answers were :")
				for _, answer := range question.GiveAnswers() {
					fmt.Println(answer)
				}
			} else {
				fmt.Println("Good job !")
			}
		}
	}

}

// Quizz is a quizz.
type Quizz struct {
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
	Answers     []Answer   `json:"answers"`
}

// QuizzConfig is a config for a quizz.
type QuizzConfig struct {
	Quizz Quizz `json:"quizz"`
}

// Question is a question.
type Question struct {
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

// Answer is an answer
type Answer struct {
	Text string `json:"text"`
	True bool   `json:"true"`
}

// Ask a question.
func (q Question) Ask() string {
	return q.Text
}

// GiveChoices for a question
func (q Question) GiveChoices() []string {
	choices := make([]string, len(q.Answers))
	for i, answer := range q.Answers {
		choices[i] = fmt.Sprintf("%d: %s", i, answer.Text)
	}
	return choices
}

// GiveAnswers for a question
func (q Question) GiveAnswers() []string {
	answers := make([]string, 0)
	for _, answer := range q.Answers {
		if !answer.True {
			continue
		}
		answers = append(answers, answer.Text)
	}
	return answers
}

// Answer a question
func (q Question) Answer(answer string) bool {
	answers := strings.Split(answer, " ")
	sort.Strings(answers)
	for i, answer := range q.Answers {
		if answer.True && sort.SearchStrings(answers, strconv.Itoa(i)) > len(answers) {
			return false
		}
		if !answer.True && sort.SearchStrings(answers, strconv.Itoa(i)) <= len(answers) {
			return false
		}
	}
	return true
}
