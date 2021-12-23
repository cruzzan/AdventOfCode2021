package main

import (
	"fmt"
	"strings"
)

func main() {
	day14part1()
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
