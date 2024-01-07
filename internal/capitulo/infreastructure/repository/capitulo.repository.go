package repository

import (
	"Novelas/database"
	"Novelas/internal/capitulo/domain"
	"fmt"
	"log"
	"sync"
)

var dbMutex sync.Mutex

type MysqlCApituloRepo struct {
}

func (sql *MysqlCApituloRepo) GetCapituloById(id string) []domain.Capitulo {

	conex, err := database.NewMysql()
	defer conex.Close()

	if err != nil {
		panic(err.Error())
	}
	defer conex.Close()

	resultadoPorPagina := 10

	numeroPagina := 1

	// Calcular el OFFSET basado en el número de página y resultados por página
	offset := (numeroPagina - 1) * resultadoPorPagina

	// Consulta SQL con LIMIT y OFFSET para la paginación
	query := fmt.Sprintf("SELECT id,titulo,Ncap,novelaId FROM capitulos where novelaId='%s' ORDER BY Ncap ASC LIMIT %d OFFSET %d", id, resultadoPorPagina, offset)

	// Realizar la consulta y escanear los resultados en una estructura de slice
	var capitulos []domain.Capitulo

	err = conex.Select(&capitulos, query)

	if err != nil {
		log.Fatal(err)
	}

	return capitulos
}
func (sql *MysqlCApituloRepo) GetCapituloPaginated(id string, page int) []domain.Capitulo {

	conex, err := database.NewMysql()
	defer conex.Close()

	if err != nil {
		panic(err.Error())
	}
	defer conex.Close()

	resultadoPorPagina := 10

	numeroPagina := page

	// Calcular el OFFSET basado en el número de página y resultados por página
	offset := (numeroPagina - 1) * resultadoPorPagina

	// Consulta SQL con LIMIT y OFFSET para la paginación
	query := fmt.Sprintf("SELECT id,titulo,Ncap,novelaId FROM capitulos  WHERE novelaId= '%s'  LIMIT %d OFFSET %d", id, resultadoPorPagina, offset)

	// Realizar la consulta y escanear los resultados en una estructura de slice
	var capitulos []domain.Capitulo

	err = conex.Select(&capitulos, query)

	if err != nil {
		log.Fatal(err)
	}

	return capitulos
}

func (sql *MysqlCApituloRepo) GetPage(id string, page int) []domain.Capitulo {

	conex, err := database.NewMysql()
	defer conex.Close()

	if err != nil {
		panic(err.Error())
	}
	defer conex.Close()

	var capitulo []domain.Capitulo

	query := fmt.Sprintf("SELECT id,titulo,Ncap,novelaId FROM capitulos WHERE novelaId ='%s' AND Ncap  LIKE '%%%d%%'", id, page)

	err = conex.Select(&capitulo, query)

	if err != nil {
		log.Println("Erro en la consulta")
		log.Fatal(err)

	}

	return capitulo
}

func (sql *MysqlCApituloRepo) FindCapitulo(id int, novela_id string, numero_cap int) domain.Capitulo {

	conex, err := database.NewMysql()

	if err != nil {
		log.Println("Error de conexcion")
		log.Panicln(err)
	}
	defer conex.Close()

	var capitulo domain.Capitulo

	err = conex.Get(&capitulo, "SELECT id,cap,Ncap,titulo,novelaId,contenido FROM capitulos where novelaId= ? and (id = ? or Ncap =?)", novela_id, id, numero_cap)
	// (Ncap = 2 or id=213 )
	if err != nil {

		log.Println("error de consulta")
		log.Panicln(err)
	}
	return capitulo

}

func (sql *MysqlCApituloRepo) AddCapitulo(data []domain.Capitulo) {
	// guardar novela
	conex, err := database.NewMysql()

	if err != nil {
		log.Println("Error de conexcion")
		log.Panicln(err)
	}
	defer conex.Close()

	// trasaccion

	transax, err := conex.Begin()

	if err != nil {
		log.Fatal()
		log.Fatal("Error al crear la transaxsion", err)
	}

	defer func() {
		if p := recover(); p != nil {
			transax.Rollback()
			log.Println("erro en la ejecucion")
			panic(p)
		} else if err != nil {
			transax.Rollback()
		} else {
			// log.Println("Comit echo ")
			fmt.Println("Inserción exitosa")

			transax.Commit()
		}
	}()
	stmt, err := transax.Prepare("INSERT INTO capitulos (cap,titulo,contenido,Ncap,novelaId) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Fatal("Error al preparar la sentencia", err)
		return
	}
	defer stmt.Close()
	// bucle de insersion
	for _, item := range data {
		// log.Println("Intentando insertar capítulo:", item.Ncap)
		dbMutex.Lock()
		_, err := stmt.Exec(item.Cap, item.Titulo, item.Contenido, item.Ncap, item.NovelaId)
		dbMutex.Unlock()

		if err != nil {
			// Manejar el error según tus necesidades
			log.Println("Error al insertar elemento:", err)
			// Hacer un rollback y salir del bucle
			transax.Rollback()
			return
		}
	}

}
