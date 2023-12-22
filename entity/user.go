package entity

type User struct {
	Id string	`json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age int16 `json:"age"`
	Email string `json:"email"`
	About string `json:"about"`
	IconColor string `json:"iconColor"`
}