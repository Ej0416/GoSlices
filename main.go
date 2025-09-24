package main

import (
	"fmt"
)

// func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
// 	message := [3]string{primary,secondary,tertiary}
// 	costPerMessage := [3]int{len(primary), len(secondary) + len(primary), len(tertiary) + len(secondary) + len(primary)}

// 	return  message, costPerMessage
// }

// const (
// 	planFree = "free"
// 	planPro  = "pro"
// )

// func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
// 	switch plan {
// 	case planPro:
// 		return messages[:], nil
// 	case planFree:
// 		return messages[:2], nil
// 	default:
// 		return nil, errors.New("unsupported plan")
// 	}
// }

// ---------------------------------------------- make slices

// func getMessageCosts(messages []string) []float64 {
// 	messageCost := make([]float64, len(messages))

// 	for i := range messages {
// 		cost := float64(len(messages[i])) * .01
// 		messageCost[i] = cost
// 	}

// 	return messageCost
// }

// -------------------------------------------------------------- variadic
// type TestType struct {
// 	Nums []int
// 	strs []string
// }

// func test(t TestType) {
// 	fmt.Println(t.Nums, t.strs)
// }

// func sum(nums ...int) int {
// 	sum := 0

// 	for i := range len(nums) {
// 		sum += nums[i]
// 	}

// 	return sum
// }

// type cost struct {
// 	day   int
// 	value float64
// }

// func getDayCosts(costs []cost, day int) []float64 {
// 	dayCost := make([]float64, 0)
// 	for i := range len(costs){
// 		if costs[i].day == day {
// 			dayCost = append(dayCost, costs[i].value)
// 		}
// 		continue
// 	}

// 	return dayCost
// }

// ------------------------------------------------------------- range
// func indexOfFirstBadWord(msg []string, badWords []string) int {
// 	for i, msgWords := range msg {

// 		for _, bw := range badWords {
// 			if msgWords == bw {
// 				return i
// 			}
// 		}

// 		// or

// 		// if slices.Contains(badWords, msgWords) {
// 		// 		return i
// 		// 	}
// 	}
// 	return -1
// }

func createMatrix(rows, cols int) [][]int {
	// make the rows
	matrix := make([][]int, rows)

	for i := range rows {
		// prepare the lenght of the rows so it wont get out of bounds
		matrix[i] = make([]int,cols)
		for j := range cols{
			matrix[i][j] = i * j
		}
	}

	return matrix
}


func main() {
	fmt.Println("app start")
}
