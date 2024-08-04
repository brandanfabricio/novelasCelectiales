package app

import (
	"Novelas/internal/user/domain"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type CaseApp struct {
	UserRepository domain.UserRepository
}

func NewCaseUser(userRepo domain.UserRepository) *CaseApp {
	return &CaseApp{UserRepository: userRepo}
}

func (c CaseApp) FindUser(id int) interface{} {

	data := c.UserRepository.FaindById(id)
	if data.Id == "" {

		return map[string]interface{}{
			"msj": "El usuario no existe",
		}
	}

	return data
}
func (c CaseApp) SaveUser(user domain.User) (interface{}, error) {
	//chanel
	chanel := make(chan error)
	chanel2 := make(chan int)
	defer close(chanel)
	defer close(chanel2)
	// validaciones

	// hash Password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Nose pudo hacer el hash del password")
		panic(err)
	}
	user.Password = string(hashPassword)

	go func() {
		c.UserRepository.SaveUser(user, chanel, chanel2)
	}()
	select {
	case err := <-chanel:
		return nil, err
	case okey := <-chanel2:
		return okey, nil
	}
}

func (c CaseApp) Login(user domain.User) (interface{}, error) {

	//   buscar usuario por mail

	isUser := c.UserRepository.FaindByEmail(user.Email)

	if isUser.Id == "" {
		return map[string]interface{}{
			"msj": "el password y o email es incorrecto",
		}, nil
	}
	// compara password
	comparePaass := bcrypt.CompareHashAndPassword([]byte(isUser.Password), []byte(user.Password))

	if comparePaass != nil {

		return map[string]interface{}{
			"msj": "el password y o email es incorrecto",
		}, nil
	}

	return isUser, nil

}
