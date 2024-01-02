-- create user_has_roles table
CREATE TABLE IF NOT EXISTS user_has_roles (
    user_id BIGSERIAL ,
    role_id BIGSERIAL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (user_id , role_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);
-- comment to user_has_roles table
COMMENT ON COLUMN user_has_roles.user_id IS 'ID ของ User';
COMMENT ON COLUMN user_has_roles.role_id IS 'ID ของ Role';
COMMENT ON COLUMN user_has_roles.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN user_has_roles.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN user_has_roles.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
