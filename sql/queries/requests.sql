-- name: TestSS :one
INSERT INTO reqs (
    source,
    ss_id,
    response_time,
    is_successful
)VALUES(
    $1,$2,$3,$4
)
RETURNING *;
