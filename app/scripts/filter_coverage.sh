#!/bin/bash
# Usage: ./filter_coverage.sh coverage.out coverage_filtered.out

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <coverage_in> <coverage_out>"
    exit 1
fi

COV_IN="$1"
COV_OUT="$2"

# Check if the coverage input file exists
if [ ! -f "$COV_IN" ]; then
    echo "Coverage file $COV_IN does not exist."
    exit 1
fi

# Check if .covignore exists
if [ ! -f ".covignore" ]; then
    echo ".covignore not found. No filtering applied."
    cp "$COV_IN" "$COV_OUT"
    exit 0
fi

# Create a grep pattern from all lines in .covignore
# Ignores empty lines and comments starting with #
PATTERN=$(grep -Ev '^\s*($|#)' .covignore | paste -sd'|' -)

# If no pattern is found, copy the original coverage
if [ -z "$PATTERN" ]; then
    cp "$COV_IN" "$COV_OUT"
    exit 0
fi

# Filter coverage file
# Excludes any line containing the paths/prefixes in .covignore
grep -Ev "$PATTERN" "$COV_IN" > "$COV_OUT"

echo "Filtered coverage generated in $COV_OUT"
