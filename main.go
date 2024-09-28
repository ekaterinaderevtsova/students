package main

import (
	"students_app/classroom"
	"students_app/student"
)

func main() {
	var students *classroom.ClassRoom = classroom.NewClassRoom("classroom 1")
	students.AddStudent(*student.NewStudent(students.GetNumStudents()+1, "Jake", 4, []int{4, 4, 5}))
	students.AddStudent(*student.NewStudent(students.GetNumStudents()+1, "Finn", 3, []int{5, 3, 5}))
	students.AddStudent(*student.NewStudent(students.GetNumStudents()+1, "Bubblegum", 6, []int{5, 5, 5, 5, 5}))
	students.AddStudent(*student.NewStudent(students.GetNumStudents()+1, "Jane", 5, []int{4, 5, 4}))
	students.RemoveStudentById(4)
	students.DisplayStudents()
}
