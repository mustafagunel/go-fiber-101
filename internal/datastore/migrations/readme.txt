$ goose create add_some_column sql
$ Created new file: 20170506082420_add_some_column.sql

Apply all available migrations.
$ goose up
Migrate up to a specific version.
$ goose up-to 20170506082420
Migrate up a single migration from the current version
$ goose up-by-one
Roll back a single migration from the current version.
$ goose down
Roll back migrations to a specific version.
$ goose down-to 20170506082527
Roll back the most recently applied migration, then run it again.
$ goose redo
Print the status of all migrations:
$ goose status



Usage: goose [OPTIONS] DRIVER DBSTRING COMMAND

or

Set environment key
GOOSE_DRIVER=DRIVER
GOOSE_DBSTRING=DBSTRING
GOOSE_MIGRATION_DIR=MIGRATION_DIR

Usage: goose [OPTIONS] COMMAND

Drivers:
    postgres
    mysql
    sqlite3
    mssql
    redshift
    tidb
    clickhouse
    vertica
    ydb
    duckdb

Examples:
    goose sqlite3 ./foo.db status
    goose sqlite3 ./foo.db create init sql
    goose sqlite3 ./foo.db create add_some_column sql
    goose sqlite3 ./foo.db create fetch_user_data go
    goose sqlite3 ./foo.db up

    goose postgres "user=postgres dbname=postgres sslmode=disable" status
    goose mysql "user:password@/dbname?parseTime=true" status
    goose redshift "postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db" status
    goose tidb "user:password@/dbname?parseTime=true" status
    goose mssql "sqlserver://user:password@dbname:1433?database=master" status
    goose clickhouse "tcp://127.0.0.1:9000" status
    goose vertica "vertica://user:password@localhost:5433/dbname?connection_load_balance=1" status
    goose ydb "grpcs://localhost:2135/local?go_query_mode=scripting&go_fake_tx=scripting&go_query_bind=declare,numeric" status
    goose duckdb ./foo.db status

    GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./foo.db goose status
    GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./foo.db goose create init sql
    GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=postgres dbname=postgres sslmode=disable" goose status
    GOOSE_DRIVER=mysql GOOSE_DBSTRING="user:password@/dbname" goose status
    GOOSE_DRIVER=redshift GOOSE_DBSTRING="postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db" goose status

Options:

  -allow-missing
        applies missing (out-of-order) migrations
  -certfile string
        file path to root CA's certificates in pem format (only support on mysql)
  -dir string
        directory with migration files (default ".", can be set via the GOOSE_MIGRATION_DIR env variable).
  -h    print help
  -no-color
        disable color output (NO_COLOR env variable supported)
  -no-versioning
        apply migration commands with no versioning, in file order, from directory pointed to
  -s    use sequential numbering for new migrations
  -ssl-cert string
        file path to SSL certificates in pem format (only support on mysql)
  -ssl-key string
        file path to SSL key in pem format (only support on mysql)
  -table string
        migrations table name (default "goose_db_version")
  -timeout duration
        maximum allowed duration for queries to run; e.g., 1h13m
  -v    enable verbose mode
  -version
        print version

Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
    validate             Check migration files without running them




-- +goose Up
CREATE TABLE post (
    id int NOT NULL,
    title text,
    body text,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE post;




-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION histories_partition_creation( DATE, DATE )
returns void AS $$
DECLARE
  create_query text;
BEGIN
  FOR create_query IN SELECT
      'CREATE TABLE IF NOT EXISTS histories_'
      || TO_CHAR( d, 'YYYY_MM' )
      || ' ( CHECK( created_at >= timestamp '''
      || TO_CHAR( d, 'YYYY-MM-DD 00:00:00' )
      || ''' AND created_at < timestamp '''
      || TO_CHAR( d + INTERVAL '1 month', 'YYYY-MM-DD 00:00:00' )
      || ''' ) ) inherits ( histories );'
    FROM generate_series( $1, $2, '1 month' ) AS d
  LOOP
    EXECUTE create_query;
  END LOOP;  -- LOOP END
END;         -- FUNCTION END
$$
language plpgsql;
-- +goose StatementEnd


https://github.com/pressly/goose/tree/master/examples/go-migrations