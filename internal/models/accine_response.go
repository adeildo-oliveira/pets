package models

import "time"

type VaccineResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"nome"`
	DateGiven time.Time `json:"dataAplicacao"`
	NextDue   time.Time `json:"proximaDose"`
}
