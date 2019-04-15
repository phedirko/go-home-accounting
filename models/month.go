package models

import "time"

type Month struct {
	ID 		   int32     `json:"id"`
	StartedAt  time.Time `json:"startedAt"`			//read about time
	FinishedAt time.Time `json:"finishedAt"`
	Balance	   float64	 `json:"balance"`
}
