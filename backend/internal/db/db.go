package db

import (
	"database/sql"
	"errors"
	"os"

	"go.uber.org/zap"
	_ "modernc.org/sqlite" // Runs init function which registers the driver with database/sql under the name "sqlite"
)

type DbBaseDao struct {
	logger *zap.SugaredLogger
	dbConn *sql.DB
}

func NewDbBaseDao(logger *zap.SugaredLogger, dbConn *sql.DB) *DbBaseDao {
	return &DbBaseDao{
		logger: logger,
		dbConn: dbConn,
	}
}

// OpenDb creates a connection to the SQLite db
func OpenDb() (*sql.DB, error) {
	dbFilePath := os.Getenv("SQLITE_FILEPATH")
	if dbFilePath == "" {
		return nil, errors.New("nothing at SQLITE_FILEPATH")
	}
	return sql.Open("sqlite", dbFilePath)
}

type schemaEntry struct {
	name string
	ddl  string
}

func (d *DbBaseDao) DbHealth() error {
	pingErr := d.dbConn.Ping()
	if pingErr != nil {
		d.logger.Errorf("db health: ", pingErr)
		return pingErr
	}
	rows, err := d.dbConn.Query("SELECT name, sql FROM sqlite_schema limit 1")
	if err != nil {
		d.logger.Errorf("db health: couldn't query: %w", err)
		return err
	}
	defer rows.Close()
	if c, err := rows.Columns(); err != nil {
		d.logger.Errorf("db health: couldn't get cols: %w", err)
		return err
	} else {
		d.logger.Infof("cols: %v", c)
	}
	for rows.Next() {
		var schemaEntry schemaEntry
		if err := rows.Scan(&schemaEntry.name, &schemaEntry.ddl); err != nil {
			d.logger.Errorf("couldn't read to schemaEntry: %v", err)
			return err
		}
		d.logger.Infof("schemaEntry: %v, %v", schemaEntry.name, schemaEntry.ddl)
	}
	if err := rows.Err(); err != nil {
		d.logger.Errorf("db health: row iteration: %v", err)
		return err
	}
	return nil
}
