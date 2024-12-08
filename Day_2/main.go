package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getData(filename string) [][]int {
	var data [][]int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return data
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		tempData := strings.Fields(line)
		row := make([]int, 0, len(tempData))

		for _, str := range tempData {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			row = append(row, num)
		}

		if len(row) > 0 {
			data = append(data, row)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return data
}

func checkReport(report []int) bool {

	delta := 0.0
	differentDelta := 0
	isSorted := true

	for i := 1; i < len(report); i++ {
		delta = math.Abs(float64(report[i-1]) - float64(report[i]))
		if delta < 1 || delta > 3 {
			differentDelta++
		}
		if report[i-1] < report[i] {
			isSorted = false
		}

	}

	return (isSorted || slices.IsSorted(report)) && differentDelta == 0

}

func main() {
	input := "input.txt"
	reports := getData(input)
	sumOfSafeReports := 0
	for _, report := range reports {
		if checkReport(report) {
			sumOfSafeReports++
		}
	}
	fmt.Printf("Part 1: %d\n", sumOfSafeReports)
}
