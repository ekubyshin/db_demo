INSERT INTO authors (id, name)
VALUES (1, 'Author 1'), (2, 'Author 2'), (3, 'Author 3');
SELECT setval('authors_id_seq', max(id))
FROM authors;

INSERT INTO books (id, title)
VALUES (1, 'Book 1'), (2, 'Book 2'), (3, 'Book 3');
SELECT setval('books_id_seq', max(id))
FROM books;

INSERT INTO authors_books (author_id, book_id)
VALUES (1, 1), (1, 2), (2, 2), (3, 3);