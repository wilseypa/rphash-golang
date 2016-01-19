#!/bin/sh
OPT=$1
echo "Plotting ..."
eval "go run ./plots/plot_$1.go";
echo "Plot finished, output [kmeans.png]"
