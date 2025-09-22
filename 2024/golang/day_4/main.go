package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type direction int

const (
	CENTER direction = iota
	UP
	DOWN
	LEFT
	RIGHT
	UPLEFT
	UPRIGHT
	DOWNLEFT
	DOWNRIGHT
)

type POS struct {
	x int
	y int
}

type searchdata struct {
	chars []searchChar
	keep  bool
	dirs  []d
	val   float32
}

type searchChar struct {
	char string
	pos  POS
	dir  direction
	keep bool
}

type d struct {
	dir   string
	chars string
}

func main() {
	total := part_one()
	fmt.Println("Part1:", total)
	total2 := part_two()
	fmt.Println("Part2:", total2)
}

func part_one() int {
	var total int
	var search []searchdata
	var matrix [][]rune
	data := read_file()

	for idx, row := range strings.Split(string(data), "\n") {
		matrix = append(matrix, make([]rune, 0))
		for _, r := range row {
			matrix[idx] = append(matrix[idx], r)
		}

	}

	for ridx, row := range matrix {
		for cidx, char := range row {
			// schar := searchChar{char: string(char), pos: POS{x: cidx, y: ridx}, keep: false}
			if char == 'X' {
				chars, keep := search_around(matrix, string(char), POS{x: cidx, y: ridx}, CENTER)

				if keep {
					search = append(search, searchdata{keep: true, chars: chars})
				}
			}
		}
	}

	for i, row := range search {
		// fmt.Println("SEARCH:", row)
		for mx1, m := range row.chars {
			if m.char == "M" {

				chars, keep := search_around(matrix, m.char, m.pos, m.dir)
				// fmt.Println("CHARS:", chars, "KEEP:", keep)
				search[i].chars[mx1].keep = keep
				if keep {
					for _, c := range chars {
						found := false
						for _, c2 := range search[i].chars {
							if c.char == c2.char && c.pos == c2.pos {
								found = true
							}
						}

						if !found {
							search[i].chars = append(search[i].chars, c)
						}
					}
				}
			}
		}
	}

	for i, row := range search {
		for ax1, a := range row.chars {
			if a.char == "A" {
				chars, keep := search_around(matrix, a.char, a.pos, a.dir)
				// fmt.Println("CHARS:", chars, "KEEP:", keep)
				search[i].chars[ax1].keep = keep
				if keep {
					for _, c := range chars {
						found := false
						for _, c2 := range search[i].chars {
							if c.char == c2.char && c.pos == c2.pos {
								found = true
							}
						}

						if !found {
							search[i].chars = append(search[i].chars, c)
						}
					}
				}
			}
		}
	}

	for i, row := range search {
		for _, char := range row.chars {
			found := false
			if char.keep && char.dir == UP {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "UP" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "UP",
						chars: "X" + char.char,
					})
				}
			}
			if char.keep && char.dir == DOWN {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "DOWN" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "DOWN",
						chars: "X" + char.char,
					})
				}
			}
			if char.keep && char.dir == LEFT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "LEFT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "LEFT",
						chars: "X" + char.char,
					})
				}
			}
			if char.keep && char.dir == RIGHT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "RIGHT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "RIGHT",
						chars: "X" + char.char,
					})
				}
			}
			if char.keep && char.dir == UPLEFT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "UPLEFT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "UPLEFT",
						chars: "X" + char.char,
					})
				}
			}
			if char.keep && char.dir == UPRIGHT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "UPRIGHT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "UPRIGHT",
						chars: "X" + char.char,
					})
				}
			}
			if char.keep && char.dir == DOWNRIGHT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "DOWNRIGHT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "DOWNRIGHT",
						chars: "X" + char.char,
					})
				}
			}
			if char.keep && char.dir == DOWNLEFT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "DOWNLEFT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "DOWNLEFT",
						chars: "X" + char.char,
					})
				}
			}
		}
	}

	for _, row := range search {
		for _, dir := range row.dirs {
			if dir.chars == "XMAS" {
				total += 1
			}
		}
		// fmt.Println("SEARCH:", row)
	}

	// fmt.Println("TEMP", temp)
	// fmt.Println(string(data))

	return total
}

