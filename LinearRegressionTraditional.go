package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

var n int = 1000000
//var n int = 5

func mean(array []float64) float64 {
	var sum float64 = 0.0
	for _, v := range array {
		sum += float64(v)
	}

	return sum / float64(n)
}
func slopeCalc(xValue []float64, yValue []float64, xMean *float64, yMean *float64, slope *float64) {
	sumX, sumY := 0.0, 0.0

	for i := 0; i < len(xValue); i++ {
        sumX += (xValue[i] - *xMean) * (yValue[i] - *yMean)
        sumY += math.Pow(xValue[i] - *xMean, 2)
    }

	*slope += sumX / sumY
}

func linearRegression(xArray []float64, yArray []float64) (float64, float64) {
	var xMean, yMean, slope, intercept float64
	
	numChunks := 10
	chunkSize := int(math.Ceil(float64(len(xArray)) / float64(numChunks)))

	xMean, yMean = mean(xArray), mean(yArray)
	slope = 0.0

	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
        end := start + chunkSize
        if end > len(xArray) {
            end = len(xArray)
        }
		slopeCalc(xArray[start:end], yArray[start:end], &xMean, &yMean, &slope)
	}

	slope = slope / float64(numChunks)
	intercept = yMean - slope * xMean

	return slope, intercept
}

func createArrayValues(min, max float64) []float64 {
	temporalArray := make([]float64, n)
	
	for i := range temporalArray {
		temporalArray[i] = min + rand.Float64() * (max - min)
	}

	return temporalArray
}

func main() {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	xValue, yValue := createArrayValues(0.0, 1000.0), createArrayValues(0.0, 1000.0)
	//xValue, yValue := []float64{0,1,2,3,4}, []float64{0,1,2,3,4}

	slope, intercept := linearRegression(xValue, yValue)
	fmt.Println("Slope: ", slope)
	fmt.Println("Intercept: ", intercept)
	fmt.Println("Execution time: ", time.Since(start))
}