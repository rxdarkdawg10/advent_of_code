package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	first  int
	second int
}

type seq struct {
	sequence    []int
	brokenrules []int
	valid       bool
}

func main() {
	total, seqs, rules := part_one()
	fmt.Println("Part1:", total)
	total2 := part_two(seqs, rules)
	fmt.Println("Part2:", total2)
}

func part_one() (int, []seq, []rule) {
	var total int
	var rules []rule
	var seqs []seq

	data := read_file()

	printing_seq := false
	for _, row := range strings.Split(string(data), "\n") {

		if row == "" {
			printing_seq = true
		} else {
			if printing_seq {
				pages := strings.Split(row, ",")
				var temp []int
				for i := range pages {
					p, _ := strconv.Atoi(pages[i])
					temp = append(temp, p)
				}
				seqs = append(seqs, seq{sequence: temp, valid: true})
			} else {
				r := strings.Split(row, "|")
				r1, _ := strconv.Atoi(r[0])
				r2, _ := strconv.Atoi(r[1])
				rules = append(rules, rule{first: r1, second: r2})
			}
		}
	}

	for i := range seqs {
		for r := range rules {
			fIndex := 1000
			sIndex := 1000
			pages := seqs[i].sequence
			for p := range pages {
				if rules[r].first == pages[p] {
					fIndex = p
				}
				if rules[r].second == pages[p] {
					sIndex = p
				}
			}

			if fIndex != 1000 && sIndex != 1000 {
				if fIndex > sIndex {
					seqs[i].brokenrules = append(seqs[i].brokenrules, r)
					seqs[i].valid = false
				}
			}
		}
	}

	for i := range seqs {
		if seqs[i].valid {
			mIndex := len(seqs[i].sequence) / 2
			total += seqs[i].sequence[mIndex]
		}
	}

	return total, seqs, rules
}

func part_two(seqs []seq, rules []rule) int {
	var total int
	var newSeq []seq
	var fixedSeq []seq

	for i := range seqs {
		if !seqs[i].valid {
			newSeq = append(newSeq, seqs[i])
		}
	}
	// fmt.Println("ORIG:", newSeq)

	timing := 0
	for timing < 100 {
		var temp []seq
		for i := range newSeq {
			if !newSeq[i].valid {
				temp = append(temp, seq{sequence: newSeq[i].sequence, brokenrules: newSeq[i].brokenrules, valid: newSeq[i].valid})
			} else {
				fixedSeq = append(fixedSeq, newSeq[i])
			}
		}

		newSeq = swap(temp, rules)

		for i := range newSeq {
			for r := range rules {
				fIndex := 1000
				sIndex := 1000
				pages := newSeq[i].sequence
				for p := range pages {
					if rules[r].first == pages[p] {
						fIndex = p
					}
					if rules[r].second == pages[p] {
						sIndex = p
					}
				}

				if fIndex != 1000 && sIndex != 1000 {

					if fIndex > sIndex {
						newSeq[i].brokenrules = append(newSeq[i].brokenrules, r)
						newSeq[i].valid = false
					}
				}
			}
		}
		timing = timing + 1
	}

	// fmt.Println("FIXED:", fixedSeq)
	for i := range fixedSeq {
		if fixedSeq[i].valid {
			mIndex := len(fixedSeq[i].sequence) / 2
			total += fixedSeq[i].sequence[mIndex]
		}
	}

	return total
}

func swap(seqs []seq, rules []rule) []seq {
	var newSeq []seq

	for i := range seqs {
		// if !seqs[i].valid {
		rule := rules[seqs[i].brokenrules[0]]
		temp := seqs[i].sequence
		fIndex := 1000
		sIndex := 1000
		for s := range seqs[i].sequence {
			sq := seqs[i].sequence[s]
			if sq == rule.first {
				fIndex = s
			}

			if sq == rule.second {
				sIndex = s
			}
		}
		temp[fIndex] = rule.second
		temp[sIndex] = rule.first
		newSeq = append(newSeq, seq{sequence: temp, valid: true})
		// }
	}

	if len(newSeq) > 0 {
		return newSeq
	}

	return seqs
}

func read_file() []byte {
	data, err := os.ReadFile("2024/data/5-act.txt")
	if err != nil {
		panic(err)
	}

	return data
}
