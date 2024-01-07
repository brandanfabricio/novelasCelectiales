package domain

type Capitulo struct {
	Id        int    `db:"id"`
	Cap       string `db:"cap"`
	Titulo    string `db:"titulo"`
	Contenido string `db:"contenido"`
	Ncap      int    `db:"Ncap"`
	NovelaId  string `db:"novelaId"`
}
