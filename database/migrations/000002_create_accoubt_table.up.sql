-- create accounts table
CREATE TABLE IF NOT EXISTS accounts (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL ,
  username VARCHAR (55) NOT NULL,
  password VARCHAR (50) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to accounts table
COMMENT ON COLUMN accounts.id IS 'ID ของ Permission';
COMMENT ON COLUMN accounts.username IS 'Username ของผู้ใช้';
COMMENT ON COLUMN accounts.password IS 'Password ของผู้ใช้';
COMMENT ON COLUMN accounts.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN accounts.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN accounts.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';