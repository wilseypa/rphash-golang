#!/bin/sh
# $1 - master port.
# $2 - volume port.
# $3 - data directory.

# Set up Seaweed.
weed server -master.port=$1 -volume.port=$2 -dir="$3"
