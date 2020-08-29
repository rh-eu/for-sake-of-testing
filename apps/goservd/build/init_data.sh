#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

mkdir -p /data/go

echo ${ARCH}

mkdir -p "/data/std/${ARCH}"

mkdir -p "${npm_config_cache}"

chown -R ${TARGET_UIDGID} /data
chmod -R a=rwX /data