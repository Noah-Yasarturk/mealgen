package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type DbBaseDao struct {
	logger *zap.SugaredLogger
	dbConn *pgx.Conn
}

func NewDbBaseDao(logger *zap.SugaredLogger, dbConn *pgx.Conn) *DbBaseDao {
	return &DbBaseDao{
		logger: logger,
		dbConn: dbConn,
	}
}

// OpenDb creates a connection to the Postgres db
func OpenDb(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("open db: %v", err)
	}
	return conn, nil
}

type schemaEntry struct {
	tableSchema string
	tableName   string
}

func (d *DbBaseDao) DbHealth(ctx context.Context) error {
	pingErr := d.dbConn.Ping(ctx)
	if pingErr != nil {
		d.logger.Errorf("db health: ", pingErr)
		return pingErr
	}
	rows, err := d.dbConn.Query(ctx, "SELECT table_schema, table_name FROM INFORMATION_SCHEMA.tables limit 10")
	if err != nil {
		d.logger.Errorf("db health: couldn't query: %w", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var schemaEntry schemaEntry
		if err := rows.Scan(&schemaEntry.tableSchema, &schemaEntry.tableName); err != nil {
			d.logger.Errorf("couldn't read to schemaEntry: %v", err)
			return err
		}
		d.logger.Infof("schemaEntry: schema - %v, table - %v", schemaEntry.tableSchema, schemaEntry.tableName)
	}
	if err := rows.Err(); err != nil {
		d.logger.Errorf("db health: row iteration: %v", err)
		return err
	}
	return nil
}
