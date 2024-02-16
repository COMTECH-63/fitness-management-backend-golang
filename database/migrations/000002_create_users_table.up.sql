CREATE TYPE Sex AS ENUM ('male','female');

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR (55) NOT NULL,
  password VARCHAR (100) NOT NULL,
  first_name VARCHAR (30) NOT NULL,
  last_name VARCHAR (50) NOT NULL,
  id_card VARCHAR (13) UNIQUE NOT NULL,
  email VARCHAR (100) UNIQUE NOT NULL,
  phone_number VARCHAR (10) NOT NULL,
  address TEXT NOT NULL,
  sex Sex NOT NULL,
  image_url TEXT NOT NULL,
  member_id VARCHAR (5) NULL,
  created_at TIMESTAMP NULL,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL
);
-- comments
COMMENT ON COLUMN users.id IS 'The user ID';
COMMENT ON COLUMN users.first_name IS 'The user first name';
COMMENT ON COLUMN users.last_name IS 'The user last name';
COMMENT ON COLUMN users.id_card IS 'The user ID Card';
COMMENT ON COLUMN users.email IS 'The user email';
COMMENT ON COLUMN users.phone_number IS 'The user phone number';
COMMENT ON COLUMN users.address IS 'The user address';
COMMENT ON COLUMN users.sex IS 'The user sex';
COMMENT ON COLUMN users.image_url IS 'The user image profile url';
COMMENT ON COLUMN users.member_id IS 'The user Member ID';
COMMENT ON COLUMN users.created_at IS 'Create time';
COMMENT ON COLUMN users.updated_at IS 'Update time';
COMMENT ON COLUMN users.deleted_at IS 'Delete time';
