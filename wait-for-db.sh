#!/bin/sh

# Wait until SQL Server is ready
while ! nc -z db 1433; do
  echo "Waiting for db to be up..."
  sleep 5
done

echo "db is up!"
