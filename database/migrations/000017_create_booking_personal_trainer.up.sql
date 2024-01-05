-- create booking_personal_trainer table
CREATE TABLE IF NOT EXISTS booking_personal_trainer (
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
-- comment to booking_personal_trainer table
COMMENT ON COLUMN booking_personal_trainer.id IS 'Booking Classes ID';
COMMENT ON COLUMN booking_personal_trainer.personal_trainer_id IS 'Classes ID';
COMMENT ON COLUMN booking_personal_trainer.user_id IS 'User ID';
COMMENT ON COLUMN booking_personal_trainer.date IS 'Booking Classes Date';
COMMENT ON COLUMN booking_personal_trainer.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN booking_personal_trainer.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN booking_personal_trainer.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
