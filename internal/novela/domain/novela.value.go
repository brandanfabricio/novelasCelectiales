package domain

type Novela struct {
	Id     string `db:"id"`
	Titulo string `db:"titulo"`
	Imagen string `db:"imagen"`
	Url    string `db:"url"`
	// Description string `db:"description"`
	Description string `db:"descripcion"`
	Paginas     string `db:"paginas"`
}
