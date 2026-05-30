package contracts

import (
	"database/sql"
	"time"
)

type RDBMSFacade interface {
	NewConnection(constr string) (RDBMSFacade, error)
	Close() error
	Ping() error
	GetPool() *sql.DB
	Query(query string, args ...any) (*sql.Rows, error)
}

type Commandesult struct {
	Err          error
	RowsAffected int64
	TimeStamp    time.Time
}
