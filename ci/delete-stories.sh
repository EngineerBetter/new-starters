#!/usr/bin/env bash

set -euo pipefail

: "${TRACKER_TOKEN:?Variable not set or empty}"
: "${TRACKER_PROJECT_ID:?Variable not set or empty}"

api_url="https://www.pivotaltracker.com/services/v5"

ids="$(curl \
    --silent \
    --show-error \
    --request GET \
    --header "X-TrackerToken: $TRACKER_TOKEN" \
    "$api_url/projects/$TRACKER_PROJECT_ID/stories" \
    | jq --raw-output '.[].id')"

xargs -I{} \
    curl \
        --silent \
        --show-error \
        --request DELETE \
        --header "X-TrackerToken: $TRACKER_TOKEN" \
        --header "Content-Type: application/json" \
        "$api_url/projects/$TRACKER_PROJECT_ID/stories/{}" \
        <<< "$ids"