
-- +goose Up
CREATE TABLE ss(
    id bigserial PRIMARY KEY,
    sub_id serial NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    ssLink TEXT NOT NULL,
    ip varchar(512) NOT NULL,
    port INT NOT NULL,
    secret varchar(512) NOT NULL,
    name varchar(128) NOT NULL,
    api_reported_ip varchar(128) ,
    unique(ip,port),
    FOREIGN KEY(sub_id) 
	  REFERENCES subs(id)
);


-- +goose Down

DROP TABLE IF EXISTS ss;
