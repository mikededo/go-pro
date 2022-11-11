package main

import (
	"fmt"
	"strconv"

	repository "github.com/mddg/go-taxi-events/internal/repository/taxi_manager"
	"github.com/mddg/go-taxi-events/internal/usecases"
)

func main() {
	var inputValue string
	tr := repository.NewMemoryTaxiManagerRepository()

	for inputValue != "x" {
		_, err := fmt.Scanf("%s", &inputValue)
		if err != nil {
			panic(err)
		}

		v, err := strconv.Atoi(inputValue)
		if err == nil {
			usecases.NewAssignTaxiUseCase(tr).AssignTaxiUseCase(v)
		} else if fmt.Sprint(inputValue) == "x" {
			return
		}
	}
}
