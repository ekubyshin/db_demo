package std

import (
	"errors"

	"github.com/ekubyshin/db_demo/models"
)

var queryListAuthors = `
SELECT * FROM authors
`

func (s *Storage) GetAuthors() ([]models.Author, error) {
	rows, err := s.db.Query(queryListAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	authors := make([]models.Author, 0)
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

var queryAuthor = `
SELECT * FROM authors WHERE id = $1 LIMIT 1
`

func (s *Storage) GetAuthor(id int64) (*models.Author, error) {
	var author models.Author
	err := s.db.QueryRow(queryAuthor, id).Scan(&author.ID, &author.Name)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

var queryAuthorBooks = `
	SELECT b.id, b.title, a.id, a.name
	FROM authors_books ab
	JOIN books b ON ab.book_id = b.id
	JOIN authors_books ab2 ON ab2.book_id = b.id
	JOIN authors a ON ab2.author_id = a.id
	WHERE ab.author_id = $1
`

func (s *Storage) GetAuthorBooks(id int64) ([]models.Book, error) {
	rows, err := s.db.Query(queryAuthorBooks, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := make([]models.Book, 0)
	var booksMap = make(map[int64]models.Book)
	for rows.Next() {
		var book models.Book
		var author models.Author
		err := rows.Scan(&book.ID, &book.Title, &author.ID, &author.Name)
		if err != nil {
			return nil, err
		}
		if book2, ok := booksMap[book.ID]; ok {
			book2.Authors = append(book2.Authors, author)
			booksMap[book.ID] = book2
		} else {
			book.Authors = []models.Author{author}
			booksMap[book.ID] = book
		}
	}
	for _, book := range booksMap {
		books = append(books, book)
	}
	return books, nil
}

var queryCreateAuthor = `
	INSERT INTO authors (name) VALUES ($1) RETURNING id
`

func (s *Storage) CreateAuthor(author models.Author) (*models.Author, error) {
	var id int64
	err := s.db.QueryRow(queryCreateAuthor, author.Name).Scan(&id)
	if err != nil {
		return nil, err
	}
	author.ID = id
	return &author, nil
}

var queryUpdateAuthor = `
	UPDATE authors SET name = $1 WHERE id = $2
`

func (s *Storage) UpdateAuthor(author models.Author) (*models.Author, error) {
	_, err := s.db.Exec(queryUpdateAuthor, author.Name, author.ID)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

var queryDeleteAuthor = `
	DELETE FROM authors WHERE id = $1
`

func (s *Storage) DeleteAuthor(id int64) error {
	_, err := s.db.Exec(queryDeleteAuthor, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) BalkCreateAuthor(authors []models.Author) ([]models.Author, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() //nolint
	result := make([]models.Author, 0, len(authors))
	for _, author := range authors {
		var id int64
		err := tx.QueryRow(queryCreateAuthor, author.Name).Scan(&id)
		if err != nil {
			return nil, err
		}
		author.ID = id
		result = append(result, author)
	}
	if len(result) != len(authors) {
		return nil, errors.New("not all authors created")
	}
	return result, tx.Commit()
}
