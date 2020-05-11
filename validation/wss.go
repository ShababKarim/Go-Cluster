package validation

import (
	"errors"
	"math"
	"sync"

	l "github.com/ShababKarim/go-cluster/lib"
)

// TotalWss computes within-cluster-variation of a clustering solution
func TotalWss(clusterMap map[int][][]int, numFeatures int) (float64, error) {
	if len(clusterMap) == 0 {
		return -1, errors.New("Invalid cluster map size")
	}

	var total float64
	var wg sync.WaitGroup

	for _, cluster := range clusterMap {
		wg.Add(1)

		go func(cl [][]int) {
			l.AddFloat64(&total, wss(cl, numFeatures))
			wg.Done()
		}(cluster)
	}

	wg.Wait()
	return total, nil
}

func wss(singleCluster [][]int, numFeatures int) float64 {
	totalSum := 0.0
	for i := 0; i < len(singleCluster); i++ {
		for j := i+1; j < len(singleCluster); j++ {
			sum := 0.0
			for p := 0; p < numFeatures; p++ {
				sum += math.Pow(math.Abs(float64(singleCluster[i][p]) - float64(singleCluster[j][p])), 2)
			}
			totalSum += math.Sqrt(sum)
		}
	}
	return totalSum / float64(len(singleCluster))
}