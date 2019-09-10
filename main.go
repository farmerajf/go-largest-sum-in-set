package main

import "fmt"

func main() {
	myset := []int{1, -2, 3, 6, -8, 10}

	currentSum, winningSum := 0, 0

	for currentPos := 0; currentPos < len(myset); currentPos++ {
		currentValue := myset[currentPos]

		currentSum += currentValue
		fmt.Println("Current sum is now", currentSum)

		if currentSum > winningSum {
			fmt.Println("Current sum is now bigger than winning sum so promoting it")
			winningSum = currentSum
			fmt.Println("Winning sum is now", winningSum)
		}
		if currentSum < 0 {
			fmt.Println("Current sum is now less that 0 so splitting the set here and starting again")
			currentSum = 0
		}
	}

	fmt.Println("The maximum sum is", winningSum)
}
