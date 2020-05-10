package validation

// Wss computes within-cluster-variation of a clustering solution
func Wss(clusterMap map[int][][]int) int{
	if len(clusterMap) == 0 {
		return -1
	}
	return 1
}