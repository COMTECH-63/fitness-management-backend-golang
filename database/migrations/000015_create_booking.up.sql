CREATE TYPE Status AS ENUM ( 'not_paid','paid');
-- create booking table
CREATE TABLE IF NOT EXISTS booking (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL ,
    date DATETIME NOT NULL,
    status BOOLEAN NOT NULL , 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to booking table
COMMENT ON COLUMN booking.id IS 'Booking ID';
COMMENT ON COLUMN booking.user_id IS 'User ID';
COMMENT ON COLUMN booking.date IS 'Booking Date';
COMMENT ON COLUMN booking.status IS 'Booking Status';
COMMENT ON COLUMN booking.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN booking.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN booking.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
