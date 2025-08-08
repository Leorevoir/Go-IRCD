#!/usr/bin/env bash

set -e

function _assert()
{
    if $1; then
        echo "SUCCESS: $1"
    else
        >&2 echo "ERROR: assertion failed: $1"
    fi
}

_assert "go build -o ircd-server ./cmd/server"
_assert "go build -o ircd-client ./cmd/client"
