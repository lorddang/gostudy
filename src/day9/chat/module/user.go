package module

type User struct {
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	sex      string    `json:"sex"`
	phone    string `json:"phone"`
}

func NewUser(id int, nickName, password string) *User  {
	return &User{
		Id: id,
		NickName:nickName,
		Password:password,
	}
}

func (u *User)Register()  {

}

