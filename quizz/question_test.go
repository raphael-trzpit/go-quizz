package quizz

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestToto(t *testing.T) {
	q := NewTrueFalse("Jaune est une couleur", true)
	var b bytes.Buffer
	WriteQuestion(q, &b)
	answer := "f"
	score, _ := q.Score(answer)
	WriteQuestionSolution(q, &b)
	out, _ := ioutil.ReadAll(&b)
	t.Logf("%s", out)
	t.Logf("%s", score)
}
