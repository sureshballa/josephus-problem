package main

import (
	"fmt"

	. "soldiers"
)

func main() {
	soldiersCompledTypeObject := SoldiersComplexType{}
	constructSoldiersGeneric(&soldiersCompledTypeObject, 50)
	finalSoldier := CompleteRound(&soldiersCompledTypeObject)
	fmt.Println("Final live soldier is ", finalSoldier)
}

func constructSoldiersGeneric(soldierOperations SoldierOperations, maxCapacity int64) {
	soldierOperations.ConstructSoldiers(maxCapacity)
}
