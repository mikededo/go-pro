package entity

import "time"

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
		time.Sleep(time.Duration(duration) * time.Millisecond)
		done(t.Id)
	}()
}
