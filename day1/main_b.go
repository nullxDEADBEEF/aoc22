package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	calorieScores := []int{}

	for scanner.Scan() {
		calorie, err := strconv.Atoi(scanner.Text())
		if err != nil {
			calorieScores = append(calorieScores, calories)
			calories = 0
		}
		calories += calorie
	}

	sort.Ints(calorieScores)
	topThree := 0
	for _, score := range calorieScores[len(calorieScores)-3:] {
		topThree += score
	}
	fmt.Println(topThree)

}
