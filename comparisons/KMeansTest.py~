import csv
import time
import numpy as np

from sklearn.cluster import KMeans

# Create a vectorized function for calculating distances
import math
def sqrDiff(x1, x2):
	return (x2-x1)*(x2-x1)
vectSqrDiff = np.vectorize(sqrDiff)
def distFunc(v1, v2):
	return math.sqrt(np.sum(vectSqrDiff(v1, v2)))

# Open and import the original dataset
dataLabels = []
dataMatrixL = []
with open('../dataset.csv', 'rb') as csvfile:
	reader = csv.reader(csvfile)
	rowIndx = -1
	for row in reader:
		if rowIndx >= 0:
			dataMatrixL.append(np.zeros((len(row) - 1,), dtype=np.float64))
			for i in range(0, len(row)):
				if i < (len(row) - 1):
					dataMatrixL[rowIndx][i] = float(row[i])
				else:
					dataLabels.append(row[i])	
		rowIndx = rowIndx + 1
dataMatrix = np.zeros((len(dataMatrixL), len(dataMatrixL[0])), dtype=np.float64)
row = 0
for item in dataMatrixL:
	dataMatrix[row][:] = item[:]
	row = row + 1

# Run KMeans on the data and obtain the centroids (and time it)
startTime = time.time()
kmeans = KMeans(n_clusters=6, random_state=0).fit(dataMatrix)
elapsedTime = 1000*(time.time() - startTime)

# Determine the distances between each point and the nearest centriod
distVals = np.zeros((len(dataMatrixL),), dtype=np.float64)
distIndx = np.zeros((len(kmeans.cluster_centers_),), dtype=np.int32)
for i in range(0, len(dataMatrixL)):
	minDist = distFunc(kmeans.cluster_centers_[0], dataMatrixL[i])
	minIndx = 0;
	for j in range(1, len(kmeans.cluster_centers_)):
		dist = distFunc(kmeans.cluster_centers_[j], dataMatrixL[i])
		if dist < minDist:
			minDist = dist
			minIndx = j
	distVals[i] = minDist
	distIndx[minIndx] = distIndx[minIndx] + 1


# Display metrics for "scoring" the clustering results
print("Computational Time:              " + str(elapsedTime) + 'ms')
print("Average Distance from Centroids: " + str(np.sum(distVals)/len(distVals)))
print("Points Clustered by Centroids:   " + str(distIndx))

# Plotting for 2D data
#xValues = np.zeros((len(dataMatrixL),), dtype=np.float64)
#yValues = np.zeros((len(dataMatrixL),), dtype=np.float64)
#for i in range(0, len(dataMatrixL)):
#	xValues[i] = dataMatrixL[i][0]
#	yValues[i] = dataMatrixL[i][1]
#xCentroids = np.zeros((len(kmeans.cluster_centers_),), dtype=np.float64)
#yCentroids = np.zeros((len(kmeans.cluster_centers_),), dtype=np.float64)
#for i in range(0, len(kmeans.cluster_centers_)):
#	xCentroids[i] = kmeans.cluster_centers_[i][0]
#	yCentroids[i] = kmeans.cluster_centers_[i][1]
#
#import matplotlib.pyplot as plt
#plt.plot(xValues, yValues, 'rx')
#plt.plot(xCentroids, yCentroids, 'bo')
#plt.show()
			
			





