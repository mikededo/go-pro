package entity

import (
	"fmt"
	"time"
)

type Taxi struct {
	Id     int
	InTrip bool
}

func NewTaxi(id int) *Taxi {
	return &Taxi{
		Id:     id,
		InTrip: false,
	}
}

func (t *Taxi) StartTrip(duration int, done func(int)) {
	go func() {
		fmt.Printf("[Taxi %d]: Started trip - duration: %d\n", t.Id, duration)
		time.Sleep(time.Duration(duration) * time.Millisecond)
		done(t.Id)
		fmt.Printf("[Taxi %d]: Finished trip\n", t.Id)
	}()
}