func part_two() int {
	var total int
	var search []searchdata
	var matrix [][]rune
	data := read_file()

	for idx, row := range strings.Split(string(data), "\n") {
		matrix = append(matrix, make([]rune, 0))
		for _, r := range row {
			matrix[idx] = append(matrix[idx], r)
		}

	}

	for ridx, row := range matrix {
		for cidx, char := range row {
			// schar := searchChar{char: string(char), pos: POS{x: cidx, y: ridx}, keep: false}
			if char == 'A' {
				chars, keep := search_around2(matrix, string(char), POS{x: cidx, y: ridx}, CENTER)

				if keep {
					search = append(search, searchdata{keep: true, chars: chars})
				}
			}
		}
	}

	// for i, row := range search {
	// 	// fmt.Println("SEARCH:", row)
	// 	for mx1, m := range row.chars {
	// 		if m.char == "M" {

	// 			chars, keep := search_around(matrix, m.char, m.pos, m.dir)
	// 			// fmt.Println("CHARS:", chars, "KEEP:", keep)
	// 			search[i].chars[mx1].keep = keep
	// 			if keep {
	// 				for _, c := range chars {
	// 					found := false
	// 					for _, c2 := range search[i].chars {
	// 						if c.char == c2.char && c.pos == c2.pos {
	// 							found = true
	// 						}
	// 					}

	// 					if !found {
	// 						search[i].chars = append(search[i].chars, c)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// for i, row := range search {
	// 	for ax1, a := range row.chars {
	// 		if a.char == "A" {
	// 			chars, keep := search_around(matrix, a.char, a.pos, a.dir)
	// 			// fmt.Println("CHARS:", chars, "KEEP:", keep)
	// 			search[i].chars[ax1].keep = keep
	// 			if keep {
	// 				for _, c := range chars {
	// 					found := false
	// 					for _, c2 := range search[i].chars {
	// 						if c.char == c2.char && c.pos == c2.pos {
	// 							found = true
	// 						}
	// 					}

	// 					if !found {
	// 						search[i].chars = append(search[i].chars, c)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	for i, row := range search {
		for _, char := range row.chars {
			found := false

			if char.keep && char.dir == UPLEFT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "UPLEFT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "UPLEFT",
						chars: "A" + char.char,
					})
				}
			}
			if char.keep && char.dir == UPRIGHT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "UPRIGHT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "UPRIGHT",
						chars: "A" + char.char,
					})
				}
			}
			if char.keep && char.dir == DOWNRIGHT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "DOWNRIGHT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "DOWNRIGHT",
						chars: "A" + char.char,
					})
				}
			}
			if char.keep && char.dir == DOWNLEFT {
				for d := range search[i].dirs {
					if search[i].dirs[d].dir == "DOWNLEFT" {
						found = true
						search[i].dirs[d].chars += char.char
					}
				}
				if !found {
					search[i].dirs = append(search[i].dirs, d{
						dir:   "DOWNLEFT",
						chars: "A" + char.char,
					})
				}
			}
		}
	}

	for _, row := range search {
		total_count := 0.0
		for _, dir := range row.dirs {

			if dir.chars == "AS" && dir.dir == "DOWNLEFT" {
				for _, d := range row.dirs {
					if d.chars == "AM" && d.dir == "UPRIGHT" {
						total_count += 0.5
					}
				}
			}

			if dir.chars == "AM" && dir.dir == "DOWNLEFT" {
				for _, d := range row.dirs {
					if d.chars == "AS" && d.dir == "UPRIGHT" {
						total_count += 0.5
					}
				}
			}

			if dir.chars == "AS" && dir.dir == "DOWNRIGHT" {
				for _, d := range row.dirs {
					if d.chars == "AM" && d.dir == "UPLEFT" {
						total_count += 0.5
					}
				}
			}

			if dir.chars == "AM" && dir.dir == "DOWNRIGHT" {
				for _, d := range row.dirs {
					if d.chars == "AS" && d.dir == "UPLEFT" {
						total_count += 0.5
					}
				}
			}
		}

		// fmt.Println("Value:", total_count, search[i])
		if total_count == 1 {
			total += 1
			// fmt.Println("SEARCH:", row)

		}
	}

	// for _, row := range search {
	// 	fmt.Println("SEARCH:", row)
	// }

	// fmt.Println("TEMP", temp)
	// fmt.Println(string(data))

	return total
}

