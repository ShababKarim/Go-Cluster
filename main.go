package main

import (
	"math/rand"

	"github.com/ShababKarim/go-cluster/kmeans"
)

func main() {
	sampleData := make([][]int, 60)
	for i := range sampleData {
		sampleData[i] = []int{rand.Intn(50)+1, rand.Intn(50)+1}
	}
	for i := 0; i < 20; i++ {
		// clearly separable
		sampleData[i+20][0] = sampleData[i+20][0] + 6
		sampleData[i+40][1] = sampleData[i+40][1] + 6
	}
	clusterCount := 3

	kmeans.InitialAssign(sampleData, clusterCount)
}