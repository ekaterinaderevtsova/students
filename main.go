package main

import (
	"students_app/classroom"
	"students_app/student"
)

func main() {
	var students classroom.ClassRoom
	students.AddStudent(student.Student{Id: students.GetNumStudents() + 1, Name: "Jake", Year: 4, Scores: []int{4, 4, 5}})
	students.AddStudent(student.Student{Id: students.GetNumStudents() + 1, Name: "Finn", Year: 3, Scores: []int{5, 3, 5}})
	students.AddStudent(student.Student{Id: students.GetNumStudents() + 1, Name: "Bubblegum", Year: 6, Scores: []int{5, 5, 5, 5, 5}})
	students.DisplayStudents()
}
