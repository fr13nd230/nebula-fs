-- name: GetAllFiles :many
select * from files
limit $1 offset $2;