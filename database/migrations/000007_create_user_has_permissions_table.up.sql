-- create user_has_permissions table
CREATE TABLE IF NOT EXISTS user_has_permissions (
    user_id BIGSERIAL,
    permission_id BIGSERIAL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (user_id, permission_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (permission_id) REFERENCES permissions (id)
);
-- comment to user_has_permissions table
COMMENT ON COLUMN user_has_permissions.user_id IS 'ID ของ User';
COMMENT ON COLUMN user_has_permissions.permission_id IS 'ID ของ Permission';
COMMENT ON COLUMN user_has_permissions.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN user_has_permissions.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN user_has_permissions.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
