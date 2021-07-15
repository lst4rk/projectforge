#!/bin/bash

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

if [ "$XSKIP_MOBILE" != "true" ]
then
  echo "building gomobile for Android..."
  mkdir -p build/dist/mobile_android_arm64
  time gomobile bind -o build/dist/mobile_android_arm64/projectforge.aar -target=android github.com/kyleu/projectforge/app/cmd
  echo "gomobile for Android completed successfully, building distribution..."
  cd "build/dist/mobile_android_arm64"
  zip -r "../projectforge_${TGT}_mobile_android.zip" .
fi
