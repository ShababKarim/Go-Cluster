package kmeans

import (
	"fmt"
	"math/rand"
)

func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func ExampleNew() {
	sampleData := make([][]int, 60)
	for i := range sampleData {
		sampleData[i] = []int{randomInt(0, 10), randomInt(0, 10)}
	}
	for i := 0; i < 20; i++ {
		// clearly separable
		sampleData[i+20][0] = sampleData[i+20][0] + 6
		sampleData[i+40][1] = sampleData[i+40][1] + 6
	}
	clusterCount := 3

	cmap, _ := New(sampleData, clusterCount)

	fmt.Printf("Within cluster variation: %.3f", cmap.Wss)
	// Output: Within cluster variation: 731.525
}