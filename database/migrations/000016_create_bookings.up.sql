-- create bookings table
CREATE TABLE IF NOT EXISTS bookings (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL ,
    date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to bookings table
COMMENT ON COLUMN bookings.id IS 'Booking ID';
COMMENT ON COLUMN bookings.user_id IS 'User ID';
COMMENT ON COLUMN bookings.date IS 'Booking Date';
COMMENT ON COLUMN bookings.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN bookings.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN bookings.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
