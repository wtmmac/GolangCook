package models

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int64  `json:"age"`
}

func (User) TableName() string {
	return "user"
}

type Student struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	School string `json:"school"`
	Age    int64  `json:"age"`
}

func (Student) TableName() string {
	return "student"
}
