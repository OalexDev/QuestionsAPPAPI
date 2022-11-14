package word

import (
	"database/sql"

	model "github.com/Five-Series/questions/app/word/model"
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

func (r *Repository) ReadeOneWord() (model.Word, error) {

	w := model.Word{}

	query := `select id, word from words ORDER BY random() 	LIMIT 1;`
	rows, err := r.DbConnection.Query(query)
	if err != nil {
		return w, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&w.ID, &w.Word)
		if err != nil {
			return w, err
		}
	}

	return w, nil

}

func (r *Repository) InsertWordToRoom(roomid, wordId int64) (int, error) {

	query := `INSERT INTO public.roomwords (roomid,wordid)	VALUES ($1,$2) RETURNING id;	`
	id := 0
	err := r.DbConnection.QueryRow(query, roomid, wordId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
