#!/usr/bin/env bash

set -euo pipefail

cd new-starters/tracker-csv-to-json

for csv in ../../csvs/*.csv; do
    json_file="../../jsons/$(basename "$csv" .csv).json"

    go run main.go "$csv" > "$json_file"
    
done
