package std

import (
	"database/sql"
	"testing"

	dblib "github.com/ekubyshin/db_demo/db"
	"github.com/stretchr/testify/require"
)

func clearDB(db *sql.DB, t *testing.T) {
	_, err := db.Exec("DELETE FROM authors WHERE true")
	require.NoError(t, err)
}

func fillTestData(db *sql.DB, t *testing.T) {
	require.NoError(t, dblib.FillTestData(db))
}

func TestSQL(t *testing.T) {
	dsn, close := dblib.UpTestingDB(t)
	defer close()
	db := NewSQLDB(dsn)
	defer db.Close()
	err := dblib.Migrate(db)
	require.NoError(t, err)
}
