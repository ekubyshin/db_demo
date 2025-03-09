package std

import (
	"testing"
	"time"

	dblib "github.com/ekubyshin/db_demo/db"
	"github.com/ekubyshin/db_demo/gorm"
	"github.com/hexops/autogold/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
	g "gorm.io/gorm"
)

func TestGormStorage_Authors(t *testing.T) {
	dsn, close := dblib.UpTestingDB(t)
	defer close()
	db, err := gorm.NewDB(dsn)
	cfg, err := pgx.ParseConfig(dsn)
	require.NoError(t, err)
	stddb := stdlib.OpenDB(*cfg)
	require.NoError(t, err)
	require.NoError(t, err)
	fillTestData(stddb, t)

	t.Run("getAuthors list", func(t *testing.T) {
		authors := []gorm.Author{}
		_ = db.Select("*").Preload("Books").Find(&authors)
		ignoreDate(authors)
		autogold.ExpectFile(t, authors)
	})

	t.Run("getAuthorsByID", func(t *testing.T) {
		var authors []gorm.Author
		_ = db.Select("*").Where("id = ?", 1).Preload("Books").Find(&authors)
		ignoreDate(authors)
		autogold.ExpectFile(t, authors, autogold.Name("TestGormStorage_Authors/getAuthorsByID"))
	})

	t.Run("CreateAuthor", func(t *testing.T) {
		author := gorm.Author{Name: "Author test"}
		_ = db.Create(&author)
		authors := []gorm.Author{}
		_ = db.Select("*").Preload("Books").Find(&authors)
		ignoreDate(authors)
		autogold.ExpectFile(t, authors)
	})

	t.Run("BalkCreateAuthor", func(t *testing.T) {
		authors := []gorm.Author{
			{Name: "Author test1"},
			{Name: "Author test2"},
			{Name: "Author test3"},
		}
		_ = db.Create(&authors)
		lst := []gorm.Author{}
		_ = db.Select("*").Preload("Books").Find(&lst)
		ignoreDate(authors)
		autogold.ExpectFile(t, authors)
	})
}

func ignoreDate(authors []gorm.Author) {
	t, _ := time.Parse(time.RFC3339, "")
	for i := 0; i < len(authors); i++ {
		authors[i].CreatedAt = t
		authors[i].UpdatedAt = t
		authors[i].DeletedAt = g.DeletedAt{
			Time:  t,
			Valid: true,
		}
	}
}
