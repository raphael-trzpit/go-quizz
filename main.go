package main

import (
	"fmt"
	"log"
	"os"

	"github.com/raphael-trzpit/go-quizz/quizz"
)

func main() {
	var questions = []quizz.Question{
		quizz.NewTrueFalse("Jaune est une couleur", true),
		quizz.NewTrueFalse("Arbre est une couleur", false),
	}

	for _, q := range questions {
		quizz.WriteQuestion(q, os.Stdout)

		var response string
		i, err := fmt.Fscanln(os.Stdin, &response)
		switch {
		case 0 == i:
			fmt.Print("You must answer the question !")
		case err != nil:
			fmt.Print(i)
			log.Fatal(err)
		default:
			score, err := q.Score(response)
			if err != nil {
				fmt.Println("Error : " + err.Error())
			}

			if score == 1 {
				fmt.Println("Good job !")
			} else {
				fmt.Println("Error !")
				fmt.Println("The solution was : ")
				quizz.WriteQuestionSolution(q, os.Stdout)
			}
		}
	}

}
