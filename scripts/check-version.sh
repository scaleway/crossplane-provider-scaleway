#!/bin/bash

# Check for the latest Terraform Provider release
latest_release=$(gh release list -R scaleway/terraform-provider-scaleway --limit 1)
latest_version=$(echo $latest_release | awk '{print $1}' | sed 's/^v//')
echo "Latest version: $latest_version"
echo "new_version=$latest_version" >> $GITHUB_ENV
