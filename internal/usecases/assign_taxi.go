package usecases

import (
	"github.com/mddg/go-taxi-events/internal/repository/taxi_manager"
)

type AssignTaxiUseCase struct {
	r repository.TaxiManagerRepository
}

func NewAssignTaxiUseCase(r repository.TaxiManagerRepository) *AssignTaxiUseCase {
	return &AssignTaxiUseCase{r: r}
}

func (at *AssignTaxiUseCase) AssignTaxiUseCase(duration int) {
	at.r.AssignPassengerToTaxi(duration)
}
