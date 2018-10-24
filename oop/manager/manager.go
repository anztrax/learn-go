package manager

import "fmt"

type manager struct{
	firstName string
	lastName string
	totalLeaves int
	leavesTaken int
}

func New(firstName  string, lastName string, totalLeaves int, leavesTaken int) manager{
	e := manager{
		firstName,
		lastName,
		totalLeaves,
		leavesTaken,
	}
	return e
}

func(e manager) LeavesRemaining(){
	fmt.Printf("\n\n%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves- e.leavesTaken))
}