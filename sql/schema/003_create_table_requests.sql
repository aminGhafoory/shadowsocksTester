-- +goose Up
CREATE TABLE reqs(
    id bigserial PRIMARY KEY,
    ss_id bigserial NOT NULL,
    source TEXT NOT NULL,
    response_time int NOT NULL default 0,
    is_successful bool NOT NULL,
    FOREIGN KEY(ss_id) 
	  REFERENCES ss(id)
);

-- +goose Down
DROP TABLE IF EXISTS reqs;