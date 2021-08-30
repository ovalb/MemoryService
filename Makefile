.PHONY: conndb, nukedb
conndb:
	PGPASSWORD=example psql -h localhost -d postgres -U postgres

nukedb:
	PGPASSWORD=example psql -h localhost -d postgres -U postgres -c "drop schema public cascade" -c "create schema public"
