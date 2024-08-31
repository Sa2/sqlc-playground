migrate-up:
	@echo "Migrating local database"
	@sql-migrate up 

migrate-down:
	@echo "Rolling back local database"
	@sql-migrate down
