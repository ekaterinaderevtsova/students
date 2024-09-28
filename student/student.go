package student

type Student struct {
	Id     int
	Name   string
	Year   int
	Scores []int
}

func NewStudent(id int, name string, year int, scores []int) *Student {
	return &Student{
		Id:     id,
		Name:   name,
		Year:   year,
		Scores: scores,
	}
}
