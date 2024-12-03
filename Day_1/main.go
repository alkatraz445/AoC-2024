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
		sum += absInt(firstColumn[i] - secondColumn[i])
	}
	return sum
}

func multiplyLists(firstColumn, secondColumn []int) int {
	sum := 0
	count := make(map[int]int)
	for _, firstVal := range firstColumn {
		for _, secondVal := range secondColumn {
			if firstVal == secondVal {
				count[firstVal] = count[firstVal] + 1
			}
		}
	}
	for k, v := range count {
		sum += k * v
	}
	return sum
}

func main() {
	array := getData("input.txt")
	a, b := splitArray(array)
	fmt.Printf("Part 1: %d\n", sumLists(a, b))
	fmt.Printf("Part 2: %d\n", multiplyLists(a, b))

}
