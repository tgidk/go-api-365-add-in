package resources

import (
	"time"
)

type Quote struct {
	Price float64   `json:"price"`
	Date  time.Time `json:"date"`
}
