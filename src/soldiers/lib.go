package soldiers

import "fmt"

type Soldier struct {
	Index  int
	Status bool
}

type SoldiersComplexType struct {
	Soldiers []*Soldier
}

type SoldierOperations interface {
	ConstructSoldiers(maxCapacity int64)
	PrintSoldiers()
}

func (soldiersComplexTypeObject *SoldiersComplexType) ConstructSoldiers(maxCapacity int64) {
	soldiersComplexTypeObject.Soldiers = make([]*Soldier, maxCapacity, maxCapacity)
	for i := range soldiersComplexTypeObject.Soldiers {
		soldiersComplexTypeObject.Soldiers[i] = &Soldier{
			Index:  i + 1,
			Status: true}
	}
}

func (soldiersComplexTypeObject *SoldiersComplexType) PrintSoldiers() {
	for i := range soldiersComplexTypeObject.Soldiers {
		fmt.Print(" { %d } ", (*soldiersComplexTypeObject.Soldiers[i]).Index)
	}
	fmt.Println()
}

func printSoldiersGeneric(soldierOperations SoldierOperations) {
	soldierOperations.PrintSoldiers()
}

func CompleteRound(soldiersComplexTypeObject *SoldiersComplexType) (ret int) {
	//printSoldiersGeneric(soldiersCompledTypeObject)
	soldiersLocalReference := soldiersComplexTypeObject.Soldiers
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

			return CompleteRound(&SoldiersComplexType{newSoldiersRightShift})
		} else {
			return CompleteRound(&SoldiersComplexType{newSoldiers})
		}
	}
}
