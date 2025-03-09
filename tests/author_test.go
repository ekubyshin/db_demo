package std

import (
	"testing"

	dblib "github.com/ekubyshin/db_demo/db"
	"github.com/ekubyshin/db_demo/models"
	"github.com/ekubyshin/db_demo/std"
	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/require"
)

func TestStorage_Authors(t *testing.T) {
	dsn, close := dblib.UpTestingDB(t)
	defer close()
	db := std.NewDB(dsn)
	defer db.Close()
	err := dblib.Migrate(db)
	require.NoError(t, err)
	fillTestData(db, t)
	storage := std.NewStorage(db)

	t.Run("getAuthors list", func(t *testing.T) {
		got, err := storage.GetAuthors()
		require.NoError(t, err)
		autogold.ExpectFile(t, got)
	})

	t.Run("getAuthorsByID", func(t *testing.T) {
		author, err := storage.GetAuthor(1)
		require.NoError(t, err)
		autogold.ExpectFile(t, author, autogold.Name("TestStorage_Authors/getAuthorsByID"))
	})

	t.Run("GetAuthorBooks", func(t *testing.T) {
		got, err := storage.GetAuthorBooks(1)
		require.NoError(t, err)
		autogold.ExpectFile(t, got)
	})

	t.Run("CreateAuthor", func(t *testing.T) {
		author := models.Author{Name: "Author test"}
		_, err := storage.CreateAuthor(author)
		require.NoError(t, err)
		lst, err := storage.GetAuthors()
		autogold.ExpectFile(t, lst)
	})

	t.Run("UpdateAuthor", func(t *testing.T) {
		author := models.Author{Name: "Author test2"}
		got, err := storage.CreateAuthor(author)
		require.NoError(t, err)
		got.Name = "Author test2 updated"
		got, err = storage.UpdateAuthor(*got)
		require.NoError(t, err)
		lst, err := storage.GetAuthors()
		require.NoError(t, err)
		autogold.ExpectFile(t, lst)
	})

	t.Run("BalkCreateAuthor", func(t *testing.T) {
		authors := []models.Author{
			{Name: "Author test1"},
			{Name: "Author test2"},
			{Name: "Author test3"},
		}
		_, err := storage.BalkCreateAuthor(authors)
		require.NoError(t, err)
		lst, err := storage.GetAuthors()
		require.NoError(t, err)
		autogold.ExpectFile(t, lst)
	})
}
