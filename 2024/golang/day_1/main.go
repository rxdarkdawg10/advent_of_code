package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	total := part_one()
	fmt.Println("Part 1 - Total", total)

	total2 := part_two()
	fmt.Println("Part 2 - Total", total2)

}

func read_file() ([]int64, []int64) {
	var (
		rhs []int64
		lhs []int64
	)
	data, err := os.ReadFile("2024/data/1-act.txt")
	if err != nil {
		panic(err)
	}

	filestr := strings.Split(string(data), "\n")
	for _, row := range filestr {
		cols := strings.Split(row, " ")
		for idx, col := range cols {

			if col != "" {
				if idx > 0 {
					// fmt.Println(fmt.Sprintf("column: {%v}", col))
					n, _ := strconv.ParseInt(col, 10, 0)
					rhs = append(rhs, n)
				} else {
					n, _ := strconv.ParseInt(col, 10, 0)
					lhs = append(lhs, n)
				}
			}
		}
	}
	return lhs, rhs
}

func part_one() int64 {

	lhs, rhs := read_file()

	slices.SortFunc(rhs, func(i, j int64) int {
		if i < j {
			return -1
		}
		return 1
	})
	slices.SortFunc(lhs, func(i, j int64) int {
		if i < j {
			return -1
		}
		return 1
	})
	var total int64 = 0

	for i := 0; i < len(lhs); i++ {
		diff := math.Abs(float64(lhs[i]) - float64(rhs[i]))
		total += int64(diff)
	}

	return total

}

func part_two() int64 {
	var (
		total int64
	)

	lhs, rhs := read_file()

	for _, n := range lhs {
		count := 0
		for _, r := range rhs {
			if n == r {
				count += 1
			}
		}

		total += int64(count) * n
	}

	return total
}
