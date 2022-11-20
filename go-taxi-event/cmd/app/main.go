package main

import (
	"fmt"
	"strconv"

	bRepository "github.com/mddg/go-pro/go-taxi-event/internal/domain/bus/repository"
	tRepository "github.com/mddg/go-pro/go-taxi-event/internal/domain/taxi_manager/repository"
	"github.com/mddg/go-pro/go-taxi-event/internal/domain/taxi_manager/usecases"
)

func main() {
	eb := bRepository.NewEventBus(1)
	tr := tRepository.NewMemoryTaxiManagerRepository(eb)

	var inputValue string
	for inputValue != "x" {
		_, err := fmt.Scanf("%s", &inputValue)
		if err != nil {
			panic(err)
		}

		v, err := strconv.Atoi(inputValue)
		if err == nil {
			usecases.NewNotifyTaxiAssignToBusUseCase(tr).Execute(v)
		} else if fmt.Sprint(inputValue) == "x" {
			usecases.NewClearTaxiBusSubscriptionUseCase(tr).Execute()
			return
		}
	}
}
