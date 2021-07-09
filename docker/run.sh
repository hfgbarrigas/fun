#!/bin/sh

set -e

[ -z "$DATA_PATH" ] && echo "DATA_PATH not set. " && exit 1;

exec litecoind -datadir="$DATA_PATH"

