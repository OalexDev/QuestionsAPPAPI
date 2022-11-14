package room

import (
	"database/sql"

	"github.com/Five-Series/questions/app/room/model"
	"github.com/Five-Series/questions/infra/environment"
)

type Repository struct {
	env          *environment.Environment
	DbConnection *sql.DB
}

func NewRepository(env *environment.Environment, conn *sql.DB) *Repository {
	return &Repository{
		env:          env,
		DbConnection: conn}
}

func (r *Repository) GetRooms() ([]model.Rooms, error) {

	query := `select id,room, created_at from room where alive = true`
	rows, err := r.DbConnection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ros := []model.Rooms{}

	for rows.Next() {
		r := model.Rooms{}
		err = rows.Scan(&r.ID, &r.Room, &r.Create)
		if err != nil {
			return nil, err
		}
		ros = append(ros, r)
	}

	return ros, nil

}
