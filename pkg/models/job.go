package models

import "time"

type Job struct {
	Uuid      string
	Message   string
	Status    int
	CreatedAt time.Time
}
