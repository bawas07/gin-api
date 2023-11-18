echo "$DB_URL"
migrate -database "$DB_URL" -path db/migrations version