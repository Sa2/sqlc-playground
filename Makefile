migrate-up:
	@echo "Migrating local database"
	@sql-migrate up 
	@docker run -it --rm -v $PWD:/root -w /root --network sqlc-playground_fixed_network postgres /bin/bash -c 'export PGPASSWORD=postgres && pg_dump -h 10.0.3.10 -U postgres -s -t public.*' > schema/db/sql/schema.sql

migrate-down:
	@echo "Rolling back local database"
	@sql-migrate down
	@docker run -it --rm -v $PWD:/root -w /root --network sqlc-playground_fixed_network postgres /bin/bash -c 'export PGPASSWORD=postgres && pg_dump -h 10.0.3.10 -U postgres -s -t public.*' > schema/db/sql/schema.sql

generate-sqlc:
	@echo "Generating SQL orm code"
	@docker pull sqlc/sqlc
	@docker run --rm -v $PWD:/src -w /src sqlc/sqlc generate

