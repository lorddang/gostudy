package Library

import "fmt"

type Book struct {
	Isbn string
	Name string
	Author string
	PublishDate string
	Num int
	ReadList []Student
}

func NewBook(isbn, name, author, pulishDate string, num int) *Book  {
	book := &Book{
		Isbn:isbn,
		Name:name,
		Author:author,
		PublishDate:pulishDate,
		Num:num,
	}
	return book

}

func (this *Book)BorrowBook(s *Student) bool {
	if this.Num <= 0 {
		return false
	}
	this.Num --
	fmt.Println(this.Num)
	if this.ReadList == nil {
		this.ReadList = make([]Student, 1)
		this.ReadList[0] = *s
		return true
	}
	this.ReadList = append(this.ReadList, *s)
	return true
}

func (this *Book)RebackBook(s *Student) bool {
	this.Num ++
	var sindex int
	for k, v := range this.ReadList{
		if v.Sid == s.Sid {
			sindex = k
			break
		}
	}
	this.ReadList = append(this.ReadList[:sindex], this.ReadList[sindex + 1:]...)
	fmt.Println(this.ReadList)
	fmt.Println(this)
	return true

}

