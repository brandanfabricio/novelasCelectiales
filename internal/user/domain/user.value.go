package domain

type User struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Status   string `db:"status"`
	Password string `db:"password"`
	Email    string `db:"email"`
	RolId    string `db:"rol_id"`
}
