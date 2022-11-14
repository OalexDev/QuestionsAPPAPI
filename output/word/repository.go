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

func (r *Repository) ReadeOneWord() ([]model.Word, error) {

	query := `select id, word from words ORDER BY random() 	LIMIT 1;`
	rows, err := r.DbConnection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ws := []model.Word{}

	for rows.Next() {
		w := model.Word{}
		err = rows.Scan(&w.ID, &w.Word)
		if err != nil {
			return nil, err
		}
		ws = append(ws, w)
	}

	return ws, nil

}
