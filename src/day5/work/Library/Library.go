package Library

import (
	"fmt"
)

type Library struct {
	StudentList []Student
	BookList []Book
}
func (this *Library)AddBook(book Book)  {
	if this.BookList == nil {
		this.BookList = make([]Book, 1)
		this.BookList[0] = book
		return
	}
	this.BookList = append(this.BookList, book)
}

func (this *Library)AddStudent(student Student)  {
	if this.StudentList == nil {
		this.StudentList = make([]Student, 1)
		this.StudentList[0] = student
		return
	}
	this.StudentList = append(this.StudentList, student)
}

func (this *Library)ShowBooks()  {
	fmt.Println("[编号] \t\t isbn \t\t name \t\t author \t\t publishDate \t\t num  \n")
	for k, v  := range this.BookList{
		fmt.Printf("[%d] \t\t %s \t\t %s \t\t %s \t\t %s \t\t %d \n", k, v.Isbn, v.Name, v.Author, v.PublishDate, v.Num)
	}

}

func (this *Library) ShowStudent()  {

	fmt.Println("[编号] \t\t 学号 \t\t 姓名 \t\t 身份证号 \t\t 年级 \t\t 性别\n\t" )
	for k, v  := range this.StudentList{
		fmt.Printf("[%d] \t\t %s \t\t %s \t\t %s \t\t %s \t\t %s\n", k, v.Sid, v.Name, v.IdCard, v.Grade, v.Sex)
	}
}

func (this *Library)BorrowBook()  {
	var (
		bindex int
		sindex int
	)
	fmt.Print("请输入图书编号: ")
	fmt.Scanf("%d\n", &bindex)
	if bindex >= len(this.BookList){
		fmt.Println("图书不存在")
		return
	}
	fmt.Print("请输入学生编号: ")
	fmt.Scanf("%d\n", &sindex)
	if sindex >= len(this.StudentList) {
		fmt.Println("学生不存在")
		return
	}
	s := this.StudentList[sindex]
	b := this.BookList[bindex]
	if b.BorrowBook(s){
		s.BorrowBook(b)
	}
	
}

func (this *Library)RebackBook(s Student, b Book)  {
	if b.RebackBook(s){
		s.RebackBook(b)
	}
}



func (this *Library)Open()  {
	this.CmdOption()
}

func ParseBookInfo() (book Book)  {
	var (
		isbn string
		name string
		author string
		publishDate string
		num int

	)
	fmt.Print("请输入书isbn: ")
	fmt.Scanf("%s\n", &isbn)
	fmt.Print("请输入书名: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入书作者: ")
	fmt.Scanf("%s\n", &author)
	fmt.Print("请输入书出版日期: ")
	fmt.Scanf("%s\n", &publishDate)
	fmt.Print("请输入书本数量: ")
	fmt.Scanf("%d\n", &num)

	book = Book{
		Name:name,
		Isbn:isbn,
		Author:author,
		PublishDate:publishDate,
		Num:num,
	}
	return
}

func ParseStudentInfo() (stu Student) {
	var (
		sid string
		name string
		idcard string
		sex string
		grade string
	)
	fmt.Print("请输入学号: ")
	fmt.Scanf("%s\n", &sid)
	fmt.Print("请输入姓名: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入身份证: ")
	fmt.Scanf("%s\n", &idcard)
	fmt.Print("请输入性别: ")
	fmt.Scanf("%s\n", &sex)
	fmt.Print("请输入年级: ")
	fmt.Scanf("%s\n", &grade)
	stu = Student{
		Sid:sid,
		Name:name,
		Sex:sex,
		Grade:grade,
	}
	return
}

func (this Library)echoCmdInfo()  {
	info := `
[0] 		addbook     		添加图书
[1] 		addstudent 			添加学生
[2] 		showbook 			展示图书信息
[3] 		showstudent			展示学生信息

`
fmt.Println(info)

}

func (this Library)CmdOption()  {
	for {
		this.echoCmdInfo()
		var option string
		fmt.Scanf("%s\n", &option)
		switch option {
		case "addbook":
			this.AddBook(ParseBookInfo())
		case "addstudent":
			this.AddStudent(ParseStudentInfo())
		case "showbook":
			this.ShowBooks()
		case "showstudent":
			this.ShowStudent()
		case "quit":
			return
		}
	}
}