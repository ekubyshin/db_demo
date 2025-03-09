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

func TestSQLCStorage_Books(t *testing.T) {
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

	t.Run("GetBooks", func(t *testing.T) {
		clearDB(stddb, t)
		fillTestData(stddb, t)
		got, err := db.BooksList(ctx)
		require.NoError(t, err)
		autogold.ExpectFile(t, got)
	})

	t.Run("GetBook", func(t *testing.T) {
		clearDB(stddb, t)
		fillTestData(stddb, t)
		type args struct {
			id int64
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "Book 1",
				args: args{
					id: 1,
				},
				wantErr: false,
			},
			{
				name: "Book 2",
				args: args{
					id: 2,
				},
				wantErr: false,
			},
			{
				name: "Book 3",
				args: args{
					id: 3,
				},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := db.AuthorBooks(ctx, int32(tt.args.id))
				require.NoError(t, err)
				autogold.ExpectFile(t, got, autogold.Name("TestSQLCStorage_Books/GetBooks_"+tt.name))
			})
		}
	})
}
