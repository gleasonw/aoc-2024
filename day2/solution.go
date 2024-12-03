package day2

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gleasonw/aoc-2024/utils"
)

func Solve() int {
	numSafe := 0
	for _, val := range readInput() {
		if val.isSafeWithOneDropAllowed() {
			numSafe++
		}
	}
	return numSafe
}

func readInput() []Report {
	file, err := os.Open("day2/input.txt")
	utils.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var reports []Report
	for scanner.Scan() {
		line := scanner.Text()
		levelsRaw := strings.Split(line, " ")
		levelsParsed := make([]int, len(levelsRaw))
		for i, val := range levelsRaw {
			num, err := strconv.ParseInt(val, 10, 64)
			utils.Check(err)
			levelsParsed[i] = int(num)
		}
		reports = append(reports, Report{levels: levelsParsed})
	}
	return reports
}

type Report struct {
	levels []int
}

func (r Report) isSafeWithOneDropAllowed() bool {
	if r.isSafe() {
		return true
	}
	for i := range r.levels {
		sliceWithoutVal := make([]int, 0, len(r.levels)-1)
		sliceWithoutVal = append(sliceWithoutVal, r.levels[:i]...)
		sliceWithoutVal = append(sliceWithoutVal, r.levels[i+1:]...)
		newReport := Report{levels: sliceWithoutVal}
		if newReport.isSafe() {
			log.Println("safe with drop", r, sliceWithoutVal, r.levels[i])
			return true
		}
	}
	return false
}

func (r Report) isSafe() bool {
	var previousDiff int
	for i, val := range r.levels {
		if i == len(r.levels)-1 {
			continue
		}
		next := r.levels[i+1]
		diff := val - next
		absDiff := math.Abs(float64(diff))
		if absDiff == 0 || absDiff > 3 {
			return false
		}
		if diff < 0 {
			if previousDiff > 0 {
				return false
			}
			previousDiff = diff
		}
		if diff > 0 {
			if previousDiff < 0 {
				return false
			}
			previousDiff = diff
		}

	}
	return true
}
