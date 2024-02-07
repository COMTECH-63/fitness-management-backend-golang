CREATE TYPE Status AS ENUM ( 'not_paid','paid');

-- create order_payments table
CREATE TABLE IF NOT EXISTS order_payments (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGSERIAL ,
    amount NUMERIC (12,2) NOT NULL, 
    status Status NOT NULL, 
    date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id)
);
-- comment to order_payments table
COMMENT ON COLUMN order_payments.id IS 'Orders Payments ID';
COMMENT ON COLUMN order_payments.order_id IS 'Order ID';
COMMENT ON COLUMN order_payments.amount IS 'Orders Payments Amount';
COMMENT ON COLUMN order_payments.status IS 'Orders Payments Status';
COMMENT ON COLUMN order_payments.date IS 'Orders Payments Date';
COMMENT ON COLUMN order_payments.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN order_payments.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN order_payments.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
