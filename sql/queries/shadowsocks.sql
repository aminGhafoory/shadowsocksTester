-- name: CreateShadowsocks :one
INSERT INTO ss(
    created_at ,
    updated_at ,
    ssLink,
    ip,
    sub_id,
    port ,
    secret,
    name  

)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8

)
RETURNING *;


-- name: GetNextSSToTest :many
SELECT 
    * 
    FROM 
        ss 
ORDER BY 
    updated_at ASC 
LIMIT $1;


-- name: UpdateDestinationIP :one
UPDATE ss
SET 
    api_reported_ip = $1,
    updated_at = NOW()
WHERE 
    id = $2
RETURNING *;

-- name: GetAllSSs :many
SELECT * FROM ss;

-- name: GetBestList :many
SELECT 
    ss.id,
    ss.secret,
    ss.ip,
    ss.port,
    ss.name,
    ss.ssLink,
    avg(response_time) as avg_response_time,
	sum(case is_successful when true then 1 else 0 end) as successful_count,
	sum(case is_successful when false then 1 else 0 end) as failure_count
FROM
    ss JOIN reqs ON ss.id = reqs.ss_id

GROUP BY 
	ss.id
HAVING 
    avg(response_time)>0
ORDER BY 
	successful_count Desc,
    avg_response_time ASC;