func search_around(data [][]rune, char string, pos POS, dir direction) ([]searchChar, bool) {
	var up, down, left, right, upright, upleft, downright, downleft string
	var nextchar []string
	var searches []searchChar
	if dir == CENTER {
		searches = append(searches, searchChar{char: char, pos: pos, dir: CENTER, keep: false})
	}
	switch char {
	case "X":
		nextchar = []string{"M"}
	case "M":
		nextchar = []string{"A"}
	case "A":
		nextchar = []string{"S", "M"}
	}

	if dir == CENTER || dir == UP {
		if pos.y != 0 {
			up = string(data[pos.y-1][pos.x])

			if !slices.Contains(nextchar, up) {
				up = "0"
			} else {
				searches = append(searches, searchChar{char: up, keep: true, pos: POS{x: pos.x, y: pos.y - 1}, dir: UP})
			}
		} else {
			up = "0"
		}
	}

	if dir == CENTER || dir == DOWN {
		if pos.y != len(data)-1 {
			down = string(data[pos.y+1][pos.x])

			if !slices.Contains(nextchar, down) {
				down = "0"
			} else {
				searches = append(searches, searchChar{char: down, keep: true, pos: POS{x: pos.x, y: pos.y + 1}, dir: DOWN})
			}
		} else {
			down = "0"
		}
	}

	if dir == CENTER || dir == LEFT {
		if pos.x != 0 {
			left = string(data[pos.y][pos.x-1])

			if !slices.Contains(nextchar, left) {
				left = "0"
			} else {
				searches = append(searches, searchChar{char: left, keep: true, pos: POS{x: pos.x - 1, y: pos.y}, dir: LEFT})
			}

		} else {
			left = "0"
		}
	}

	if dir == CENTER || dir == RIGHT {
		if pos.x != len(data[pos.y])-1 {
			right = string(data[pos.y][pos.x+1])
			if !slices.Contains(nextchar, right) {
				right = "0"
			} else {
				searches = append(searches, searchChar{char: right, keep: true, pos: POS{x: pos.x + 1, y: pos.y}, dir: RIGHT})

			}
		} else {
			right = "0"
		}
	}

	if dir == CENTER || dir == DOWNLEFT {
		if pos.x == 0 {
			downleft = "0"
		} else {
			if pos.y == len(data)-1 {
				downleft = "0"
			} else {
				downleft = string(data[pos.y+1][pos.x-1])
				if !slices.Contains(nextchar, downleft) {
					downleft = "0"
				} else {
					searches = append(searches, searchChar{char: downleft, keep: true, pos: POS{x: pos.x - 1, y: pos.y + 1}, dir: DOWNLEFT})

				}
			}
		}
	}

	if dir == CENTER || dir == DOWNRIGHT {
		if pos.x == len(data[pos.y])-1 {
			downright = "0"
		} else {
			if pos.y == len(data)-1 {
				downright = "0"
			} else {
				downright = string(data[pos.y+1][pos.x+1])
				if !slices.Contains(nextchar, downright) {
					downright = "0"
				} else {
					searches = append(searches, searchChar{char: downright, keep: true, pos: POS{x: pos.x + 1, y: pos.y + 1}, dir: DOWNRIGHT})

				}
			}
		}
	}

	if dir == CENTER || dir == UPLEFT {
		if pos.y == 0 {
			upleft = "0"
		} else {
			if pos.x == 0 {
				upleft = "0"
			} else {
				upleft = string(data[pos.y-1][pos.x-1])
				if !slices.Contains(nextchar, upleft) {
					upleft = "0"
				} else {
					searches = append(searches, searchChar{char: upleft, keep: true, pos: POS{x: pos.x - 1, y: pos.y - 1}, dir: UPLEFT})

				}
			}
		}
	}

	if dir == CENTER || dir == UPRIGHT {
		if pos.y == 0 {
			upright = "0"
		} else {
			if pos.x == len(data[pos.y])-1 {
				upright = "0"
			} else {
				upright = string(data[pos.y-1][pos.x+1])
				if !slices.Contains(nextchar, upright) {
					upright = "0"
				} else {
					searches = append(searches, searchChar{char: upright, keep: true, pos: POS{x: pos.x + 1, y: pos.y - 1}, dir: UPRIGHT})

				}
			}

		}
	}

	// fmt.Println("NEXT CHAR:", string(nextchar))
	// fmt.Println("Char:", string(char), "NEXTCHAR:", nextchar, "UP:", string(up), "DOWN:", string(down), "LEFT:", string(left), "RIGHT:", string(right), "UPLEFT:", string(upleft), "UPRIGHT:", string(upright), "DOWNLEFT:", string(downleft), "DOWNRIGHT:", string(downright))
	var keep bool
	if len(searches) > 0 {
		keep = true
		searches[0].keep = true
	}
	return searches, keep
}

