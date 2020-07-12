package main

import (
	"fmt"
	"math/rand"
)

// Copied sort from http://codeforces.com/contest/145/submission/1107770
func QuicksortTourist(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, x := l, r, a[l+rand.Intn(r-l+1)]
	for i <= j {
		for a[i] < x {
			i++
		}
		for a[j] > x {
			j--
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	QuicksortTourist(a, l, j)
	QuicksortTourist(a, i, r)
}

// Beautiful!
// To prove there is no infinite recursion, you have to prove this:
// - For any array that runs this code, "i++" is called at least once AND "j--" is called at least once
//   (Otherwise we could have i == l or j == r which would result in recurring on an array of the same size -- i.e. infinite recursion)
//
// The easiest way to do this is to prove that the third case (where we do the swap) is called at least once. You can
// prove this by contradiction: assume it is never called, and you will conclude that a[k] < pivot or a[k] > pivot for
// every k in the array, which is a contradiction because pivot IS inside this array.
func QuicksortShort(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, x := l, r, a[l+rand.Intn(r-l+1)]
	for i <= j {
		if a[i] < x {
			i++
		} else if a[j] > x {
			j--
		} else {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	QuicksortShort(a, l, j)
	QuicksortShort(a, i, r)
}

// Idea from https://en.wikipedia.org/wiki/Dutch_national_flag_problem
func QuicksortDutch(a []int, l, r int) {
	if l >= r {
		return
	}
	n := r - l + 1
	pivot := a[l+rand.Intn(n)]
	// Entries from 0 up to (but not including) i are values less than pivot,
	// entries from i up to (but not including) j are values equal to pivot,
	// entries from j up to (but not including) k are values not yet sorted, and
	// entries from k to the end of the array are values greater than pivot.
	i, j, k := l, l, l+n
	for j < k {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
			j++
		} else if a[j] > pivot {
			k--
			a[j], a[k] = a[k], a[j]
		} else { // a[j] == pivot
			j++
		}
	}
	QuicksortDutch(a, l, i-1)
	QuicksortDutch(a, k, r)
}

func SortAndPrint(a []int) {
	fmt.Println("Before sorting: ", a)
	QuicksortDutch(a, 0, len(a)-1)
	fmt.Println("After sorting: ", a)
	fmt.Println()
}

func main() {
	SortAndPrint([]int{5, 7, 1})
	SortAndPrint([]int{1, 1, 1, 1})
	SortAndPrint([]int{1, 2, 3, 4, 5})
	SortAndPrint([]int{5, 4, 3, 2, 1})
	SortAndPrint([]int{7, 2, 8, 14, 11, 13, 19, 20, 1, 9, 12, 18, 15, 6, 4, 10, 17, 5, 3, 16})
}
