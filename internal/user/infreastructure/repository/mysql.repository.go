package repository

import (
	"Novelas/database"
	"Novelas/internal/user/domain"
	"fmt"
	"log"
)

type UserMysql struct {
}

func (qry UserMysql) FaindById(id int) domain.User {

	conex, err := database.NewMysql()

	if err != nil {
		fmt.Println("error de consulta")
		fmt.Println(err)
	}
	defer conex.Close()

	var user domain.User

	query := fmt.Sprintf("Select * from users where id = '%d'", id)

	err = conex.Get(&user, query)

	if err != nil {

		fmt.Println("error de consulta")
		fmt.Println(err)
	}
	return user
}

func (qry UserMysql) SaveUser(user domain.User, channel chan error, okey chan int) {
	// guardar novela
	conex, err := database.NewMysql()
	if err != nil {
		log.Println("Error de conexcion")
		log.Panicln(err)
	}
	defer conex.Close()
	query := "INSERT into users (name,email,password) VALUE (?,?,?);"
	result, err := conex.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		channel <- err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	okey <- 201

	fmt.Printf("Se insertó una fila con ID %d\n", lastInsertID)

}
func (qry UserMysql) FaindRol(id int) {

}
func (qry UserMysql) FaindByEmail(email string) domain.User {

	return domain.User{}
}
