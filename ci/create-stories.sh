#!/usr/bin/env bash

set -euo pipefail

: "${TRACKER_TOKEN:?Variable not set or empty}"
: "${TRACKER_PROJECT_ID:?Variable not set or empty}"

api_url="https://www.pivotaltracker.com/services/v5"

while IFS= read -r line; do
  curl \
    --request POST \
    --header "X-TrackerToken: $TRACKER_TOKEN" \
    --header "Content-Type: application/json" \
    --data "$line" \
    "$api_url/projects/$TRACKER_PROJECT_ID/stories"  
done <<< "$(cat jsons/*.json)"

