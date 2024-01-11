-- create roles table
CREATE TABLE IF NOT EXISTS roles (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR (10) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);
-- comment to roles table
COMMENT ON COLUMN roles.id IS 'ID ของ Role';
COMMENT ON COLUMN roles.name IS 'ชื่อ Role';
COMMENT ON COLUMN roles.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN roles.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN roles.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
