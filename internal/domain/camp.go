package domain

type Camp struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type CampRepository interface {
	FindByID(campID int64) *Camp
	FindByName(name string) *Camp
	FindAllCamp() *[]Camp
}

type CampService interface {
}
