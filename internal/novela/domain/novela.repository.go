package domain

type NovelaRepository interface {
	FindAll() ([]Novela, error)
	SaveNovela(novela Novela)
	GetById(id string) Novela
	UpdateUrlNovela(novela_id string, url string)
}
