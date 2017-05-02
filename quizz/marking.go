package quizz

type Marking struct {
	Question *Question
	Answer   string
}

type MarkingCollection struct {
	Markings []Marking
}
