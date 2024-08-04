package domain

type UserRepository interface {
	FaindByEmail(id string) User
	FaindById(id int) User
	SaveUser(user User, channel chan error, okey chan int)
	FaindRol(id int)
}
