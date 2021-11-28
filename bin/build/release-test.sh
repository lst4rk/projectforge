#!/bin/bash

## Runs goreleaser in test mode

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

[[ -f "$HOME/bin/oauth" ]] && . $HOME/bin/oauth

export PUBLISH_TEST=true
goreleaser -f ./tools/release/.goreleaser.yml --snapshot --timeout 60m --skip-publish --rm-dist
unset PUBLISH_TEST
