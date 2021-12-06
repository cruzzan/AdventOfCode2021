package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	day6part1()
	day6part2()
}

func day6part1()  {
	sf := strings.Split(startFishies, ",")
	fish := make([]int, len(sf))

	for key, f := range sf {
		var err error
		fish[key], err = strconv.Atoi(f)
		if err != nil {
			fmt.Println("Error on fish #", key, f, err)
		}
	}

	for i := 0; i < 80; i++ {
		for key, f := range fish {
			switch f {
			case 0:
				fish[key] = 6
				fish = append(fish, 8)
				break
			default:
				fish[key] = f - 1
			}
		}
	}

	fmt.Println("Fish count after 80 days:", len(fish))
}

func day6part2() {
	sf := strings.Split(startFishies, ",")
	fish := make([]int, 9)

	for _, f := range sf {
		var err error
		age, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("Error on fish", f, err)
		}

		fish[age]++
	}

	for i := 0; i < 256; i++ {
		tempFish := 0
		for j := 8; j >= 0; j-- {
			if j == 0 {
				fish[6] += fish[j]
				fish[8] = fish[j]
				fish[j] = tempFish
			} else if j == 8 {
				tempFish = fish[j]
				fish[j] = 0
			} else {
				temp := fish[j]
				fish[j] = tempFish
				tempFish = temp
			}
		}
	}

	fishCount := 0
	for i := 0; i < 9; i++ {
		fishCount += fish[i]
	}

	fmt.Println("Fish count after 256 days:", fishCount)
}

var startFishies = `2,1,1,4,4,1,3,4,2,4,2,1,1,4,3,5,1,1,5,1,1,5,4,5,4,1,5,1,3,1,4,2,3,2,1,2,5,5,2,3,1,2,3,3,1,4,3,1,1,1,1,5,2,1,1,1,5,3,3,2,1,4,1,1,1,3,1,1,5,5,1,4,4,4,4,5,1,5,1,1,5,5,2,2,5,4,1,5,4,1,4,1,1,1,1,5,3,2,4,1,1,1,4,4,1,2,1,1,5,2,1,1,1,4,4,4,4,3,3,1,1,5,1,5,2,1,4,1,2,4,4,4,4,2,2,2,4,4,4,2,1,5,5,2,1,1,1,4,4,1,4,2,3,3,3,3,3,5,4,1,5,1,4,5,5,1,1,1,4,1,2,4,4,1,2,3,3,3,3,5,1,4,2,5,5,2,1,1,1,1,3,3,1,1,2,3,2,5,4,2,1,1,2,2,2,1,3,1,5,4,1,1,5,3,3,2,2,3,1,1,1,1,2,4,2,2,5,1,2,4,2,1,1,3,2,5,5,3,1,3,3,1,4,1,1,5,5,1,5,4,1,1,1,1,2,3,3,1,2,3,1,5,1,3,1,1,3,1,1,1,1,1,1,5,1,1,5,5,2,1,1,5,2,4,5,5,1,1,5,1,5,5,1,1,3,3,1,1,3,1`
