#!/bin/sh
set -e

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

host="$1"
shift
cmd="$@"

until mysql -h "$host" -u"$DB_USER" -p"$DB_PASS" -e 'SELECT 1' 2>&1; do
  >&2 echo "MySQL is unavailable - sleeping"
  sleep 2
done

>&2 echo "MySQL is up - executing command"
exec $cmd
