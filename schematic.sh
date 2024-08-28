#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <problem_number>"
    exit 1
fi

# Format problem number with zero-padding to 3 digits for directories and main file
problem_number_padded=$(printf "%03d" $1)
file_name="leetcode/$problem_number_padded/$problem_number_padded.go"

# Use the original problem number for the test file naming
file_name_test="leetcode/$problem_number_padded/${1}_test.go"

# Check if the problem directory exists, if not, create it
if [ ! -d "leetcode/$problem_number_padded" ]; then
    mkdir -p "leetcode/$problem_number_padded"
fi

# Check if the problem file already exists
if [ -e $file_name ]; then
    echo "Problem $problem_number_padded already exists!"
    exit 1
fi

# Copy the template and rename the module inside it
cp template/problem.go $file_name
cp template/problem_test.go $file_name_test

echo "Problem $1 created in directory $problem_number_padded"
