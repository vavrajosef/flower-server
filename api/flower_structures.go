package server

import (
	"time"
)

type Flower struct {
	Id   int
	Name string
}

type FlowerDetail struct {
	Id          int
	Name        string
	Period      int
	LastWatered time.Time
}
