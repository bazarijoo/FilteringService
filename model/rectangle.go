package model

import "time"

type Rectangle struct {
	X         int
	Y         int
	Width     int
	Height    int
	CreatedAt time.Time
}
