package classroom

import (
	"fmt"
	"students_app/student"
)

type ClassRoom struct {
	name         string
	studentsList []student.Student
}

func (cr *ClassRoom) AddStudent(student student.Student) {
	cr.studentsList = append(cr.studentsList, student)
}

func (cr ClassRoom) DisplayStudents() {
	for i, v := range cr.studentsList {
		fmt.Println("Student ", i+1, ":")
		fmt.Println("name - ", v.Name, "; year - ", v.Year, "; scores - ", v.Scores)
	}
}
