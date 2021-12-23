package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	day14part1()
	day14part2()
}

func day14part1() {
	ir := make(map[string]string)
	lines := strings.Split(day14insertionrules, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		ir[parts[0]] = parts[1]
	}

	template := day14template
	for i := 0; i < 10; i++ {
		newTemplate := template

		cur := 1
		for j := 0; j < len(template)-1; j++ {
			pair := fmt.Sprintf("%s%s", string(template[j]), string(template[j+1]))
			newTemplate = fmt.Sprintf("%s%s%s", newTemplate[:cur], ir[pair], newTemplate[cur:])
			cur += 2
		}

		template = newTemplate
	}

	polymerCount := make(map[string]int)
	for i := 0; i < len(template); i++ {
		polymerCount[string(template[i])]++
	}

	fmt.Println("Polymer counts", polymerCount)
}

func day14part2() {
	ir := make(map[string]string)
	lines := strings.Split(day14insertionrules, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		ir[parts[0]] = parts[1]
	}

	pairs := make(map[string]int64)
	for i := 0; i < len(day14template)-1; i++ {
		pair := fmt.Sprintf("%s%s", string(day14template[i]), string(day14template[i+1]))
		pairs[pair]++
	}

	for i := 0; i < 40; i++ {
		tempPairs := make(map[string]int64)
		for key, count := range pairs {
			if count > 0 {
				tempPairs[key] = count
			}
		}
		for key, count := range pairs {
			first := fmt.Sprintf("%s%s", string(key[0]), ir[key])
			second := fmt.Sprintf("%s%s", ir[key], string(key[1]))
			tempPairs[key] -= count
			tempPairs[first] += count
			tempPairs[second] += count
		}
		pairs = tempPairs
	}

	polymerCount := make(map[string]int64)
	for pair, count := range pairs {
		for _, char := range pair {
			polymerCount[string(char)] += count
		}
	}

	fmt.Println("Polymer counts")
	for p, c := range polymerCount {
		adjusted := int(math.Ceil(float64(c) / float64(2)))
		fmt.Println(p, adjusted)
	}
}

var test14template = `NNCB`

var test14ir = `CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

var day14template = `FSKBVOSKPCPPHVOPVFPC`

var day14insertionrules = `BV -> O
OS -> P
KP -> P
VK -> S
FS -> C
OK -> P
KC -> S
HV -> F
HC -> K
PF -> N
NK -> F
SC -> V
CO -> K
PO -> F
FB -> P
CN -> K
KF -> N
NH -> S
SF -> P
HP -> P
NP -> F
OV -> O
OP -> P
HH -> C
FP -> P
CS -> O
SK -> O
NS -> F
SN -> S
SP -> H
BH -> B
NO -> O
CB -> N
FO -> N
NC -> C
VF -> N
CK -> C
PC -> H
BP -> B
NF -> O
BB -> C
VN -> K
OH -> K
CH -> F
VB -> N
HO -> P
FH -> K
PK -> H
CC -> B
VH -> B
BF -> N
KS -> V
PV -> B
CP -> N
PB -> S
VP -> V
BO -> B
HS -> H
BS -> F
ON -> B
HB -> K
KH -> B
PP -> H
BN -> C
BC -> F
KV -> K
VO -> P
SO -> V
OF -> O
BK -> S
PH -> V
SV -> F
CV -> H
OB -> N
SS -> H
VV -> B
OO -> V
CF -> H
KB -> F
NV -> B
FV -> V
HK -> P
VS -> P
FF -> P
HN -> N
FN -> F
OC -> K
SH -> V
KO -> C
HF -> B
PN -> N
SB -> F
VC -> B
FK -> S
KK -> N
FC -> F
NN -> P
NB -> V
PS -> S
KN -> S`
