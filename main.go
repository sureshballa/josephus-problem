package main

import "fmt"

func main() {
	soldiers := Soldiers{}
	constructSoldiersGeneric(&soldiers, 100)
	finalSoldier := completeRound(&soldiers)
	fmt.Println("Final soldier is ", finalSoldier)
}

func completeRound(soldiers *Soldiers) (ret int) {
	printSoldiersGeneric(soldiers)
	soldiersLocalReference := soldiers.Soldiers
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

			return completeRound(&Soldiers{newSoldiersRightShift})
		} else {
			return completeRound(&Soldiers{newSoldiers})
		}
	}
}

type Soldier struct {
	Index  int
	Status bool
}

type Soldiers struct {
	Soldiers []*Soldier
}

func (soldiers *Soldiers) constructSoldiers(maxCapacity int64) {
	soldiers.Soldiers = make([]*Soldier, maxCapacity, maxCapacity)
	for i := range soldiers.Soldiers {
		soldiers.Soldiers[i] = &Soldier{
			Index:  i + 1,
			Status: true}
	}
}

func (soldiers *Soldiers) printSoldiers() {
	for i := range soldiers.Soldiers {
		fmt.Print(" { %d } ", (*soldiers.Soldiers[i]).Index)
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
