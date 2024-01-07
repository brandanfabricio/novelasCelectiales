package application

import (
	"Novelas/internal/novela/domain"
	"Novelas/internal/novela/domain/lib"
	"fmt"
	"log"
)

type CaseNovelaInt interface {
	FindAllNovela() []domain.Novela
	SaveNovela(url string) bool
}

type CaseNovela struct {
	NovelaRepository domain.NovelaRepository
}

func NewCaseNovela(NovelaRepo domain.NovelaRepository) *CaseNovela {

	return &CaseNovela{NovelaRepository: NovelaRepo}

}



func (c *CaseNovela) FindAllNovela() []domain.Novela {

	data, err := c.NovelaRepository.FindAll()

	if err != nil {
		log.Println("erro en el caso")

		panic(err)

	}

	return data

}

func (c *CaseNovela) SaveNovela(url string) bool {

	fmt.Println(url)
	novela := lib.ExtrarNovle(url)

	c.NovelaRepository.SaveNovela(novela)

	// fmt.Println(novela)

	if url != "" {
		return true

	}
	return false

}
