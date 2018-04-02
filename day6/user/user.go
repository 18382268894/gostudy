package user

type User struct {
	Name     string
	Age      uint8
	PhotoUrl string
}

func NewUser(name string, age uint8, photoUrl string) User {
	var user = User{
		Name:     name,
		Age:      age,
		PhotoUrl: photoUrl,
	}
	return user
}
