package main

import (
	"fmt"
	"math/rand"
)
func main() {
	var coins = [3]float32{0.05, 0.10, 0.25}
	var gift float32 = 20
	var monyeInPiggyBank float32 = 0
	for monyeInPiggyBank < gift {
		monyeInPiggyBank += coins[rand.Intn(3)]
		fmt.Printf("%05.2f \n", monyeInPiggyBank)
	}
}