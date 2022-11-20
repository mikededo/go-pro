package repository

import (
	bus "github.com/mddg/go-pro/go-taxi-event/internal/domain/bus/repository"
	"github.com/mddg/go-pro/go-taxi-event/internal/domain/taxi_manager/entity"
)

type MemoryTaxiManagerRepository struct {
	tm *entity.TaxiManager
	b  bus.Bus
}

var repoInstance *MemoryTaxiManagerRepository

func NewMemoryTaxiManagerRepository(b bus.Bus) *MemoryTaxiManagerRepository {
	if repoInstance == nil {
		repoInstance = &MemoryTaxiManagerRepository{
			tm: entity.NewTaxiManager(10),
			b:  b,
		}

		b.Subscribe("assign_taxi", repoInstance.assignPassengerToTaxi)
	}

	return repoInstance
}

func (m *MemoryTaxiManagerRepository) NotifyTaxiAssignToBus(d int) {
	m.b.Notify("assign_taxi", d)
}

func (m *MemoryTaxiManagerRepository) ClearBusSubscription() {
	m.b.Clear("assign_taxi")
}

func (m *MemoryTaxiManagerRepository) assignPassengerToTaxi(d int) error {
	return m.tm.AssignPassenger(d)
}
