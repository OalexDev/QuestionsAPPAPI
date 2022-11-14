package model

import "time"

type Rooms struct {
	ID     int64     `json:"id" db:"id"`
	Room   int64     `json:"room" db:"room"`
	Create time.Time `json:"create" db:"create_at"`
}
