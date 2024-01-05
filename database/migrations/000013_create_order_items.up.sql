-- create order_items table
CREATE TABLE IF NOT EXISTS order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGSERIAL ,
    service_id BIGSERIAL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (service_id) REFERENCES services (id)
);
-- comment to order_items table
COMMENT ON COLUMN order_items.id IS 'Orders Itmes ID';
COMMENT ON COLUMN order_items.order_id IS 'Order ID';
COMMENT ON COLUMN order_items.service_id IS 'Service ID';
COMMENT ON COLUMN order_items.created_at IS 'บันทึกเวลาที่ข้อมูลถูกสร้าง';
COMMENT ON COLUMN order_items.updated_at IS 'บันทึกเวลาที่ข้อมูลถูกแก้ไข';
COMMENT ON COLUMN order_items.deleted_at IS 'บันทึกเวลาที่ข้อมูลถูกลบ';
