package quizz

import "testing"

var roundTest = []struct {
	s        Score
	roundOn  float64
	expected Score
}{
	{1, 1, 1},
	{0, 1, 0},
}

func TestRound(t *testing.T) {
	for _, tt := range roundTest {
		round := NewRound(tt.roundOn)

		if result := round(tt.s); result != tt.expected {
			t.Errorf("NewRound(%f)(%f) => %f, want %f", tt.roundOn, tt.s, result, tt.expected)
		}
	}
}
