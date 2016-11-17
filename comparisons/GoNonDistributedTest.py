import os
import sys

filename = 'NonDistResults.txt'
os.system('go run ../rphash_cl.go ../dataset.csv ' + filename)
os.system('python CentroidLabelAssignment.py ' + filename)
os.system('python ClusterMetrics.py ' + filename)
os.remove(filename)
