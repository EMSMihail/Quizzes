#!/bin/bash

host="db"
port=5433

echo "Waiting for database to become available..."

while ! nc -z $host $port; do
  sleep 1
done

echo "Database is now available. Starting the application."

exec "$@"
