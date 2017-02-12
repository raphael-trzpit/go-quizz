package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
		choices[i] = fmt.Sprintf("%d: %s", i+1, answer.Text)
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
	answers := strings.Split(answer, ",")
AnswersLoop:
	for i, answer := range q.Answers {
		i++
		for _, j := range answers {
			i := strconv.Itoa(i)
			if i == j && !answer.True {
				return false
			}
			if i == j && answer.True {
				continue AnswersLoop
			}
		}
		if answer.True {
			return false
		}
	}
	return true
}
