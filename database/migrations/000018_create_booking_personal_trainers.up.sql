-- create booking_personal_trainers table
CREATE TABLE IF NOT EXISTS booking_personal_trainers (
    id BIGSERIAL PRIMARY KEY,
    personal_trainer_id BIGSERIAL ,
    user_id BIGSERIAL ,
    date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (personal_trainer_id) REFERENCES personal_trainers (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to booking_personal_trainers table
COMMENT ON COLUMN booking_personal_trainers.id IS 'Booking Classes ID';
COMMENT ON COLUMN booking_personal_trainers.personal_trainer_id IS 'Classes ID';
COMMENT ON COLUMN booking_personal_trainers.user_id IS 'User ID';
COMMENT ON COLUMN booking_personal_trainers.date IS 'Booking Classes Date';
COMMENT ON COLUMN booking_personal_trainers.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN booking_personal_trainers.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN booking_personal_trainers.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
