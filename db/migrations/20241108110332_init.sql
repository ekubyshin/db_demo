-- +goose Up
-- +goose StatementBegin
create table if not exists authors (
    id serial primary key,
    name text not null
);
create table if not exists books (
    id serial primary key,
    title text not null
);
create table if not exists authors_books (
    author_id int not null references authors (id) on delete cascade,
    book_id int not null references books (id) on delete cascade
);
create unique index if not exists authors_books_unique_idx on authors_books (author_id, book_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table authors_books;
drop table books;
drop table authors;
-- +goose StatementEnd
