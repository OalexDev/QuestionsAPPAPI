package word

import (
	"database/sql"

	model "github.com/OalexDev/QuestionsAPPAPI/app/word/model"
	"github.com/OalexDev/QuestionsAPPAPI/infra/environment"
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

func (r *Repository) ReadeOneNewWord(room int64) (model.Word, error) {

	w := model.Word{}

	query := `SELECT 
			w.id, 
			w.word  
		FROM words w
		WHERE w.id not in (
			SELECT DISTINCT wordid  
			FROM roomwords 
			WHERE roomid = $1)
		ORDER BY random() LIMIT 1;`

	rows, err := r.DbConnection.Query(query, room)
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

func (r *Repository) ReadWord(room int64) (model.Word, error) {

	w := model.Word{}

	query := `SELECT id, word FROM  words WHERE id IN (
		SELECT wordid  FROM roomwords rw
		WHERE roomid = $1
		ORDER BY created_at  DESC 
		LIMIT 1 );`

	rows, err := r.DbConnection.Query(query, room)
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
