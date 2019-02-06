package main

import (
	"fmt"
	"sync"

	. "soldiers"
)

func main() {
	
	var wg sync.WaitGroup
	operationStatusChannel := make(chan OperationStatus, 10)

	operations := []OperationStatus{
		{50, false, 0},
		{100, false, 0},
		{200, false, 0},
		{300, false, 0},
		{400, false, 0},
		{500, false, 0},
		{1000, false, 0},
		{10000, false, 0},
		{100000, false, 0},
	}

	for _, operation := range operations {
		// Increment the WaitGroup counter.
		wg.Add(1)
		go func(operation OperationStatus) {
			operation.solve()
			operationStatusChannel <- operation
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
		}(operation)
	}

	wg.Wait()
	close(operationStatusChannel)

	for received := range operationStatusChannel {
		fmt.Printf("For %d, final live soldier is %d\n", received.Number, received.LiveNumber)
	}
	
}

func (operationStatus *OperationStatus) solve() {
	soldiersCompledTypeObject := SoldiersComplexType{}
	constructSoldiersGeneric(&soldiersCompledTypeObject, operationStatus.Number)
	operationStatus.LiveNumber = CompleteRound(&soldiersCompledTypeObject)
}

func constructSoldiersGeneric(soldierOperations SoldierOperations, maxCapacity int64) {
	soldierOperations.ConstructSoldiers(maxCapacity)
}

type OperationStatus struct {
	Number int64
	Status bool
	LiveNumber	int
}
