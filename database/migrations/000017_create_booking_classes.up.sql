-- create booking_classes table
CREATE TABLE IF NOT EXISTS booking_classes (
    id BIGSERIAL PRIMARY KEY,
    class_id BIGSERIAL ,
    user_id BIGSERIAL ,
    date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (class_id) REFERENCES classes (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to booking_classes table
COMMENT ON COLUMN booking_classes.id IS 'Booking Classes ID';
COMMENT ON COLUMN booking_classes.class_id IS 'Classes ID';
COMMENT ON COLUMN booking_classes.user_id IS 'User ID';
COMMENT ON COLUMN booking_classes.date IS 'Booking Classes Date';
COMMENT ON COLUMN booking_classes.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN booking_classes.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN booking_classes.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
