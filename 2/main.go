package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = `1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,19,5,23,2,9,23,27,1,5,27,31,1,5,31,35,1,35,13,39,1,39,9,43,1,5,43,47,1,47,6,51,1,51,13,55,1,55,9,59,1,59,13,63,2,63,13,67,1,67,10,71,1,71,6,75,2,10,75,79,2,10,79,83,1,5,83,87,2,6,87,91,1,91,6,95,1,95,13,99,2,99,13,103,1,103,9,107,1,10,107,111,2,111,13,115,1,10,115,119,1,10,119,123,2,13,123,127,2,6,127,131,1,13,131,135,1,135,2,139,1,139,6,0,99,2,0,14,0
`

func main() {
	inputStrArr := strings.Split(input, ",")
	inputIntArr := []int{}
	for i, _ := range inputStrArr{
		intVal, _ := strconv.Atoi(inputStrArr[i])
		inputIntArr = append(inputIntArr, intVal)
	}

	res1Arr := executeProgram(inputIntArr)
	fmt.Printf("Part 1: %d\n", res1Arr[0])
	goal := 19690720

	for i := 0; i < len(inputIntArr); i++ {
		for j := 0; j < len(inputIntArr); j++ {
			inputIntArr[1] = i
			inputIntArr[2] = j

			checkArr := executeProgram(inputIntArr)
			if checkArr[0] == goal {
				fmt.Printf("Part 2: %d", 100*i+j)
			}
		}
	}

}

func executeProgram(inputIntArr []int) []int {
	halting := false
	arrToUse := make([]int, len(inputIntArr))
	copy(arrToUse, inputIntArr)

	for x := 0; x < len(arrToUse)/4; x++ {
		var ints [4]int
		copy(ints[:4], arrToUse[x*4:x*4+4])
		arrToUse, halting = performOperation(ints, arrToUse)
		if halting {
			break
		}
	}
	return arrToUse
}

func performOperation(slice[4]int, input []int) (output []int, halting bool){
	operation := slice[0]
	pos1 := slice[1]
	pos2 := slice[2]
	outputPos := slice[3]
	if operation == 99 {
		return input, true
	}
	if operation == 1 {
		input[outputPos] = input[pos1]+input[pos2]
		return input, false
	}
	if operation == 2 {
		input[outputPos] = input[pos1]*input[pos2]
		return input, false
	}
	return input, true
}
