package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* Sortowanie tablicy i używanie wartości absolutnej
//
// Problem z otrzymanym plikiem. Może czytanie od razu do int?
// Jak narazie program czyta tylko pierwszą cyfrę w liczbie.
//
*/

// Helper functions
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Main code

func getData(filename string) []int {
	var twoArraysString []string
	var twoArraysTemp []string
	var twoArrays []int
	// data, err := os.ReadFile(filename)
	// check(err)

	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		twoArraysString = append(twoArraysString, fileScanner.Text())
	}

	readFile.Close()

	for _, val := range twoArraysString {
		temp := strings.Split(val, "   ")
		twoArraysTemp = append(twoArraysTemp, temp...)
	}

	for _, val := range twoArraysTemp {
		newVal, err := strconv.Atoi(val)
		check(err)
		twoArrays = append(twoArrays, newVal)
	}

	fmt.Println(twoArrays)

	return twoArrays
}

func splitArray(array []int) ([]int, []int) {
	var firstColumn []int
	var secondColumn []int
	for i := range array {
		if i%2 == 0 {
			firstColumn = append(firstColumn, array[i])
		} else {
			secondColumn = append(secondColumn, array[i])
		}

	}

	sort.Slice(firstColumn, func(i, j int) bool {
		return firstColumn[i] < firstColumn[j]
	})

	sort.Slice(secondColumn, func(i, j int) bool {
		return secondColumn[i] < secondColumn[j]
	})

	return firstColumn, secondColumn
}

func sumLists(firstColumn, secondColumn []int) int {
	var sum int = 0

	for i := range firstColumn {
		fmt.Println(firstColumn[i])
		sum += absInt(firstColumn[i] - secondColumn[i])
		fmt.Println(sum)
	}
	return sum
}

func main() {
	array := getData("input.txt")
	a, b := splitArray(array)
	result := sumLists(a, b)
	println(result)

}
