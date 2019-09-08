package main

import (
	"fmt"
)

func main() {
	fmt.Println("DCP-27")

	pba := [][2]int{} // holds parenthesis pairs
	cba := [][2]int{} // holds curly brace pairs
	sba := [][2]int{} // holds square bracket pairs

	tStr := "((){[]}()[{((()))}])"

	for x := 0; x < len(tStr); x++ {
		tmp := 0
		if tStr[x] == '(' {
			a := x
			tmp++
			for y := x + 1; y < len(tStr); y++ {
				if tStr[y] == '(' {
					tmp++
				} else if tStr[y] == ')' {
					tmp--
					if tmp == 0 {
						b := y
						pba = append(pba, [2]int{a, b})
						break
					}
				}
			}
		}
		if tmp != 0 {
			fmt.Println("unbalanced ()")
		}
	}

	for x := 0; x < len(tStr); x++ {
		tmp := 0
		if tStr[x] == '{' {
			a := x
			tmp++
			for y := x + 1; y < len(tStr); y++ {
				if tStr[y] == '{' {
					tmp++
				} else if tStr[y] == '}' {
					tmp--
					if tmp == 0 {
						b := y
						cba = append(cba, [2]int{a, b})
						break
					}
				}
			}
		}
		if tmp != 0 {
			fmt.Println("unbalanced {}")
		}
	}

	for x := 0; x < len(tStr); x++ {
		tmp := 0
		if tStr[x] == '[' {
			a := x
			tmp++
			for y := x + 1; y < len(tStr); y++ {
				if tStr[y] == '[' {
					tmp++
				} else if tStr[y] == ']' {
					tmp--
					if tmp == 0 {
						b := y
						sba = append(sba, [2]int{a, b})
						break
					}
				}
			}
		}
		if tmp != 0 {
			fmt.Println("unbalanced []")
		}
	}

	fmt.Println(pba)
	fmt.Println(cba)
	fmt.Println(sba)

	for i := 0; i < len(pba); i++ {
		for j := i + 1; j < len(pba); j++ {
			fmt.Printf("Comparing pba%v to pba%v\n", pba[i], pba[j])
		}
		for k := 0; k < len(cba); k++ {
			fmt.Printf("Comparing pba%v to cba%v\n", pba[i], cba[k])
		}
		for l := 0; l < len(sba); l++ {
			fmt.Printf("Comparing pba%v to sba%v\n", pba[i], sba[l])
		}
	}
}

// Int pairs can reside outside eachother or inside each other
// ie [1,2] [5,6] is okay because those pairs are completely

/*func compareIntPairs(pair1 [2]int, pair2 [2]int ) bool {

	if pair1[0] < pair2[0] && pair[1] < pair1[1] {
		return true
	}

	if pair1[0] > pair2[1]




}*/
