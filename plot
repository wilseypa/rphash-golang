#!/bin/sh
OPT=$1
echo "Plotting ..."
eval "go run ./tests/plot_$1.go";
echo "Plot finished, output [cluster.png]"
