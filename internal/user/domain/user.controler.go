package domain

type CaseUserInt interface {
	FindUser(id int) interface{}
	SaveUser(user User) (interface{}, error)
	Login(user User) (interface{}, error)
}
