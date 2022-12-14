package usecases

import (
	"github.com/mddg/go-pro/go-taxi-event/internal/domain/taxi_manager/repository"
)

type NotifyTaxiAssignToBus struct {
	r repository.TaxiManagerRepository
}

func NewNotifyTaxiAssignToBusUseCase(r repository.TaxiManagerRepository) *NotifyTaxiAssignToBus {
	return &NotifyTaxiAssignToBus{r: r}
}

func (nt *NotifyTaxiAssignToBus) Execute(duration int) {
	nt.r.NotifyTaxiAssignToBus(duration)
}
