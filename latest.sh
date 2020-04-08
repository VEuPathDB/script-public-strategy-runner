#!/usr/bin/env bash

#
# Verification
#
if [ -z "$@" ]; then
  echo "At least one argument containing either the site url or \"-h\" required"
  exit 1
fi

if [[ "$OSTYPE" == "linux-gnu" ]]; then
  readonly OS="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  readonly OS="darwin"
else
  echo "Unsupported OS: $OSTYPE"
  exit 1
fi

#
# Constants
#
readonly GITHUB_API_PATH="https://api.github.com"
readonly REPO_API_PATH="${GITHUB_API_PATH}/repos/VEuPathDB/script-public-strategy-runner"
readonly REPO_TARGET="${REPO_API_PATH}/releases/latest"
readonly BINARY_NAME="public-strategies"


#
# Execution Time
#
readonly FILE_URL="$(curl -s "${REPO_TARGET}" \
  | grep "browser_download_url" \
  | grep "${OS}" \
  | cut -d '"' -f 4)"
readonly FILE_NAME="$(basename "${FILE_URL}")"


#
# Download and run the tool
#
wget -q "${FILE_URL}" \
  && tar -xzf "${FILE_NAME}" \
  && rm "${FILE_NAME}" \
  && ./public-strategies $@
rm -rf ./public-strategies