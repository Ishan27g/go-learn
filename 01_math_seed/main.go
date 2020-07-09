package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGaurds := rand.Intn(100)

	leftSafely := rand.Intn(5)
	if isHeistOn {
		if eludedGaurds >= 50 {
			fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
		} else {
			isHeistOn = false
			fmt.Println("Plan a better disguise next time?")
		}
		openedVault := rand.Intn(100)
		if isHeistOn && openedVault >= 10 {
			fmt.Println("Grab and GO!")
		} else if isHeistOn {
			isHeistOn = false
			fmt.Println("Vault can't be opened!")
		}
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Heist Failed-0")
		case 1:
			isHeistOn = false
			fmt.Println("Heist Failed-1")
		case 2:
			isHeistOn = false
			fmt.Println("Heist Failed-2")
		case 3:
			isHeistOn = false
			fmt.Println("Heist Failed-3")
		default:
			fmt.Println("Start the getaway car!")
		}
		if isHeistOn {
			amtStolen := 10000 + rand.Intn(1000000)
			fmt.Println("Amount Stolen : ", amtStolen)
		}
	}
	fmt.Println(isHeistOn)
}
