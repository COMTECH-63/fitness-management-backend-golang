-- create class_has_users table
CREATE TABLE IF NOT EXISTS class_has_users (
    class_id BIGSERIAL ,
    user_id BIGSERIAL ,
    date_start DATE NOT NULL , 
    date_end DATE NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (class_id, user_id),
    FOREIGN KEY (class_id) REFERENCES classes (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to class_has_users table
COMMENT ON COLUMN class_has_users.class_id IS 'Classes ID';
COMMENT ON COLUMN class_has_users.user_id IS 'User ID';
COMMENT ON COLUMN class_has_users.date_start IS 'วันเริ่มต้น class ที่ user จอง';
COMMENT ON COLUMN class_has_users.date_end IS 'วันสิ้นสุด class ที่ user จอง';
COMMENT ON COLUMN class_has_users.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN class_has_users.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN class_has_users.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
