-- create services table
CREATE TABLE IF NOT EXISTS services (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR (30) NOT NULL,
  description VARCHAR (50) NOT NULL,
  price NUMERIC (12,2) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);
-- comment to group_services table
COMMENT ON COLUMN services.id IS 'ID ของ Permission';
COMMENT ON COLUMN services.name IS 'ชื่อของบริการ';
COMMENT ON COLUMN services.description IS 'คำอธิบาย';
COMMENT ON COLUMN services.price IS 'ราคาบริการ';
COMMENT ON COLUMN services.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN services.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN services.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';