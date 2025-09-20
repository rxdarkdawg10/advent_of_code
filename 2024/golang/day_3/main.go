package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	total1 := part_one()
	fmt.Println("Part1:", total1)
	total2 := part_two()
	fmt.Println("Part2:", total2)
}

func part_one() int {
	var total int
	data := read_file(`mul\(\d{1,3}\,\d{1,3}\)`)

	for _, d := range data {
		regex, err := regexp.Compile(`\d{1,3}`)
		if err != nil {
			fmt.Println("ERROR", err)
		}

		vals := regex.FindAllString(d, -1)
		val1, _ := strconv.Atoi(vals[0])
		val2, _ := strconv.Atoi(vals[1])
		calc := val1 * val2

		total += calc
	}

	return total
}

func part_two() int {
	var total int
	data := read_file(`(mul\(\d{1,3}\,\d{1,3}\))|(don't\(\))|do\(\)`)
	runcalc := true
	for _, d := range data {
		regex, err := regexp.Compile(`\d{1,3}`)
		if err != nil {
			fmt.Println("ERROR", err)
		}
		if d == "don't()" {
			runcalc = false
		}
		if d == "do()" {
			runcalc = true
		}

		if runcalc {
			vals := regex.FindAllString(d, -1)
			if len(vals) > 0 {
				val1, _ := strconv.Atoi(vals[0])
				val2, _ := strconv.Atoi(vals[1])
				calc := val1 * val2

				total += calc
			}
		}
	}

	return total
}

func read_file(rgx string) []string {
	var (
		rows []string
	)
	data, err := os.ReadFile("2024/data/3-act.txt")
	if err != nil {
		panic(err)
	}

	regex, err := regexp.Compile(rgx)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	rows = regex.FindAllString(string(data), -1)

	return rows
}
