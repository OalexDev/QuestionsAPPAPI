package model

type Game struct {
	ID       int64  `json:"id,omitempty"`
	RoomID   int64  `json:"room,omitempty"`
	PlayerID int64  `json:"player,omitempty"`
	Text     string `json:"text,omitempty"`
}