func search_around2(data [][]rune, char string, pos POS, dir direction) ([]searchChar, bool) {
	var upright, upleft, downright, downleft string
	var nextchar []string
	var searches []searchChar
	if dir == CENTER {
		searches = append(searches, searchChar{char: char, pos: pos, dir: CENTER, keep: false})
	}
	switch char {
	case "A":
		nextchar = []string{"S", "M"}
	}

	if dir == CENTER || dir == DOWNLEFT {
		if pos.x == 0 {
			downleft = "0"
		} else {
			if pos.y == len(data)-1 {
				downleft = "0"
			} else {
				downleft = string(data[pos.y+1][pos.x-1])
				if !slices.Contains(nextchar, downleft) {
					downleft = "0"
				} else {
					searches = append(searches, searchChar{char: downleft, keep: true, pos: POS{x: pos.x - 1, y: pos.y + 1}, dir: DOWNLEFT})

				}
			}
		}
	}

	if dir == CENTER || dir == DOWNRIGHT {
		if pos.x == len(data[pos.y])-1 {
			downright = "0"
		} else {
			if pos.y == len(data)-1 {
				downright = "0"
			} else {
				downright = string(data[pos.y+1][pos.x+1])
				if !slices.Contains(nextchar, downright) {
					downright = "0"
				} else {
					searches = append(searches, searchChar{char: downright, keep: true, pos: POS{x: pos.x + 1, y: pos.y + 1}, dir: DOWNRIGHT})

				}
			}
		}
	}

	if dir == CENTER || dir == UPLEFT {
		if pos.y == 0 {
			upleft = "0"
		} else {
			if pos.x == 0 {
				upleft = "0"
			} else {
				upleft = string(data[pos.y-1][pos.x-1])
				if !slices.Contains(nextchar, upleft) {
					upleft = "0"
				} else {
					searches = append(searches, searchChar{char: upleft, keep: true, pos: POS{x: pos.x - 1, y: pos.y - 1}, dir: UPLEFT})

				}
			}
		}
	}

	if dir == CENTER || dir == UPRIGHT {
		if pos.y == 0 {
			upright = "0"
		} else {
			if pos.x == len(data[pos.y])-1 {
				upright = "0"
			} else {
				upright = string(data[pos.y-1][pos.x+1])
				if !slices.Contains(nextchar, upright) {
					upright = "0"
				} else {
					searches = append(searches, searchChar{char: upright, keep: true, pos: POS{x: pos.x + 1, y: pos.y - 1}, dir: UPRIGHT})

				}
			}

		}
	}

	// fmt.Println("NEXT CHAR:", string(nextchar))
	// fmt.Println("Char:", string(char), "NEXTCHAR:", nextchar, "UP:", string(up), "DOWN:", string(down), "LEFT:", string(left), "RIGHT:", string(right), "UPLEFT:", string(upleft), "UPRIGHT:", string(upright), "DOWNLEFT:", string(downleft), "DOWNRIGHT:", string(downright))
	var keep bool
	if len(searches) > 0 {
		keep = true
		searches[0].keep = true
	}
	return searches, keep
}

func read_file() []byte {
	data, err := os.ReadFile("2024/data/4-act.txt")
	if err != nil {
		panic(err)
	}

	return data
}
