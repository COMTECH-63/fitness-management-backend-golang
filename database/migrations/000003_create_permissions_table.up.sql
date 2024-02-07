-- create permissions table
CREATE TABLE IF NOT EXISTS permissions (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR (25) NOT NULL,
  description VARCHAR (50) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);
-- comment to group_permissions table
COMMENT ON COLUMN permissions.id IS 'ID ของ Permission';
COMMENT ON COLUMN permissions.description IS 'คำอธิบาย';
COMMENT ON COLUMN permissions.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN permissions.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN permissions.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';