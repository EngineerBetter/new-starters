#!/usr/bin/env bash

set -euo pipefail

: "${TRACKER_TOKEN:?Variable not set or empty}"
: "${TRACKER_PROJECT_ID:?Variable not set or empty}"

api_url='https://www.pivotaltracker.com/services/v5'

while IFS= read -r line; do
  echo -n "Creating story \"$( jq -r '.name' <<< "$line" )\"..."

  curl \
    --silent \
    --show-error \
    --request POST \
    --header "X-TrackerToken: $TRACKER_TOKEN" \
    --header 'Content-Type: application/json' \
    --data "$line" \
    "$api_url/projects/$TRACKER_PROJECT_ID/stories" \
    1> /dev/null

  echo " done!"
done <<< "$( cat $(ls jsons/*.json | sort --reverse) | grep name )"
