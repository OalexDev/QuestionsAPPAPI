package model

import "time"

type Rooms struct {
	ID     int64     `json:"id,omitempty" db:"id"`
	Room   int64     `json:"room,omitempty" db:"room"`
	UserID int64     `json:"user_id,omitempty" db:"user_id"`
	Create time.Time `json:"create,omitempty" db:"create_at"`
}
