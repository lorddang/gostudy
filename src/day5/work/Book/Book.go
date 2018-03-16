package Book

import "day5/work/Student"

type Book struct {
	Isbn string
	Name string
	Author string
	PublishDate string
	Num int
	ReadList []Student.Student
}

