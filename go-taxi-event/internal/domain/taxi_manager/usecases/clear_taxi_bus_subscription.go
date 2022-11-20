package usecases

import (
	"github.com/mddg/go-pro/go-taxi-event/internal/domain/taxi_manager/repository"
)

type ClearTaxiBusSubscriptionUseCase struct {
	r repository.TaxiManagerRepository
}

func NewClearTaxiBusSubscriptionUseCase(r repository.TaxiManagerRepository) *ClearTaxiBusSubscriptionUseCase {
	return &ClearTaxiBusSubscriptionUseCase{r: r}
}

func (ct *ClearTaxiBusSubscriptionUseCase) Execute() {
	ct.r.ClearBusSubscription()
}
