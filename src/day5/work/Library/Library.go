package Library

import (
	"day5/work/Student"
	"day5/work/Book"
)

type Library struct {
	StudentList []Student.Student
	BookList []Book.Book
}
func (this *Library)AddBook(book Book.Book)  {
	if this.BookList == nil {
		this.BookList = make([]Book.Book, 1)
	}
	this.BookList = append(this.BookList, book)
}

func (this *Library)AddStudent(student Student.Student)  {
	if this.StudentList == nil {
		this.StudentList = make([]Student.Student, 1)
	}
	this.StudentList = append(this.StudentList, student)
}