-- create role_has_users table
CREATE TABLE IF NOT EXISTS role_has_users (
    role_id BIGSERIAL,
    user_id BIGSERIAL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (role_id , user_id),
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to role_has_users table
COMMENT ON COLUMN role_has_users.role_id IS 'ID ของ Role';
COMMENT ON COLUMN role_has_users.user_id IS 'ID ของ User';
COMMENT ON COLUMN role_has_users.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN role_has_users.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN role_has_users.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
