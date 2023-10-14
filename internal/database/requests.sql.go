// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: requests.sql

package database

import (
	"context"
)

const testSS = `-- name: TestSS :one
INSERT INTO reqs (
    source,
    ss_id,
    response_time,
    is_successful
)VALUES(
    $1,$2,$3,$4
)
RETURNING id, ss_id, source, response_time, is_successful
`

type TestSSParams struct {
	Source       string
	SsID         int64
	ResponseTime int32
	IsSuccessful bool
}

func (q *Queries) TestSS(ctx context.Context, arg TestSSParams) (Req, error) {
	row := q.db.QueryRowContext(ctx, testSS,
		arg.Source,
		arg.SsID,
		arg.ResponseTime,
		arg.IsSuccessful,
	)
	var i Req
	err := row.Scan(
		&i.ID,
		&i.SsID,
		&i.Source,
		&i.ResponseTime,
		&i.IsSuccessful,
	)
	return i, err
}