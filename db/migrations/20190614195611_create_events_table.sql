-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- +goose StatementBegin
CREATE TABLE events (
  id serial primary key,
  user_id integer not null references users(id),
  text varchar(150) not null,
  done Boolean not null default false,
  created_at timestamp default localtimestamp,
  updated_at timestamp default localtimestamp
);

-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose StatementBegin
DROP TABLE events;
-- +goose StatementEnd
