package main

import (
	"fmt"
	"strings"
)

func main() {
	day12part1()
	day12part2()
}

func day12part1() {
	caves := make(map[string][]string)

	for _, line := range strings.Split(day12input, "\n") {
		parts := strings.Split(line, "-")

		found := false

		for _, c := range caves[parts[0]] {
			if parts[1] == c {
				found = true
				break
			}
		}

		if !found {
			caves[parts[0]] = append(caves[parts[0]], parts[1])
			caves[parts[1]] = append(caves[parts[1]], parts[0])
		}
	}

	count := walk(caves, "start", []string{}, func(c string, visited []string) bool {
		if c == strings.ToLower(c) {
			for _, v := range visited {
				if c == v {
					return true
				}
			}
		}

		return false
	})

	fmt.Println("Path count:", count)
}

func day12part2() {
	caves := make(map[string][]string)

	for _, line := range strings.Split(day12input, "\n") {
		parts := strings.Split(line, "-")

		found := false

		for _, c := range caves[parts[0]] {
			if parts[1] == c {
				found = true
				break
			}
		}

		if !found {
			caves[parts[0]] = append(caves[parts[0]], parts[1])
			caves[parts[1]] = append(caves[parts[1]], parts[0])
		}
	}

	count := walk(caves, "start", []string{}, func(c string, visited []string) bool {
		if c == "start" {
			return true
		}

		if c == strings.ToLower(c) {
			doublesUsed := false
			for _, v := range visited {
				counter := 0
				for _, p := range visited {
					if strings.ToLower(v) == p {
						counter++
					}
				}
				if counter >= 2 {
					doublesUsed = true
					break
				}
			}

			seen := 0
			for _, v := range visited {
				if c == v {
					seen++
				}
			}

			if (doublesUsed && seen > 0) || (!doublesUsed && seen > 1) {
				return true
			}
		}

		return false
	})

	fmt.Println("Path count:", count)
}

func walk(s map[string][]string, c string, visited []string, skipCave func(c string, visited []string) bool) int {
	count := 0

	visited = append(visited, c)

	if c == "end" {
		return 1
	}

	for _, m := range s[c] {
		skip := skipCave(m, visited)

		if !skip {
			count += walk(s, m, visited, skipCave)
		}
	}

	return count
}

var test12 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var test12b = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var day12input = `BC-gt
gt-zf
end-KH
end-BC
so-NL
so-ly
start-BC
NL-zf
end-LK
LK-so
ly-KH
NL-bt
gt-NL
start-zf
so-zf
ly-BC
BC-zf
zf-ly
ly-NL
ly-LK
IA-bt
bt-so
ui-KH
gt-start
KH-so`
