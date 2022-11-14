package model

type Word struct {
	ID   int64  `json:"id" db:"id"`
	Word string `json:"word"`
}
