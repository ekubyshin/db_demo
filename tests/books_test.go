package std

import (
	"testing"

	dblib "github.com/ekubyshin/db_demo/db"
	"github.com/ekubyshin/db_demo/models"
	"github.com/ekubyshin/db_demo/std"
	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/require"
)

func TestStorage_Books(t *testing.T) {
	dsn, close := dblib.UpTestingDB(t)
	defer close()
	db := std.NewDB(dsn)
	defer db.Close()
	err := dblib.Migrate(db)
	require.NoError(t, err)

	t.Run("GetBooks", func(t *testing.T) {
		clearDB(db, t)
		fillTestData(db, t)
		s := std.NewStorage(db)
		got, err := s.GetBooks()
		require.NoError(t, err)
		autogold.ExpectFile(t, got)
	})

	t.Run("GetBook", func(t *testing.T) {
		clearDB(db, t)
		fillTestData(db, t)
		s := std.NewStorage(db)
		type args struct {
			id int64
		}
		tests := []struct {
			name    string
			args    args
			want    *models.Book
			wantErr bool
		}{
			{
				name: "Book 1",
				args: args{
					id: 1,
				},
				want: &models.Book{
					ID:    1,
					Title: "Book 1",
					Authors: []models.Author{
						{ID: 1, Name: "Author 1"},
					},
				},
				wantErr: false,
			},
			{
				name: "Book 2",
				args: args{
					id: 2,
				},
				want: &models.Book{
					ID:    2,
					Title: "Book 2",
					Authors: []models.Author{
						{ID: 1, Name: "Author 1"},
						{ID: 2, Name: "Author 2"},
					},
				},
				wantErr: false,
			},
			{
				name: "Book 3",
				args: args{
					id: 3,
				},
				want: &models.Book{
					ID:    3,
					Title: "Book 3",
					Authors: []models.Author{
						{ID: 3, Name: "Author 3"},
					},
				},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := s.GetBook(tt.args.id)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			})
		}
	})
}
