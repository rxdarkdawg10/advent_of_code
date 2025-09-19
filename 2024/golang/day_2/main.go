package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	total1 := part_one()
	fmt.Println("Part 1:", total1)

	total2 := part_two()
	fmt.Println("Part 2:", total2)
}

func read_file() [][]int64 {
	var (
		rows [][]int64
	)
	data, err := os.ReadFile("2024/data/2-act.txt")
	if err != nil {
		panic(err)
	}

	filestr := strings.Split(string(data), "\n")
	for _, row := range filestr {
		cols := strings.Split(row, " ")
		var r []int64
		for _, col := range cols {
			// fmt.Println(fmt.Sprintf("column: {%v}", col))
			n, _ := strconv.ParseInt(col, 10, 0)
			r = append(r, n)
		}
		rows = append(rows, r)
	}
	return rows
}

func part_one() int64 {
	var total int64
	data := read_file()

	var levels []int64
	for _, row := range data {
		safe := false
		levels = row

		if check_safe(levels) {
			safe = true
		}

		if safe {
			total += 1
		}
	}

	return total
}

func check_safe(levels []int64) bool {
	var increases, decreases bool

	increases = true
	decreases = true

	for idx := range levels {
		if idx+1 < len(levels) {
			diff := math.Abs(float64(levels[idx+1]) - float64(levels[idx]))

			if diff < 1 || diff > 3 {
				return false
			}

			if levels[idx] >= levels[idx+1] {
				increases = false
			}
			if levels[idx] <= levels[idx+1] {
				decreases = false
			}
		}
	}

	return increases || decreases
}

func part_two() int64 {
	var total int64
	data := read_file()

	var levels []int64
	for _, row := range data {
		safe := false
		levels = row

		if check_safe(levels) {
			safe = true
		}

		for idx := range levels {
			var levels_adj []int64
			for nidx := range levels {
				if idx != nidx {
					levels_adj = append(levels_adj, levels[nidx])
				}
			}

			if check_safe(levels_adj) {
				safe = true
			}
		}

		if safe {
			total += 1
		}
	}

	return total
}
