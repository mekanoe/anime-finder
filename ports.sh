#!/bin/bash

SET=export
SPLIT="="

if [[ "$SHELL" == *fish* ]]; then
    SET="set -gx"
    SPLIT=" "
fi

echo "$SET NATS_ADDR$SPLIT$(docker-compose port nats 4222);"
echo "$SET DGRAPH_ADDR$SPLIT$(docker-compose port dgraph 9080);"
echo "$SET CRDB_ADDR$SPLIT$(docker-compose port crdb1 26257);"