import os
import csv
import time
import numpy as np

from sklearn.cluster import KMeans
filename = 'KMeansTest.txt'

# Open and import the original dataset
dataLabels = []
dataMatrixL = []
labelsRef = []
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
					dataLabels.append(int(row[i]))
					if int(row[i]) not in labelsRef:
						labelsRef.append(int(row[i]))	
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

# Write results to file
File = open(filename, 'w')
for i in range(0, len(kmeans.cluster_centers_)):
	for j in range(0, len(kmeans.cluster_centers_[i])):
		if j < (len(kmeans.cluster_centers_[i]) - 1):
			File.write(str(kmeans.cluster_centers_[i][j]) + ",")
		else:
			File.write(str(kmeans.cluster_centers_[i][j]))
	File.write('\n')
File.write(str(elapsedTime))
File.close()

# Run the metrics and delete the results
os.system('python CentroidLabelAssignment.py ' + filename)
os.system('python ClusterMetrics.py ' + filename)
os.remove(filename)

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
			
			





