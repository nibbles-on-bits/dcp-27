package main

import (
	"fmt"
)

func main() {
	fmt.Println("DCP-27")
	res := IsWellBalanced("((){[]}()[{((()))}])")
	fmt.Println(res)
}

func IsWellBalanced(s string) bool {

	pba := [][2]int{} // holds parenthesis pairs
	cba := [][2]int{} // holds curly brace pairs
	sba := [][2]int{} // holds square bracket pairs

	for x := 0; x < len(s); x++ {
		tmp := 0
		if s[x] == '(' {
			a := x
			tmp++
			for y := x + 1; y < len(s); y++ {
				if s[y] == '(' {
					tmp++
				} else if s[y] == ')' {
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
			return false
		}
	}

	for x := 0; x < len(s); x++ {
		tmp := 0
		if s[x] == '{' {
			a := x
			tmp++
			for y := x + 1; y < len(s); y++ {
				if s[y] == '{' {
					tmp++
				} else if s[y] == '}' {
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
			return false
		}
	}

	for x := 0; x < len(s); x++ {
		tmp := 0
		if s[x] == '[' {
			a := x
			tmp++
			for y := x + 1; y < len(s); y++ {
				if s[y] == '[' {
					tmp++
				} else if s[y] == ']' {
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
			return false
		}
	}

	fmt.Println(pba)
	fmt.Println(cba)
	fmt.Println(sba)

	for i := 0; i < len(pba); i++ {
		for j := i + 1; j < len(pba); j++ {
			fmt.Printf("Comparing pba%v to pba%v\n", pba[i], pba[j])
			if !CompareIntPairs(pba[i], pba[j]) {
				return false
			}
		}
		for k := 0; k < len(cba); k++ {
			fmt.Printf("Comparing pba%v to cba%v\n", pba[i], cba[k])
			if !CompareIntPairs(pba[i], cba[k]) {
				return false
			}
		}
		for l := 0; l < len(sba); l++ {
			fmt.Printf("Comparing pba%v to sba%v\n", pba[i], sba[l])
			if !CompareIntPairs(pba[i], sba[l]) {
				return false
			}
		}
	}
	return true
}

// Int pairs can reside outside eachother or inside each other
// ie [1,2] [5,6] is okay because those pairs are completely
func CompareIntPairs(pa [2]int, pb [2]int) bool {
	if isBetween(pa[0], pb[0], pb[1]) {
		if isBetween(pa[1], pb[0], pb[1]) {
			return true
		}
		return false
	}

	return true
}

func isBetween(i int, x int, y int) bool {
	if i > x && i < y {
		return true
	}
	return false
}
