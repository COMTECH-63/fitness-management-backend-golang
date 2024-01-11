-- create orders table
CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL ,
    total NUMERIC (12,2) NOT NULL , 
    vat NUMERIC (12,2) NOT NULL ,
    total_vat NUMERIC (12,2) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- comment to orders table
COMMENT ON COLUMN orders.id IS 'Orders ID';
COMMENT ON COLUMN orders.user_id IS 'User ID';
COMMENT ON COLUMN orders.total IS 'Orders Total Price';
COMMENT ON COLUMN orders.vat IS 'Orders Vat Price';
COMMENT ON COLUMN orders.total_vat IS 'Orders Total Vat Price';
COMMENT ON COLUMN orders.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN orders.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN orders.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
