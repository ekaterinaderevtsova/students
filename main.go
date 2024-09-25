package main

import (
	"fmt"
	"students_app/student"
)

type ClassRoom [2]student.Student

func (cr ClassRoom) AddStudent() {
	// I need to change array to slice to make this method work!
}

func main() {
	var students ClassRoom
	students2 := ClassRoom{{1, "Jake", 2, [3]int{4, 4, 5}}, {2, "Emily", 4, [3]int{5, 3, 4}}}

	fmt.Println(students)
	fmt.Println(students2)
}
