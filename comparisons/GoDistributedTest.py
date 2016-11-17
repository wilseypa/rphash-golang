import os
import sys

filename = 'DistResults_' + sys.argv[1] + '.txt'
os.system('go run ../rphash_cl.go ../dataset.csv ' + filename + ' ' + sys.argv[1])
os.system('python CentroidLabelAssignment.py ' + filename)
os.system('python ClusterMetrics.py ' + filename)
os.remove(filename)
