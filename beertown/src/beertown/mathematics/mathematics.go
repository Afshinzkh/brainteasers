package mathematics

import (
	"fmt"
	"math"
	"sort"
)

func FindBeers(n int) (goodBeers []int) {

	if n > 1000 {
		fmt.Println("That's a huge number, go get a beer until I'm done!")
	}
	// we know 1,2,3 are not the ones so we start from 4
	if n < 4 {
		return nil
	}

	for i := 4; i <= n; i++ {

		if isGood(i) {
			goodBeers = append(goodBeers, i)
		}

	}

	return goodBeers
}

func isGood(val int) bool {
	devisors := []int{1}
	sum := 0
	for i := 2; i <= int(math.Sqrt(float64(val))); i++ {
		if val%i == 0 {
			if i == val/i {
				sum += i
				devisors = append(devisors, i)
			} else {
				sum += (i + val/i)
				devisors = append(devisors, i)
				devisors = append(devisors, val/i)
			}

		}
	}

	if sum+1 <= val {
		return false
	}

	sort.Ints(devisors)

	if isSubSetSum(devisors, len(devisors), val) {
		return false
	}

	return true
}

func isSubSetSum(devisors []int, len int, val int) bool {
	if val <= 0 {
		return true
	}

	if len == 0 && val > 0 {
		return false
	}

	if devisors[len-1] > val {
		return isSubSetSum(devisors, len-1, val)
	}
	return isSubSetSum(devisors, len-1, val) ||
		isSubSetSum(devisors, len-1, val-devisors[len-1])
}
