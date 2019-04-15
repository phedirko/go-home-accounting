package models

import "time"

type Transaction struct {
	ID 		  int32
	CreatedAt time.Time
	MonthId   int32
}