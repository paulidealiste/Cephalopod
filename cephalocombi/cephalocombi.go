// Package cephalocombi provides simple combinatorics electors
package cephalocombi

// CartesianProduct provides Cartesian product between any number of integer vectors
func CartesianProduct(vectors [][]int) [][]int {
	arraynumber := len(vectors)

	digits := make([]int, arraynumber)
	arraylens := make([]int, arraynumber)
	var totalcombinations uint64
	totalcombinations = 1

	for i, vec := range vectors {
		aln := len(vec)
		digits[i] = 0
		totalcombinations = totalcombinations * uint64(aln)
		arraylens[i] = aln
	}

	combined := make([][]int, totalcombinations)
	var idx uint64
	for idx = 0; idx < totalcombinations; idx++ {
		item := make([]int, arraynumber)
		for idy := 0; idy < arraynumber; idy++ {
			item[idy] = vectors[idy][digits[idy]]
		}
		combined[idx] = item
		for idz := 0; idz < arraynumber; idz++ {
			if digits[idz] == arraylens[idz]-1 {
				digits[idz] = 0
			} else {
				digits[idz] = digits[idz] + 1
				break
			}
		}
	}
	return combined
}
