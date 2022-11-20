package repository

type TaxiManagerRepository interface {
	NotifyTaxiAssignToBus(int)
	ClearBusSubscription()
}
