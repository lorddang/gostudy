package Library

type Student struct {
	Sid string
	Name string
	Grade string
	IdCard string
	Sex string
	BookList []Book
}

func NewStudent(sid, name, grade, idcard, sex string) *Student  {
	stu := &Student{
		Sid:sid,
		Name:name,
		Grade:grade,
		IdCard:idcard,
		Sex:sex,
	}
	return stu
}

func (this Student)BorrowBook(b Book)  {
	if this.BookList == nil {
		this.BookList = make([]Book, 1)
	}
	this.BookList = append(this.BookList, b)
}

func (this Student) RebackBook(b Book)  {
	var sindex int
	for k, v := range this.BookList{
		if v.Isbn == b.Isbn {
			sindex = k
			break
		}
	}
	this.BookList = append(this.BookList[:sindex], this.BookList[sindex + 1:]...)
}