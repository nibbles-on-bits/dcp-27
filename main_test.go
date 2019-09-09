package main

import (
	"fmt"
	"testing"
)

func TestCompareIntPairs(t *testing.T) {

	pa := [2]int{0, 3}
	pb := [2]int{4, 8}
	got := CompareIntPairs(pa, pb)
	if got != true {
		t.Error(fmt.Sprintf("pa=%v pb=%v got=%v", pa, pb, got))
	}

	pa = [2]int{5, 6}
	pb = [2]int{4, 8}
	got = CompareIntPairs(pa, pb)
	if got != true {
		t.Error(fmt.Sprintf("pa=%v pb=%v got=%v", pa, pb, got))
	}

	pa = [2]int{2, 10}
	pb = [2]int{4, 8}
	got = CompareIntPairs(pa, pb)
	if got != true {
		t.Error(fmt.Sprintf("pa=%v pb=%v got=%v", pa, pb, got))
	}

	pa = [2]int{6, 10}
	pb = [2]int{4, 8}
	got = CompareIntPairs(pa, pb)
	if got != false {
		t.Error(fmt.Sprintf("pa=%v pb=%v got=%v", pa, pb, got))
	}
}

func TestIsWellBalanced(t *testing.T) {

	shouldPass := []string{"()", "[]", "{}", "({})", "[[(({}))]]"}
	for _, s := range shouldPass {
		if IsWellBalanced(s) == false {
			t.Error(fmt.Sprintf("Test failed %v should return true", s))
		}
	}

	shouldFail := []string{")", "(", "[({})", "{()})", "[][({})", "{[[[(({}))]]"}
	for _, s := range shouldFail {
		if IsWellBalanced(s) == false {
			t.Error(fmt.Sprintf("Test failed %v should return true", s))
		}
	}

}
