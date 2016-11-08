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
	for i := 0; i < len(distCentroids)-1; i++ {
		for j := i + 1; j < len(distCentroids); j++ {
			distance := getDist(i, j, dataMatrix)
			distCentroids[indx] = DistCentroid{i, j, distance}
			indx++
		}
	}
	return distCentroids
}

func sortList(distCentroids []DistCentroid, numPairs int) []DistCentroid {
	for i := 0; i < numPairs-1; i++ {
		for j := i + 1; j < numPairs; j++ {
			if distCentroids[i].dist > distCentroids[j].dist {
				tmp := distCentroids[i]
				distCentroids[i] = distCentroids[j]
				distCentroids[j] = tmp
			}
		}
	}
	return distCentroids
}

func createDistCentroidList(dataMatrix [][]float64) []DistCentroid {

	// Allocate the DistCentroid objects
	numPairs := (len(dataMatrix) * (len(dataMatrix) - 1)) / 2
	distCentroids := make([]DistCentroid, numPairs)

	// Calculate the distsances between each centroid
	distCentroids = calcDistances(distCentroids, dataMatrix)

	// Sort the list
	distCentroids = sortList(distCentroids, numPairs)

	// Return the list
	return distCentroids
}

func CombineClusters(dataMatrix [][]float64, numClusters int) [][]float64 {

	// Get the centroid list
	distCentroids := createDistCentroidList(dataMatrix)

	// Combine the closest centroids until the list is the desired
	// number of centroids
	indx := 0
	for i := 0; i < numClusters; i++ {

		// Average the two closest centroids
		pt1 := distCentroids[0].indx1
		pt2 := distCentroids[0].indx2
		newPt := averagePoints(dataMatrix[pt1], dataMatrix[pt2])

		// Combine them in the original matrix
		deletDataSlice(pt1, dataMatrix)
		dataMatrix[pt2] = newPt

		// Repeat distance calculations
		distCentroids = createDistCentroidList(dataMatrix)
		indx++
	}

	// Return the new list of centroids
	return dataMatrix
}
