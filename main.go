package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var (
		nValues       = 0
		initFlag      = 0
		min, max, sum float64
	)

	// Check if have some args to work with
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more aguments")
		return
	}

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Println(arguments[i], "is not aa valid value")
			continue
		}

		nValues++
		sum += n

		if initFlag == 0 {
			min = n
			max = n
			initFlag = 1
			continue
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Number of values:", nValues)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	if nValues == 0 {
		return
	}

	meanValue := sum / float64(nValues)
	fmt.Printf("Mean value: %.5f\n", meanValue)

	// Standar deviation
	var square float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		square += math.Pow((n - meanValue), 2)
	}

	stdDeviation := math.Sqrt(square / float64(nValues))
	fmt.Printf("Standard deviation: %.5f\n", stdDeviation)
}
