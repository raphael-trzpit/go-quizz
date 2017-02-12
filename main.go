package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
			fmt.Print(i)
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
