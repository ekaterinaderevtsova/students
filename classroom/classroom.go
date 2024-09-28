package classroom

import (
	"fmt"
	"students_app/student"
)

type ClassRoom struct {
	name         string
	studentsList []student.Student
	numStudents  int
}

func (cr ClassRoom) GetNumStudents() int {
	return cr.numStudents
}

func (cr *ClassRoom) AddStudent(student student.Student) {
	cr.studentsList = append(cr.studentsList, student)
	cr.numStudents++
}

func (cr ClassRoom) DisplayStudents() {
	for _, v := range cr.studentsList {
		fmt.Println("Student ", v.Id, ":")
		fmt.Println("name - ", v.Name, "; year - ", v.Year, "; scores - ", v.Scores)
	}
}
