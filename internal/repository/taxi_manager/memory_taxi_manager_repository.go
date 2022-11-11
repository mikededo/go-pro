package repository

import "github.com/mddg/go-taxi-events/internal/domain"

type MemoryTaxiManagerRepository struct {
	tm *domain.TaxiManager
}

var repoInstance *MemoryTaxiManagerRepository

func NewMemoryTaxiManagerRepository() *MemoryTaxiManagerRepository {
	if repoInstance == nil {
		repoInstance = &MemoryTaxiManagerRepository{
			tm: domain.NewTaxiManager(10),
		}
	}

	return repoInstance
}

func (m *MemoryTaxiManagerRepository) AssignPassengerToTaxi(d int) error {
	return m.tm.AssignPassenger(d)
}
