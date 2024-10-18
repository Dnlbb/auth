#!/bin/bash
source .env

export MIGRATION_DSN="host=$DB_HOST port=5432 dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

# Попробуйте накатывать миграции до успешного завершения
until goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v; do
  echo "Migration failed, retrying in 5 seconds..."
  sleep 5
done

echo "Migrations applied successfully!"
