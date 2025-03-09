package std

import (
	"context"
	"testing"

	dblib "github.com/ekubyshin/db_demo/db"
	"github.com/ekubyshin/db_demo/sqlc"
	"github.com/hexops/autogold/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
)

func TestSQLCStorage_Authors(t *testing.T) {
	dsn, close := dblib.UpTestingDB(t)
	defer close()
	ctx := context.Background()
	cfg, err := pgx.ParseConfig(dsn)
	require.NoError(t, err)
	stddb := stdlib.OpenDB(*cfg)
	require.NoError(t, err)
	dbx, err := pgx.Connect(ctx, dsn)
	require.NoError(t, err)
	db := sqlc.New(dbx)
	defer dbx.Close(ctx)
	err = dblib.Migrate(stddb)
	require.NoError(t, err)
	fillTestData(stddb, t)

	t.Run("getAuthors list", func(t *testing.T) {
		got, err := db.AuthorsList(ctx)
		require.NoError(t, err)
		autogold.ExpectFile(t, got)
	})

	t.Run("getAuthorsByID", func(t *testing.T) {
		author, err := db.AuthorByID(ctx, 1)
		require.NoError(t, err)
		autogold.ExpectFile(t, author, autogold.Name("TestSQLCStorage/getAuthorsByID"))
	})

	t.Run("GetAuthorBooks", func(t *testing.T) {
		got, err := db.AuthorBooks(ctx, 1)
		require.NoError(t, err)
		autogold.ExpectFile(t, got)
	})

	t.Run("CreateAuthor", func(t *testing.T) {
		_, err := db.CreateAuthor(ctx, "Author sqlc test")
		require.NoError(t, err)
		lst, err := db.AuthorsList(ctx)
		autogold.ExpectFile(t, lst)
	})

	t.Run("UpdateAuthor", func(t *testing.T) {
		got, err := db.CreateAuthor(ctx, "Author sqlc test2")
		require.NoError(t, err)
		err = db.UpdateAuthor(ctx, sqlc.UpdateAuthorParams{
			ID:   got,
			Name: "Author test2 updated",
		})
		require.NoError(t, err)
		lst, err := db.AuthorsList(ctx)
		require.NoError(t, err)
		autogold.ExpectFile(t, lst)
	})

	t.Run("BalkCreateAuthor", func(t *testing.T) {
		err := db.BatchCreateAuthors(ctx, []string{"Author sqlc test1", "Author sqlc test2", "Author sqlc test3"}).Close()
		require.NoError(t, err)
		lst, err := db.AuthorsList(ctx)
		require.NoError(t, err)
		autogold.ExpectFile(t, lst)
	})
}
