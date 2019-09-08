package main

import (
	"fmt"
)

func main() {
	fmt.Println("DCP-27")

	pba := [][2]int{} // holds parenthesis pairs
	//cba := [][2]int{} // holds curly brace pairs
	//sba := [][2]int{} // holds square bracket pairs

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
	}

	fmt.Println(pba)

}
