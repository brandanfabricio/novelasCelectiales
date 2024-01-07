package repository

import (
	"Novelas/database"
	"Novelas/internal/novela/domain"
	"fmt"
	"log"
)

type MysqlNovelaRepo struct {
}

func (sql *MysqlNovelaRepo) SaveNovela(novela domain.Novela) {

	conex, err := database.NewMysql()
	defer conex.Close()

	if err != nil {
		panic(err.Error())
	}

	query := "INSERT INTO novelas (id,titulo,imagen,url,descripcion,paginas) VALUES (?, ?,?,?,?,?)"

	// id	varchar(255)	NO	PRI	NULL
	// titulo	varchar(255)	YES		NULL
	// imagen	varchar(255)	YES		NULL
	// url	varchar(255)	YES		NULL
	// descripcion	text	YES		NULL
	// paginas	varchar(255)	YES		NULL

	// Ejecutar la consulta de inserción
	result, err := conex.Exec(query, novela.Id, novela.Titulo, novela.Imagen, novela.Url, novela.Description, novela.Paginas)

	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Se insertó una fila con ID %d\n", lastInsertID)

}

func (sql *MysqlNovelaRepo) FindAll() ([]domain.Novela, error) {

	conex, err := database.NewMysql()
	if err != nil {

		panic(err.Error())
	}

	defer conex.Close()
	data, err := conex.Query("select * from novelas")

	if err != nil {
		log.Println("conex")

		return nil, err

	}

	novela := domain.Novela{}
	listNovelas := []domain.Novela{}

	for data.Next() {

		var id string
		var titulo string
		var imagen string
		var url string
		var description string
		var paginas string

		err = data.Scan(&id, &titulo, &imagen, &url, &description, &paginas)
		if err != nil {
			log.Println("erro peticion")

			panic(err.Error())
		}

		novela.Id = id
		novela.Titulo = titulo
		novela.Imagen = imagen
		novela.Url = url
		novela.Description = description
		novela.Paginas = paginas

		listNovelas = append(listNovelas, novela)
	}

	return listNovelas, nil

}

func (sql *MysqlNovelaRepo) GetById(id string) domain.Novela {

	conex, err := database.NewMysql()
	defer conex.Close()

	if err != nil {
		panic(err.Error())
	}

	var novela domain.Novela

	err = conex.Get(&novela, "select * from novelas where id = ?", id)

	if err != nil {
		// log.Fatalln("sdadasdasdasdasdasdasdasdasasdadsdas")
		log.Println(err)
		panic(err.Error())
	}

	return novela

}

func (sql *MysqlNovelaRepo) UpdateUrlNovela(novela_id string, url string) {

	conex, err := database.NewMysql()
	defer conex.Close()

	if err != nil {
		panic(err.Error())
	}

	consulta := `
	UPDATE novelas
	SET url = :url
	WHERE id = :id
`

	// Valores de parámetros con nombres
	// Valores de parámetros con nombres
	parametros := map[string]interface{}{
		"url": url,
		"id":  novela_id,
	}

	// Ejecuta la consulta nombrada
	resultado, err := conex.NamedExec(consulta, parametros)

	if err != nil {
		log.Fatal(err)
	}

	// Verifica la cantidad de filas afectadas
	filasAfectadas, err := resultado.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Filas afectadas: %d\n", filasAfectadas)
}
