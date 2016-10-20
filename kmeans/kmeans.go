package kmeans

import (
	"container/list"
	"math/rand"

	"github.com/gonum/matrix/mat64"
)

// KMeansMat structure definition
type KMeansMat struct {
	matCols    int          // Integer used to store the data's dimensionality
	clusters   int          // Integer used to store the number of clusers
	matrixList *list.List   // Linked list used to store the point data
	centroids  *mat64.Dense // Matrix used to store the centroids of the clusters
}

// Function used to return an instance of the KMeansMat structure
func create(cols int, k int) KMeansMat {
	initList := list.New()
	cents := mat64.NewDense(cols, k, nil)
	newMat := KMeansMat{cols, k, initList, cents}
	return newMat
}

// Function used to add data to the list maintained for the data (add point)
func (mat KMeansMat) addDataVect(data []float64) {
	mat.matrixList.PushBack(data)
}

// Function used to get a value from the list for the data setup
func (mat KMeansMat) getMatrixListVal(indx int, moveToBack bool) []float64 {

	// Traverse the list to find the desireed element
	elem := mat.matrixList.Front()
	for i := 0; i < (indx - 1); i++ {
		elem = elem.Next()
	}

	// Extract the value to return
	valueToReturn := elem.Value.([]float64)

	// Move the current item to the back of the list (if desired)
	if moveToBack {
		mat.matrixList.MoveToBack(elem)
	}

	// Return the requested value
	return valueToReturn
}

// Function used to get randomized centroids based on the setup data
func (mat KMeansMat) getInitCentroids() {

	// Keep track of the size of the list (decremented each time an item is
	// extracted from the list to ensure the same point isn't chosen twice)
	size := mat.matrixList.Len()

	// Populate the centroids field of the KMeansMat structure
	for indx := 0; indx < mat.clusters; indx++ {

		// Obtain a random value (that hasn't been chosen yet) from the list
		randInt := rand.Int() * size
		rowData := mat.getMatrixListVal(randInt, true)
		size = size - 1

		// Set the centroid data from the list value (point)
		mat.centroids.SetRow(indx, rowData)
	}
}

// Function used to perform the KMeans alorithm on the data (cluster it)
func (mat KMeansMat) performKMeans() {
	listLen := mat.matrixList.Len()
	for indx := 0; indx < listLen; indx++ {

	}
}
