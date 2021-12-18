package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	day11part1()
	day11part2()
}

func day11part1() {
	dumbos := make([]*dumbo, 0)
	lines := strings.Split(day11input, "\n")

	for _, l := range lines {
		for _, c := range l {
			e, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println(err)
			}

			dumbos = append(dumbos, &dumbo{flashed: false, energy: e})
		}
	}

	flashes := 0
	for i := 1; i <= 100; i++ {
		for j := 0; j < len(dumbos); j++ {
			dumbos[j].energy++
			dumbos[j].flashed = false
		}

		for j := 0; j < len(dumbos); j++ {
			if dumbos[j].energy > 9 && !dumbos[j].flashed {
				flash(j, dumbos)
			}
		}

		for j := 0; j < len(dumbos); j++ {
			if dumbos[j].flashed {
				flashes++
				dumbos[j].energy = 0
			}
		}
	}

	fmt.Println("Flash count:", flashes)
}

func day11part2() {
	dumbos := make([]*dumbo, 0)
	lines := strings.Split(day11input, "\n")

	for _, l := range lines {
		for _, c := range l {
			e, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println(err)
			}

			dumbos = append(dumbos, &dumbo{flashed: false, energy: e})
		}
	}

	step := 1
	for {
		flashes := 0
		for j := 0; j < len(dumbos); j++ {
			dumbos[j].energy++
			dumbos[j].flashed = false
		}

		for j := 0; j < len(dumbos); j++ {
			if dumbos[j].energy > 9 && !dumbos[j].flashed {
				flash(j, dumbos)
			}
		}

		for j := 0; j < len(dumbos); j++ {
			if dumbos[j].flashed {
				flashes++
				dumbos[j].energy = 0
			}
		}

		if flashes == 100 {
			break
		}

		step++
	}

	fmt.Println("First synchronization:", step)
}

func flash(i int, dumbos []*dumbo)  {
	dumbos[i].flashed = true

	for _, ix := range []int{i -1, i +1, i -11, i -10, i -9, i +9, i +10, i +11} {
		o := i + 1
		ox := ix + 1
		if (ox % 10 == o % 10) || (o % 10 > 1 && i % 10 < 9) || ((o % 10 == 1) && (ox % 10 == 2)) || ((o % 10 == 0) && (ox % 10 == 9)) {
			if ix >= 0 && ix < len(dumbos) {
				dumbos[ix].energy++

				if dumbos[ix].energy > 9 && !dumbos[ix].flashed {
					flash(ix, dumbos)
				}
			}
		}
	}
}

type dumbo struct {
	flashed bool
	energy int
}

var test11 = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var day11input = `5665114554
4882665427
6185582113
7762852744
7255621841
8842753123
8225372176
7212865827
7758751157
1828544563`
