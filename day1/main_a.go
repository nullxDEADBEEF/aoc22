package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read input data
	inputData, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputData.Close()

	scanner := bufio.NewScanner(inputData)
	calories := 0
	largestCalories := 0

	for scanner.Scan() {
		calorie, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if calories > largestCalories {
				largestCalories = calories
			}
			calories = 0
		}
		calories += calorie
	}

	fmt.Println(largestCalories)

}
