package day1

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/gleasonw/aoc-2024/utils"
)

func Solve() int {
	arr1, arr2 := readInput()
	slices.Sort(arr1)
	slices.Sort(arr2)
	sumDiff := 0
	for i, val := range arr1 {
		arr2Val := arr2[i]
		diff := math.Abs(float64(val - arr2Val))
		sumDiff += int(diff)
	}
	return sumDiff

}

func SolvePartTwo() int {
	arr1, arr2 := readInput()
	slices.Sort(arr1)
	slices.Sort(arr2)
	arr2Counts := make(map[int64]int)
	for _, val := range arr2 {
		if _, ok := arr2Counts[val]; !ok {
			arr2Counts[val] = 1
		} else {
			arr2Counts[val]++
		}
	}
	similaritySum := 0
	for _, val := range arr1 {
		if count, ok := arr2Counts[val]; !ok {
			continue
		} else {
			similaritySum += count * int(val)
		}
	}

	return similaritySum

}

func readInput() ([]int64, []int64) {
	array1 := make([]int64, 100)
	array2 := make([]int64, 100)
	file, error := os.Open("day1/input.txt")
	utils.Check(error)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		toAdd := strings.Split(line, "   ")
		first, err := strconv.ParseInt(toAdd[0], 0, 64)
		utils.Check(err)
		second, err := strconv.ParseInt(toAdd[1], 0, 64)
		utils.Check(err)
		array1 = append(array1, first)
		array2 = append(array2, second)

	}
	if len(array1) != len(array2) {
		panic("unexpected array length input mismatch")
	}
	return array1, array2
}
