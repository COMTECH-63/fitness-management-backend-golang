CREATE TYPE MotivationsP AS ENUM ( 'lean','performance','strong','well_being');
CREATE TYPE MinuteP AS ENUM ('30','45','60','95');

-- create personal_trainers table
CREATE TABLE IF NOT EXISTS personal_trainers (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL ,
    description VARCHAR (50) NOT NULL , 
    motivations MotivationsP NOT NULL , 
    minute MinuteP NOT NULL , 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to personal_trainers table
COMMENT ON COLUMN personal_trainers.id IS 'Personal Trainers ID';
COMMENT ON COLUMN personal_trainers.user_id IS 'User ID';
COMMENT ON COLUMN personal_trainers.description IS 'Personal Trainers Description';
COMMENT ON COLUMN personal_trainers.motivations IS 'Personal Trainers Motivations';
COMMENT ON COLUMN personal_trainers.minute IS 'Personal Trainers Minute';
COMMENT ON COLUMN personal_trainers.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN personal_trainers.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN personal_trainers.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
