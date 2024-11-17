package models

type Author struct {
	ID   int64
	Name string
}

type AuthorWithBooks struct {
	ID    int64
	Name  string
	Books []Book
}

type Book struct {
	ID      int64
	Title   string
	Authors []Author
}
