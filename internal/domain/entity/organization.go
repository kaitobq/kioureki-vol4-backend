package entity

import "time"

type Organization struct {
	ID 	      uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
