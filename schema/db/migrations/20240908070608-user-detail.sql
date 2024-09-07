
-- +migrate Up
CREATE TABLE user_details (
  id VARCHAR(36) PRIMARY KEY,
  user_id VARCHAR(36) NOT NULL,
  detail_info1 varchar(255) NOT NULL,
  detail_info2 varchar(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT user_details_ibfk_1 FOREIGN KEY (user_id) REFERENCES users (id)
);


-- +migrate Down
ALTER TABLE user_details DROP FOREIGN KEY user_details_ibfk_1;
DROP TABLE user_details;
