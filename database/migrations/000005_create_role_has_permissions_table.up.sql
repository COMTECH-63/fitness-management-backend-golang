-- create role_has_permissions table
CREATE TABLE IF NOT EXISTS role_has_permissions (
    role_id BIGSERIAL,
    permission_id BIGSERIAL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (permission_id) REFERENCES permissions (id)
);
-- comment to role_has_permissions table
COMMENT ON COLUMN role_has_permissions.role_id IS 'ID ของ Role';
COMMENT ON COLUMN role_has_permissions.permission_id IS 'ID ของ Permission';
COMMENT ON COLUMN role_has_permissions.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN role_has_permissions.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN role_has_permissions.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
