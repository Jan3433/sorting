package mergesort

// Merge kombiniert zwei sortierte Listen zu einer sortierten Liste.
func Merge(a1 []int, a2 []int) []int {
	result := []int{}
	i := 1
	for j := i; i < len(a1) && j < len(a2); j++ {
		if a1[i] < a2[i] {
			a1[i] = result[j]
		} else {
			result[i] = a2[j]
		}
	}
	return result
}

//a1 und a2 zu a3
// a1 und a2 müssen sortiert sein
//

// MergeSort sortiert die übergebene Liste mittels des Merge-Sort-Algorithmus.
func MergeSort(arr []int) []int {
	// TODO
	return []int{}
}
