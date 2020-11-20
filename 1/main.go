package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputsArr := strings.Split(inputsStr, "\n")
	baseTotal := 0
	fuelTotal := 0
	for _, input := range inputsArr{
		inputFloat, err := strconv.ParseFloat(input, 32)
		if err != nil {
			log.Fatal(err)
		}
		fuelForMass := calcFuelForMass(inputFloat)
		baseTotal += fuelForMass
		_, fuelToAdd := calcFuelForFuel(float64(fuelForMass), float64(fuelForMass))
		fuelTotal += int(fuelToAdd)
	}

	fmt.Println(fmt.Sprintf("Part 1 Total: %d", baseTotal))
	fmt.Println(fmt.Sprintf("Part 2 Total: %d", fuelTotal))
}

func calcFuelForMass(mass float64) int {
	return int(math.Floor(mass/3) - 2)
}

func calcFuelForFuel(fuel float64, runningTotal float64) (float64, float64){
	// If less than 7 then we know it's zero, no calculation needed
	if fuel < 7 {
		return 0.0, runningTotal
	}
	newFuel := float64(calcFuelForMass(fuel))
	runningTotal+=newFuel
	return calcFuelForFuel(newFuel, runningTotal)
}

const inputsStr = `146561
98430
131957
81605
70644
55060
93217
107158
110769
94650
141070
72381
100736
105705
99003
94057
110662
74429
55509
63492
102007
72627
95183
112072
122313
116884
125451
106093
140678
121751
149018
58459
138306
149688
82927
72676
95010
88439
51807
103175
107633
126439
128879
112054
52873
114493
77365
76768
60838
89692
66217
96060
100338
139063
126869
106490
128967
116312
56822
52422
124579
117120
106245
105255
66975
115340
145764
149427
64228
64237
67887
103345
134901
50226
126991
122314
140818
129687
149792
101148
73411
87078
121272
108804
96063
81155
62058
112684
134263
128454
99455
91689
141448
143892
103257
64352
90769
78307
111855
130153`
