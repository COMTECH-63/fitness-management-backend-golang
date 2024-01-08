-- create roles_has_users table
CREATE TABLE IF NOT EXISTS roles_has_users (
    user_id BIGSERIAL ,
    role_id BIGSERIAL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (user_id , role_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);
-- comment to roles_has_users table
COMMENT ON COLUMN roles_has_users.user_id IS 'ID ของ User';
COMMENT ON COLUMN roles_has_users.role_id IS 'ID ของ Role';
COMMENT ON COLUMN roles_has_users.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN roles_has_users.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN roles_has_users.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
