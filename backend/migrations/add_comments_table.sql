-- +migrate Up
CREATE TABLE comments (
  id varchar(255),
  topic_id varchar(255),
  body varchar(255)
);

-- +migrate Down
DROP TABLE IF EXISTS comments;

