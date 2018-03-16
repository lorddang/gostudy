package Student

import (
	"day5/work/Book"
)

type Student struct {
	Sid string
	Name string
	Grade string
	IdCard string
	Sex string
	BookList []Book.Book
}
