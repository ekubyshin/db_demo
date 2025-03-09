-- name: BooksList :many
select
    b.id,
    b.title,
    a.id,
    a.name
from
    books b
    join authors_books ab on ab.book_id = b.id
    join authors a on ab.author_id = a.id;

-- name: BookByID :one
select
    b.id,
    b.title,
    a.id,
    a.name
from
    books b
    join authors_books ab on ab.book_id = b.id
    join authors a on ab.author_id = a.id
where
    b.id = @id
limit 1;

