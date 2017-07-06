package main

import (
	"fmt"
	"log"
)

//find the k-th element of two sorted array
//suppose all elements are in asc order
func kelem(s1, s2 []int, k int) (int, error) {
	log.Printf("search %d element in %v and %v\n", k, s1, s2)
	if k < 1 || k > len(s1)+len(s2) {
		return 0, fmt.Errorf("k must be in range [1,%d] (total number of two arrays)", len(s1)+len(s2))
	}

	if len(s1) == 0 {
		return s2[k-1], nil //k is 1 based
	}
	if len(s2) == 0 {
		return s1[k-1], nil
	}

	ms1 := len(s1) / 2
	ms2 := len(s2) / 2
	if k == ms1+ms2+2 { //ms1 and ms2 are zero based
		return max(s1[ms1], s2[ms2]), nil
	} else if k > ms1+ms2+2 {
		s1right := s1[ms1+1:] //right part of the array, exclude the mid element
		s2right := s2[ms2+1:]
		knew := k - ms1 - ms2 - 2
		log.Println("search in the right part")
		return kelem(s1right, s2right, knew)
	} else {
		var s1left, s2left []int
		//ignore the bigger mid element, as it's at leat k+1 big in two arrays
		if s1[ms1] < s2[ms2] {
			s1left = s1[:ms1+1]
			s2left = s2[:ms2]
		} else {
			s1left = s1[:ms1]
			s2left = s2[:ms2+1]
		}

		log.Println("search in the left part")
		return kelem(s1left, s2left, k)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := []int{2, 4, 6}
	k := 4
	ele, err := kelem(s1, s2, k)
	if err != nil {
		//print error and return
		log.Fatal(err.Error())
	}
	log.Printf("The %d element is: %d", k, ele)
}
