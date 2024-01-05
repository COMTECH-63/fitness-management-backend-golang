-- create service_has_users table
CREATE TABLE IF NOT EXISTS service_has_users (
    service_id BIGSERIAL,
    user_id BIGSERIAL ,
    date_start DATE , 
    date_end DATE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (service_id , user_id),
    FOREIGN KEY (service_id) REFERENCES services (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to service_has_users table
COMMENT ON COLUMN service_has_users.service_id IS 'ID ของ Service';
COMMENT ON COLUMN service_has_users.user_id IS 'ID ของ User';
COMMENT ON COLUMN service_has_users.date_start IS 'วันเริ่มต้น service';
COMMENT ON COLUMN service_has_users.date_end IS 'วันสิ้นสุดservice';
COMMENT ON COLUMN service_has_users.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN service_has_users.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN service_has_users.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
