package std

import "github.com/ekubyshin/db_demo/models"

var queryListBooks = `
SELECT b.id, b.title, a.id, a.name
FROM books b
JOIN authors_books ab ON ab.book_id = b.id
JOIN authors a ON ab.author_id = a.id
`

func (s *Storage) GetBooks() ([]models.Book, error) {
	rows, err := s.db.Query(queryListBooks)
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

func (s *Storage) GetBook(id int64) (*models.Book, error) {
	rows, err := s.db.Query(queryListBooks+" WHERE b.id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	book := models.Book{
		Authors: make([]models.Author, 0),
	}
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&book.ID, &book.Title, &author.ID, &author.Name)
		if err != nil {
			return nil, err
		}
		book.Authors = append(book.Authors, author)
	}

	return &book, nil
}

func (s *Storage) CreateBook(book models.Book) (*models.Book, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() //nolint
	return nil, nil
}
