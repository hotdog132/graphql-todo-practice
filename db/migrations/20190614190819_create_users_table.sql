-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- +goose StatementBegin
CREATE TABLE users (
  id integer primary key,
  name varchar(20) not null,
  created_at timestamp default localtimestamp,
  updated_at timestamp default localtimestamp
);

-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
