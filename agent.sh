#!/bin/sh
# $1 - master host and port.
# $2 - Available disk space.
# $3 - data directory.

# Set up Seaweed.
weed volume -max=$2 -mserver="$1" -dir="$3"
