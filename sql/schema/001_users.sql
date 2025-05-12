-- +goose Up
CREATE TABLE users (
	id UUID UNIQUE NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL, 
	updated_at TIMESTAMP NOT NULL, 
	name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
