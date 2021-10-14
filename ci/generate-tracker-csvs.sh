#!/usr/bin/env bash

set -euo pipefail

source_dir=${1}
dest_dir=${2}

find ${source_dir}*.prolific -name '* *' | while read fname

do
        new_fname=`echo $fname | tr " " "_"`

        if [ -e $new_fname ]
        then
                echo "File $new_fname already exists. Not replacing $fname"
        else
                echo "Creating new file $new_fname to replace $fname"
                mv "$fname" $new_fname
        fi
done

echo "Looking in ${source_dir} for .prolific files"
for file in ${source_dir}*.prolific; do

  echo "Processing $file"
  without_extension="$(basename "$file" .prolific)"
  csv_filepath="${dest_dir}${without_extension}.csv"
  prolific "$file" > "${csv_filepath}"
  echo "Wrote ${csv_filepath}"
done

echo "Done."
