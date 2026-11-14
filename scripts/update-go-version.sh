#!/usr/bin/env bash

set -euo pipefail

readonly VERSION_INPUT="${1:-}"

# Manage temporary files robustly, even if the script crashes or dies early.
declare -a CLEANUP_FILES=()
cleanup() {
  if [[ ${#CLEANUP_FILES[@]} -gt 0 ]]; then
    rm -f "${CLEANUP_FILES[@]}"
  fi
}
trap cleanup EXIT

# FIX: Redirect to stderr (>&2) so command substitutions $(...) don't capture log output
log() {
  echo "[update-go-version] $*" >&2
}

die() {
  echo "[update-go-version] ERROR: $*" >&2
  exit 1
}

require_file() {
  local file="$1"
  [[ -f "$file" ]] || die "Required file not found: $file"
}

require_cmd() {
  local cmd="$1"
  command -v "$cmd" >/dev/null 2>&1 || die "Required command not found: $cmd"
}

validate_version() {
  [[ "$VERSION_INPUT" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || die "Invalid Go version '$VERSION_INPUT'. Expected format: MAJOR.MINOR.PATCH"
}

replace_or_fail() {
  local file="$1"
  local match_regex="$2"
  local sed_expr="$3"

  grep -Eq "$match_regex" "$file" || die "Expected pattern not found in $file"
  sed -E -i "$sed_expr" "$file"
}


update_known_version_files() {
  log "Updating Go version references in known files"

  replace_or_fail "go.mod" '^go [0-9]+\.[0-9]+\.[0-9]+$' "s/^go [0-9]+\.[0-9]+\.[0-9]+$/go ${VERSION_INPUT}/"
  replace_or_fail "go.tool.mod" '^go [0-9]+\.[0-9]+\.[0-9]+$' "s/^go [0-9]+\.[0-9]+\.[0-9]+$/go ${VERSION_INPUT}/"
  replace_or_fail ".golangci.yaml" '^[[:space:]]*go:[[:space:]]*"?[0-9]+\.[0-9]+(\.[0-9]+)?"?[[:space:]]*$' "s/^([[:space:]]*go:[[:space:]]*).*/\\1\"${VERSION_INPUT}\"/"
}

# Constructs and returns the fully assembled replacement string,
# preventing the need for global state variables.
resolve_base_image() {
  local parsed
  parsed="$(extract_repo_and_stage_from_dockerfile)"

  local repo="${parsed%%$'\t'*}"
  local stage_suffix="${parsed#*$'\t'}"

  local tag
  tag="$(find_latest_dated_tag "$repo")"

  local image_ref="${repo}:${tag}"
  log "Resolving digest for $image_ref"

  local digest
  digest="$(resolve_digest_with_docker "$image_ref")"

  printf 'FROM %s@%s%s\n' "$image_ref" "$digest" "$stage_suffix"
}


show_detected_go_version_refs() {
  log "Scanning tracked files for Go version keys"
  grep -HnE '(^go [0-9]+\.[0-9]+\.[0-9]+$|^[[:space:]]*go:[[:space:]]*"?[0-9]+\.[0-9]+(\.[0-9]+)?"?[[:space:]]*$|^[[:space:]]*(go|golang)[[:space:]]*=[[:space:]]*"[^"]+"[[:space:]]*$)' go.mod go.tool.mod .golangci.yaml  2>/dev/null || true
}

main() {
  validate_version

  # Ensure file availability
  require_file "go.mod"
  require_file "go.tool.mod"
  require_file ".golangci.yaml"

  # Ensure tool availability
  require_cmd grep
  require_cmd sed
  require_cmd awk
  require_cmd curl
  require_cmd jq
  require_cmd mktemp

  show_detected_go_version_refs
  update_known_version_files

  log "Done. Updated Go references to $VERSION_INPUT"
}

main "$@"
