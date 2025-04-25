#!/bin/bash

set -euo pipefail

# Check for the latest Terraform Provider release
latest_release=$(gh release list -R scaleway/terraform-provider-scaleway --limit 1)
latest_version=$(echo "$latest_release" | awk '{print $1}' | sed 's/^v//')
echo "Latest version: $latest_version"
export new_version="$latest_version"

# If running under GH Actions, also write it into $GITHUB_ENV
if [[ -n "${GITHUB_ENV-}" ]]; then
  echo "new_version=$latest_version" >> "$GITHUB_ENV"
fi
