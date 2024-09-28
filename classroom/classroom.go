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

func NewClassRoom(name string) *ClassRoom {
	return &ClassRoom{
		name:         name,
		studentsList: make([]student.Student, 0, 10),
		numStudents:  0,
	}
}

func (cr ClassRoom) GetNumStudents() int {
	return cr.numStudents
}

func (cr *ClassRoom) AddStudent(student student.Student) {
	cr.studentsList = append(cr.studentsList, student)
	cr.numStudents++
}

func (cr *ClassRoom) RemoveStudentById(id int) {
	for i, v := range cr.studentsList {
		if v.Id == id {
			indexOfStudentToRemove := i
			cr.studentsList = append(cr.studentsList[:indexOfStudentToRemove], cr.studentsList[indexOfStudentToRemove+1:]...)
			fmt.Println("Removed student: ", v.Id)
			return
		}
	}
	fmt.Println("Student not found!")
}

func (cr *ClassRoom) RemoveStudentByName(name string) {
	for i, v := range cr.studentsList {
		if v.Name == name {
			indexOfStudentToRemove := i
			cr.studentsList = append(cr.studentsList[:indexOfStudentToRemove], cr.studentsList[indexOfStudentToRemove+1:]...)
			fmt.Println("Removed student: ", v.Name)
			return
		}
	}
	fmt.Println("Student not found!")
}

func (cr ClassRoom) DisplayStudents() {
	for _, v := range cr.studentsList {
		fmt.Println("Student ", v.Id, ":")
		fmt.Println("name - ", v.Name, "; year - ", v.Year, "; scores - ", v.Scores)
	}
}
