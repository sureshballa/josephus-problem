package main

import "fmt"

func main() {
	soldiersCompledTypeObject := SoldiersComplexType{}
	constructSoldiersGeneric(&soldiersCompledTypeObject, 50)
	finalSoldier := completeRound(&soldiersCompledTypeObject)
	fmt.Println("Final live soldier is ", finalSoldier)
}

func completeRound(soldiersCompledTypeObject *SoldiersComplexType) (ret int) {
	printSoldiersGeneric(soldiersCompledTypeObject)
	soldiersLocalReference := soldiersCompledTypeObject.Soldiers
	if len(soldiersLocalReference) == 1 {
		return (*soldiersLocalReference[0]).Index
	} else {
		var newLength int
		if len(soldiersLocalReference)%2 == 0 {
			newLength = len(soldiersLocalReference) / 2
		} else {
			newLength = (len(soldiersLocalReference) / 2) + 1
		}
		newSoldiers := make([]*Soldier, newLength, newLength)
		for i, j := 1, 0; i < len(soldiersLocalReference); i, j = i+2, j+1 {
			(*soldiersLocalReference[i]).Status = false
			newSoldiers[j] = soldiersLocalReference[i-1]
		}
		if len(soldiersLocalReference)%2 != 0 {
			newSoldiers[len(newSoldiers)-1] = soldiersLocalReference[len(soldiersLocalReference)-1]

			newSoldiersRightShift := make([]*Soldier, newLength, newLength)
			newSoldiersRightShift[0] = newSoldiers[len(newSoldiers)-1]

			for i := 0; i+1 < len(newSoldiers); i++ {
				newSoldiersRightShift[i+1] = newSoldiers[i]
			}

			return completeRound(&SoldiersComplexType{newSoldiersRightShift})
		} else {
			return completeRound(&SoldiersComplexType{newSoldiers})
		}
	}
}

type Soldier struct {
	Index  int
	Status bool
}

type SoldiersComplexType struct {
	Soldiers []*Soldier
}

func (soldiersCompledTypeObject *SoldiersComplexType) constructSoldiers(maxCapacity int64) {
	soldiersCompledTypeObject.Soldiers = make([]*Soldier, maxCapacity, maxCapacity)
	for i := range soldiersCompledTypeObject.Soldiers {
		soldiersCompledTypeObject.Soldiers[i] = &Soldier{
			Index:  i + 1,
			Status: true}
	}
}

func (soldiersCompledTypeObject *SoldiersComplexType) printSoldiers() {
	for i := range soldiersCompledTypeObject.Soldiers {
		fmt.Print(" { %d } ", (*soldiersCompledTypeObject.Soldiers[i]).Index)
	}
	fmt.Println()
}

type SoldierOperations interface {
	constructSoldiers(maxCapacity int64)
	printSoldiers()
}

func constructSoldiersGeneric(soldierOperations SoldierOperations, maxCapacity int64) {
	soldierOperations.constructSoldiers(maxCapacity)
}

func printSoldiersGeneric(soldierOperations SoldierOperations) {
	soldierOperations.printSoldiers()
}
