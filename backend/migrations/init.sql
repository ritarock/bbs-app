-- +migrate Up
CREATE TABLE topics (
  id varchar(255),
  title varchar(255),
  detail varchar(255)
);

-- +migrate Down
DROP TABLE IF EXISTS topics;
