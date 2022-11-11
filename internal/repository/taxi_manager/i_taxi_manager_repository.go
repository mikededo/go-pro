package repository

type TaxiManagerRepository interface {
	AssignPassengerToTaxi(int) error
}
