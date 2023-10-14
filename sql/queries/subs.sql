
-- name: AddToSubs :one
INSERT INTO subs(
    created_at,
    updated_at,
    url
)
VALUES(
    $1,$2,$3
)
RETURNING *;

-- name: GetNextSubsToFetch :many
SELECT * FROM subs 
ORDER BY updated_at ASC 
LIMIT $1;

-- name: GetAllSubs :many
SELECT * FROM subs
ORDER BY created_at DESC;

