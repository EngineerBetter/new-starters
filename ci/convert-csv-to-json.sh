#!/usr/bin/env bash

set -euo pipefail

go install github.com/jiangmiao/csv2json@latest

for csv in csvs/*.csv
do
    json_file="jsons/$(basename "$csv" .csv).json"
    jq_query='[.[] '\
'| .["name"] = .Title '\
'| .["story_type"] = .Type '\
'| .["description"] = .Description '\
'| .["labels"] = .Labels '\
'| del(.Title, .Type, .Description, .Labels)] '\
'| map(.labels |= split(","))'

    csv2json < "$csv" | jq "$jq_query" > "$json_file"
done
