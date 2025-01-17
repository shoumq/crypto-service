#!/usr/bin/env sh
# wait-for-it.sh

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until nc -z -v -w30 $host $port; do
  echo "Waiting for $host:$port..."
  sleep 1
done

>&2 echo "$host:$port is up - executing command"
exec $cmd
