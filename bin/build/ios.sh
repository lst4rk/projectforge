#!/bin/bash

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

if [ "$XSKIP_MOBILE" != "true" ]
then
  echo "building gomobile for iOS..."
  time gomobile bind -o build/dist/mobile_ios_arm64/projectforge.framework -target=ios github.com/kyleu/projectforge/app/cmd
  echo "gomobile for iOS completed successfully, building distribution..."
  cd "build/dist/mobile_ios_arm64/projectforge.framework"
  zip --symlinks -r "../../projectforge_${TGT}_mobile_ios.zip" .
fi
