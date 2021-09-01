#!/usr/bin/env bash

set -euo pipefail

: "${TRACKER_TOKEN:?Variable not set or empty}"
: "${TRACKER_PROJECT_ID:?Variable not set or empty}"

api_url="https://www.pivotaltracker.com/services/v5"

ids="$(curl -X GET -H "X-TrackerToken: $TRACKER_TOKEN" "$api_url/projects/$TRACKER_PROJECT_ID/stories" | jq -r '.[].id')" 

xargs -I{} curl -X DELETE -H "X-TrackerToken: $TRACKER_TOKEN" -H "Content-Type: application/json" "$api_url/projects/$TRACKER_PROJECT_ID/stories/{}" <<< "$ids"