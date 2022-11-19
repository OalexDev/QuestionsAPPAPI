package room

import (
	"database/sql"
	"time"

	"github.com/OalexDev/QuestionsAPPAPI/app/room/model"
	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"
	_ "github.com/lib/pq"
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

func (r *Repository) EntryRoom(room *model.Rooms) (int, error) {

	query := `INSERT INTO public.roomuser (roomid,userid) VALUES ($1,$2) RETURNING id;`
	id := 0
	err := r.DbConnection.QueryRow(query, room.ID, room.UserID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *Repository) MakeRoom() (int64, error) {
	now := time.Now()
	idRoom := now.Unix()
	query := `INSERT INTO public.room (room)VALUES ($1) RETURNING id;`
	id := 0
	err := r.DbConnection.QueryRow(query, idRoom).Scan(&id)
	if err != nil {
		return 0, err
	}

	return idRoom, nil

}
