CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    role int not null default(0),
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE INDEX name_index ON users USING btree(name);

CREATE TABLE IF NOT EXISTS access_token (
    token text PRIMARY KEY not null,
    user_id BIGINT not null,
    created_at timestamp not null,
    valid_thru timestamp not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);