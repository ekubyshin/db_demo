-- name: AuthorsList :many
select
    *
from
    authors;

-- name: AuthorByID :one
select
    *
from
    authors
where
    id = @id
limit 1;

-- name: AuthorBooks :many
select
    b.id,
    b.title,
    a.id,
    a.name
from
    authors_books ab
    join books b on ab.book_id = b.id
    join authors_books ab2 on ab2.book_id = b.id
    join authors a on ab2.author_id = a.id
where
    ab.author_id = @author_id;

-- name: CreateAuthor :one
insert into authors(
    name)
values (
    @name)
returning
    id;

-- name: BatchCreateAuthors :batchexec
insert into authors(
    name)
values (
    @name);

-- name: UpdateAuthor :exec
update
    authors
set
    name = @name
where
    id = @id;

-- name: DeleteAuthor :exec
delete from authors
where id = @id;

