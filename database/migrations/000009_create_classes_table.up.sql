CREATE TYPE Motivations AS ENUM ( 'lean','performance','strong','well_being');
CREATE TYPE Intensity AS ENUM ( 'low','mid','high');
CREATE TYPE Minute AS ENUM ( 30,45,60,95);
-- create classes table
CREATE TABLE IF NOT EXISTS classes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL ,
    name VARCHAR (30) NOT NULL , 
    description VARCHAR (50) NOT NULL , 
    motivations Motivations NOT NULL , 
    intensity Intensity NOT NULL , 
    minute Minute NOT NULL , 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
);
-- comment to classes table
COMMENT ON COLUMN classes.id IS 'Classes ID';
COMMENT ON COLUMN classes.user_id IS 'User ID';
COMMENT ON COLUMN classes.name IS 'Classes Name';
COMMENT ON COLUMN classes.motivations IS 'Classes Motivations';
COMMENT ON COLUMN classes.intensity IS 'Classes Intensity';
COMMENT ON COLUMN classes.minute IS 'Classes Minute';
COMMENT ON COLUMN classes.description IS 'Classes Description';
COMMENT ON COLUMN classes.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN classes.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN classes.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
