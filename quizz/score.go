package quizz

// Score is a float (usually between 0 and 1)
type Score float64

// Round a score.
type Round func(s Score) Score

// NewRound Create a round
func NewRound(roundOn float64) Round {
	return func(s Score) (r Score) {
		return Score(float64(s) / roundOn)
	}
}
