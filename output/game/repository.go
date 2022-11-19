package game

import (
	"database/sql"
	"fmt"

	"github.com/OalexDev/QuestionsAPPAPI/app/game/model"
	modelWord "github.com/OalexDev/QuestionsAPPAPI/app/word/model"
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

func (r *Repository) InsertMaessage(g *model.Game) error {

	query := `INSERT INTO public.gameplay (roomid,playerid,message)	VALUES ($1, $2,$3) RETURNING id;`
	id := 0
	err := r.DbConnection.QueryRow(query, g.RoomID, g.PlayerID, g.Text).Scan(&id)
	if err != nil {
		fmt.Println("deu erro ")
		fmt.Println(err)
		return err
	}

	return nil

}

func (r *Repository) GetWordByID(id int64) (*string, error) {

	w := modelWord.Word{}
	query := `SELECT wordid AS id FROM roomwords r WHERE roomid = $1 ORDER BY created_at  DESC LIMIT 1;`
	rows, err := r.DbConnection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&w.ID)
		if err != nil {
			return nil, err
		}
	}

	query = `SELECT id, word FROM words  WHERE id = $1 LIMIT 1;`
	rows, err = r.DbConnection.Query(query, w.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&w.ID, &w.Word)
		if err != nil {
			return nil, err
		}
	}

	return &w.Word, nil

}

func (r *Repository) GetGameMessages(id int64) ([]model.Game, error) {

	result := []model.Game{}
	query := `SELECT id, roomid, playerid, message  FROM gameplay g WHERE roomid = $1 ORDER BY created_at ASC;`
	rows, err := r.DbConnection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		g := model.Game{}
		err = rows.Scan(&g.ID, &g.RoomID, &g.PlayerID, &g.Text)
		if err != nil {
			return nil, err
		}
		result = append(result, g)
	}

	return result, nil

}
