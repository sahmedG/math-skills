package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/montanaflynn/stats"
)

func main() {
	//check if more than two arguments return with error
	if len(os.Args) == 2 {
		//reading the file name from os.args and opening it
		readFile, err := os.Open(os.Args[1])
		//case error opening the file throw error and exit the program
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//creating a scanner for each new line in the file case the file has multiple lines
		fileScanner := bufio.NewScanner(readFile)
		//splitting the data from the input file
		fileScanner.Split(bufio.ScanLines)
		// array to hold the converted strings to float64
		var numbers []float64
		// appending the values from scanner into float64 array
		for fileScanner.Scan() {
			number,_ := strconv.Atoi(fileScanner.Text())
			numbers = append(numbers, float64(number))
		}
		// closing the file after all the lines where scanned
		readFile.Close()
		//sorting the array
		sort.Float64s(numbers)
		//calculating median value
		median,_ := stats.Median(numbers)
		//caculating varience
		varience,_ := stats.Variance(numbers)
		//calculating standard deviation
		std_dev,_ := stats.StandardDeviation(numbers)
		//printing the output to the terminal using the required format
		fmt.Println("Average: ",math.Round(Average(numbers)))
		fmt.Println("Median: ",math.Round(median))
		fmt.Println("Variance:",int64(math.Round(varience)))
		fmt.Println("Standard Deviation:",math.Round(std_dev))
	} else {
		fmt.Println("Wrong nummbr of arguments")
	}
}

//function to calculate the average by adding all digits and dividing them by their count
func Average(array []float64) float64{
	n := len(array)
	sum := 0.0
	for i:=0;i<n;i++{
		sum += array[i]
	}
	return (sum) / float64(n)
}
