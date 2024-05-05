package domain

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type UserRepository interface {
	FindByID() User
	FindByUsername(username string) User
}
