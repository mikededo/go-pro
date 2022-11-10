package entity

import "errors"

type TaxiManager struct {
	taxis     map[int]*Taxi
	freeQueue []int
}

func NewTaxiManager(taxis int) *TaxiManager {
	t := make(map[int]*Taxi, taxis)
	a := make([]int, taxis)
	for i := 1; i <= taxis; i++ {
		t[i] = NewTaxi(i)
		a[i-1] = i
	}

	return &TaxiManager{
		taxis:     t,
		freeQueue: a,
	}
}

func (t *TaxiManager) AssignPassenger(duration int) error {
	f, err := findAndPopFirstFreeTaxi(t)
	if err != nil {
		return err
	}

	taxi := t.taxis[f]
	taxi.StartTrip(duration, t.FreePassenger)

	return nil
}

func (t *TaxiManager) FreePassenger(id int) {
	t.freeQueue = append(t.freeQueue, id)
}

func findAndPopFirstFreeTaxi(tm *TaxiManager) (int, error) {
	if len(tm.freeQueue) == 0 {
		return 0, errors.New("no available taxis")
	}

	first := tm.freeQueue[0]
	tm.freeQueue = tm.freeQueue[1:]
	return first, nil
}
