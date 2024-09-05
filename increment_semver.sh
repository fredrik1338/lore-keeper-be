#!/bin/bash

# Input semver in the format "x.y.z"
semver=$1
if [ "$1" = "" ]; then
    # $var is empty
    echo "1.0.0"
    exit 0
fi

# Split semver into major, minor, and patch parts
IFS='.' read -r -a parts <<< "$semver"
major="${parts[0]}"
minor="${parts[1]}"
patch="${parts[2]}"

# Increment minor version
((patch++))

# Output the new semver
echo "$major.$minor.$patch"