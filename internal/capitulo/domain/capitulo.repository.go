package domain

type CapituloRepository interface {
	GetCapituloById(id string) []Capitulo
	GetCapituloPaginated(id string, page int) []Capitulo
	GetPage(id string, page int) []Capitulo
	FindCapitulo(capitulo_id int, novela_id string, numero_cap int) Capitulo
	AddCapitulo(data []Capitulo)
}
