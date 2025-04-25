#!/usr/bin/env bash
set -euo pipefail

# ------------------------------------------------------------------
# scripts/update_provider.sh
#
# This will:
#   1. check for a new TF provider release (via ./scripts/check_version.sh)
#   2. bump Makefile versions
#   3. save old provider-metadata.yaml
#   4. regenerate with `make generate`
#   5. diff old vs new metadata, code-gen any new resources
#   6. run go mod tidy + go fmt
# ------------------------------------------------------------------

COMMIT_CHANGES=false
if [[ "${1:-}" = "--commit" ]]; then
  COMMIT_CHANGES=true
fi

# using BSD or GNU sed style
if sed --version >/dev/null 2>&1; then
  SED_INPLACE=(-i)
else
  SED_INPLACE=(-i '')
fi

echo "➡️  Checking for new Terraform provider release…"
# check_version.sh must export NEW_VERSION env var
source ./scripts/check-version.sh
if [[ -z "${new_version:-}" ]]; then
  echo "✋ check_version.sh did not set \$new_version" >&2
  exit 1
fi

current_version=$(awk -F ' := ' '/TERRAFORM_PROVIDER_VERSION/{print $2}' Makefile)
echo "   current: $current_version"
echo "   latest : $new_version"

if [[ "$new_version" == "$current_version" ]]; then
  echo "✅ Already up-to-date; nothing to do."
  exit 0
fi

echo "🔧 Bumping Makefile to v$new_version"
sed "${SED_INPLACE[@]}" \
  -e "s|^export TERRAFORM_PROVIDER_VERSION := .*|export TERRAFORM_PROVIDER_VERSION := ${new_version}|" \
  -e "s|^export TERRAFORM_NATIVE_PROVIDER_BINARY := .*|export TERRAFORM_NATIVE_PROVIDER_BINARY := terraform-provider-scaleway_v${new_version}|" \
  Makefile

# back up old metadata
cp config/provider-metadata.yaml config/provider-metadata.old.yaml

echo "🔄 Regenerating provider metadata…"
make generate

echo "🔍 Diffing metadata for new resources…"
export CURRENT_METADATA="$(< config/provider-metadata.old.yaml)"
export NEW_METADATA="$(< config/provider-metadata.yaml)"
# run comparator and capture its JSON output
raw=$(
  go run ./config/tools/comparator/main.go \
    --current config/provider-metadata.old.yaml \
    --new   config/provider-metadata.yaml
)
resource_configs=$(printf '%s\n' "$raw" | tail -n1)

if [[ "$resource_configs" != "[]" && -n "$resource_configs" ]]; then
  echo "✨ Generating configs…"
  printf '%s' "$resource_configs" \
    | go run ./config/tools/generator/main.go
else
  echo "⏭️ No new resources"
fi

echo "🧹 Tidy & format Go modules"
go mod tidy
go fmt ./...

echo "🎉 Done."
