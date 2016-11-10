package utils

import "math"

type DistCentroid struct {
	indx1 int
	indx2 int
	dist  float64
}

func deleteCentroid(indx int, distCentroids []DistCentroid) []DistCentroid {
	distCentroids = append(distCentroids[:indx], distCentroids[indx+1:]...)
	return distCentroids
}

func deletDataSlice(indx int, dataSlice [][]float64) [][]float64 {
	dataSlice = append(dataSlice[:indx], dataSlice[indx+1:]...)
	return dataSlice
}

func averagePoints(pt1 []float64, pt2 []float64) []float64 {
	newPt := make([]float64, len(pt1))
	for i := 0; i < len(pt1); i++ {
		newPt[i] = (pt1[i] + pt2[i]) / 2
	}
	return newPt
}

func getDist(i1 int, i2 int, data [][]float64) float64 {
	total := 0.0
	for i := 0; i < len(data[i1]); i++ {
		total = total + ((data[i1][i] - data[i2][i]) * (data[i1][i] - data[i2][i]))
	}
	return math.Sqrt(total)
}

func calcDistances(distCentroids []DistCentroid, dataMatrix [][]float64) []DistCentroid {
	indx := 0
	for i := 0; i < len(dataMatrix)-1; i++ {
		for j := i + 1; j < len(dataMatrix); j++ {
			distance := getDist(i, j, dataMatrix)
			distCentroids[indx] = DistCentroid{i, j, distance}
			indx++
		}
	}
	return distCentroids
}

func getShortest(distCentroids []DistCentroid) DistCentroid {
	shortest := distCentroids[0]
	for i := 1; i < len(distCentroids); i++ {
		if distCentroids[i].dist < shortest.dist {
			shortest = distCentroids[i]
		}
	}
	return shortest
}

func createDistCentroidList(dataMatrix [][]float64) []DistCentroid {

	// Allocate the DistCentroid objects
	numPairs := (len(dataMatrix) * (len(dataMatrix) - 1)) / 2
	distCentroids := make([]DistCentroid, numPairs)

	// Calculate the distsances between each centroid
	distCentroids = calcDistances(distCentroids, dataMatrix)

	// Return the list
	return distCentroids
}

func CombineClusters(dataMatrix [][]float64, numClusters int) [][]float64 {

	// Get the centroid list
	distCentroids := createDistCentroidList(dataMatrix)

	// Combine the closest centroids until the list is the desired
	// number of centroids
	for len(dataMatrix) > numClusters {

		// Average the two closest centroids
		shortest := getShortest(distCentroids)
		pt1 := shortest.indx1
		pt2 := shortest.indx2
		newPt := averagePoints(dataMatrix[pt1], dataMatrix[pt2])

		// Combine them in the original matrix
		if pt2 > pt1 {
			tmp := pt1
			pt1 = pt2
			pt2 = tmp
		}
		dataMatrix = deletDataSlice(pt1, dataMatrix)
		dataMatrix[pt2] = newPt

		// Repeat distance calculations
		distCentroids = createDistCentroidList(dataMatrix)
	}

	// Return the new list of centroids
	return dataMatrix
}
