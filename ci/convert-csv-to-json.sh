#!/usr/bin/env bash

set -euo pipefail

go get github.com/jiangmiao/csv2json

for csv in csvs/*.csv
do
    json_file="jsons/$(basename "$csv" .csv).json"
    csv2json < "$csv" > "$json_file"
done
