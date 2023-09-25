package server

import (
	"time"
)

type Flower struct {
	Id   string
	Name string
}

type FlowerDetail struct {
	Id          string
	Name        string
	Period      int
	LastWatered time.Time
}
