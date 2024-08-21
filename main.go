package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Check if have some args to work with
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("No arguments were provided to the program.")
		return
	} else if len(arguments) > 2 {
		fmt.Println("Please only provide the quantity of floats to generate.")
		return
	}

	values := []float64{}

	//TODO: add some default value to the program, so this can run without arguments
	quantityToGenerate, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(arguments[1], "is not a valid value")
		return
	}

	for i := 0; i < quantityToGenerate; i++ {
		val := randomFloat(-10, 10)
		values = append(values, val)
	}

	sort.Float64s(values)
	fmt.Println("Number of values:", len(values))
	fmt.Println("Min:", values[0])
	fmt.Println("Max:", values[len(values)-1])

	var sum float64 //0.0
	for _, v := range values {
		sum += v
	}

	meanVal := sum / float64(len(values))
	fmt.Printf("Mean value: %.5f\n", meanVal)

	var squared float64
	for i := 0; i < len(values); i++ {
		squared += math.Pow((values[i] - meanVal), 2)
	}

	stdDeviation := math.Sqrt(squared / float64(len(values)))
	fmt.Printf("Standard deviation: %.5f\n", stdDeviation)

	normalized := normalize(stdDeviation, meanVal, values)
	fmt.Println("normalized:", normalized)
}

// https://en.wikipedia.org/wiki/Standard_score
func normalize(stdDeviation float64, mean float64, data []float64) []float64 {
	if stdDeviation == 0 {
		return data
	}

	normalized := make([]float64, len(data))
	for i, v := range data {
		// The number of zeros determines the level of accuracy for the result
		normalized[i] = math.Floor((v-mean)/stdDeviation*10_000) / 10_000
	}

	return normalized
}

func randomFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
