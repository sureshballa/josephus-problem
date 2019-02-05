package main

import "fmt"

func main() {
	soldiers := constructSoldiers(1000000)
	finalSoldier := completeRound(soldiers)
	fmt.Println("Final soldier is ", finalSoldier)
}

func completeRound(soldiers []*Soldier) (ret int) {
	printSoldiers(soldiers)
	if len(soldiers) == 1 {
		return (*soldiers[0]).Index
	} else {
		var newLength int
		if len(soldiers)%2 == 0 {
			newLength = len(soldiers) / 2
		} else {
			newLength = (len(soldiers) / 2) + 1
		}
		newSoldiers := make([]*Soldier, newLength, newLength)
		for i, j := 1, 0; i < len(soldiers); i, j = i+2, j+1 {
			(*soldiers[i]).Status = false
			newSoldiers[j] = &Soldier{
				Index:  (*soldiers[i-1]).Index,
				Status: (*soldiers[i-1]).Status}
		}
		if len(soldiers)%2 != 0 {
			newSoldiers[len(newSoldiers)-1] = &Soldier{
				Index:  (*soldiers[len(soldiers)-1]).Index,
				Status: (*soldiers[len(soldiers)-1]).Status}

			newSoldiersRightShift := make([]*Soldier, newLength, newLength)
			newSoldiersRightShift[0] = newSoldiers[len(newSoldiers)-1]

			for i := 0; i+1 < len(newSoldiers); i++ {
				newSoldiersRightShift[i+1] = newSoldiers[i]
			}

			return completeRound(newSoldiersRightShift)
		} else {
			return completeRound(newSoldiers)
		}
	}
}

func constructSoldiers(maxCapacity int64) (soldiers []*Soldier) {
	soldiers1 := make([]*Soldier, maxCapacity, maxCapacity)
	for i := range soldiers1 {
		soldiers1[i] = &Soldier{
			Index:  i + 1,
			Status: true}
	}
	soldiers = soldiers1
	return
}

func printSoldiers(soldiers []*Soldier) {
	for i := range soldiers {
		fmt.Print(" { %d } ", (*soldiers[i]).Index)
	}
	fmt.Println()
}

type Soldier struct {
	Index  int
	Status bool
}

//type Soldiers []Soldiers
