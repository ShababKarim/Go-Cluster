package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/ShababKarim/go-cluster/kmeans"
)

func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func main() {
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

	clusters, err := kmeans.InitialAssign(sampleData, clusterCount)
	if err != nil {
		fmt.Println("Error: ", err.Error())
        os.Exit(2)
	}
	cl, err := kmeans.Update(clusters, 2)
	if err != nil {
		fmt.Println("Error: ", err.Error())
        os.Exit(2)
	}
	fmt.Println(cl)
}