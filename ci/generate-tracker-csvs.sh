#!/usr/bin/env bash

set -euo pipefail

source_dir=${1}
dest_dir=${2}

echo "Looking in ${source_dir} for .prolific files"
for yaml in ${source_dir}/*.prolific; do
  echo "Processing $yaml"
  without_extension="$(basename "$yaml" .prolific)"
  csv_filepath="${dest_dir}${without_extension}.csv"
  prolific "$yaml" > "${csv_filepath}"
  echo "Wrote ${csv_filepath}"
done

echo "Done."