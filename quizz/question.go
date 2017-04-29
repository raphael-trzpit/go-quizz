package quizz

import (
	"fmt"
	"io"
	"text/template"
)

// Question is composed by a sentence and a list of answers
type Question interface {
	Score(string) (Score, error)
	Template() QuestionTemplate
}

type QuestionTemplate struct {
	Question string
	Answer   string
}

func WriteQuestion(q Question, w io.Writer) error {
	tmpl := template.Must(template.New("test").Parse(q.Template().Question))
	return tmpl.Execute(w, q)
}

func WriteQuestionSolution(q Question, w io.Writer) error {
	tmpl := template.Must(template.New("test").Parse(q.Template().Answer))
	return tmpl.Execute(w, q)
}

type QuestionCollection struct {
	Questions []Question
}

type QuizzQuestion struct {
	Question *Question
	Answer   string
}

type TrueFalse struct {
	Text             string
	Answer           bool
	QuestionTemplate QuestionTemplate
}

var TrueFalseDefaultTemplate = QuestionTemplate{
	Question: "{{.Text}}\nTrue or False ? [y,n] :",
	Answer:   "{{.Text}} : {{if .Answer}}True{{else}}False{{end}}",
}

func NewTrueFalse(text string, isTrue bool) TrueFalse {
	return TrueFalse{text, isTrue, QuestionTemplate{}}
}

func (q TrueFalse) Template() QuestionTemplate {
	if q.QuestionTemplate.Question == "" || q.QuestionTemplate.Answer == "" {
		return TrueFalseDefaultTemplate
	}
	return q.QuestionTemplate
}

func (q TrueFalse) Score(s string) (Score, error) {
	b, e := stringToBool(s)
	if e != nil {
		return Score(0), e
	}
	if b == q.Answer {
		return Score(1), nil
	}
	return Score(0), nil
}

func stringToBool(s string) (bool, error) {
	if s == "true" || s == "t" || s == "yes" || s == "y" {
		return true, nil
	}
	if s == "false" || s == "f" || s == "no" || s == "n" {
		return false, nil
	}
	return false, fmt.Errorf("Cannot convert %s to boolean. Valid input should be true, t, yes, y, false, f, no, n", s)
}
