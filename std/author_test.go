package std

import (
	"testing"

	dblib "github.com/ekubyshin/db_demo/db"
	"github.com/ekubyshin/db_demo/models"
	"github.com/stretchr/testify/require"
)

func TestStorage_Authors(t *testing.T) {
	dsn, close := dblib.UpTestingDB(t)
	defer close()
	db := NewSQLDB(dsn)
	defer db.Close()
	err := dblib.Migrate(db)
	require.NoError(t, err)

	t.Run("GetAuthors empty", func(t *testing.T) {
		storage := NewStorage(db)
		authors, err := storage.GetAuthors()
		require.NoError(t, err)
		require.Equal(t, 0, len(authors))
	})

	t.Run("test getAuthors list", func(t *testing.T) {
		fillTestData(db, t)
		s := NewStorage(db)
		got, err := s.GetAuthors()
		require.NoError(t, err)
		want := []models.Author{
			{ID: 1, Name: "Author 1"},
			{ID: 2, Name: "Author 2"},
			{ID: 3, Name: "Author 3"},
		}
		require.Equal(t, want, got)
		got2, err := s.GetAuthor(1)
		require.NoError(t, err)
		require.Equal(t, models.Author{ID: 1, Name: "Author 1"}, got2)
	})

	t.Run("GetAuthorBooks", func(t *testing.T) {
		clearDB(db, t)
		fillTestData(db, t)
		s := NewStorage(db)
		got, err := s.GetAuthorBooks(1)
		require.NoError(t, err)
		want := []models.Book{
			{ID: 1, Title: "Book 1", Authors: []models.Author{{ID: 1, Name: "Author 1"}}},
			{ID: 2, Title: "Book 2", Authors: []models.Author{{ID: 1, Name: "Author 1"}, {ID: 2, Name: "Author 2"}}},
		}
		require.Equal(t, want, got)
	})
}
