import sys
import csv
import time
import numpy as np

# Open and import the original dataset
dataLabels = []
dataMatrixL = []
labelsRef = []
with open('../dataset.csv', 'rb') as csvfile:
	reader = csv.reader(csvfile)
	rowIndx = -1
	for row in reader:
		if rowIndx >= 0:
			dataLabels.append(int(row[len(row) - 1]))
			if int(row[len(row) - 1]) not in labelsRef:
				labelsRef.append(int(row[len(row) - 1]))
		rowIndx = rowIndx + 1

# Open and import the resulting cendtroids (and time) file results
# from the command-line value
guessedLabels = []
timeVal = ''
with open(sys.argv[1], 'rb') as csvfile:
	reader = csv.reader(csvfile)
	rowIndx = -1
	for row in reader:
		if len(row) == 1:
			timeVal = row[0]
		else:
			for i in range(0, len(row)):
				guessedLabels.append(float(row[i]))	
		rowIndx = rowIndx + 1

# Display metrics for "scoring" the clustering results
from sklearn import metrics
hcv = metrics.homogeneity_completeness_v_measure(dataLabels, guessedLabels)
adjRandScore = metrics.adjusted_rand_score(dataLabels, guessedLabels)
print("")
print("Computational Time: " + timeVal + "ms")
#print("Average Distance from Centroids: " + str(np.sum(distVals)/len(distVals)))
#print("Points Clustered by Centroids:   " + str(distIndx))
print("")
print("Scores:")
print("   Homogeneity:    " + str(hcv[0]))
print("   Completeness:   " + str(hcv[1]))
print("   V-Measure:      " + str(hcv[2]))
print("   Adj. Rand:      " + str(adjRandScore))
print("")



