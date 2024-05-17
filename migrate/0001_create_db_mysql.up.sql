CREATE TABLE IF NOT EXISTS tb_product (
    id INT AUTO_INCREMENT PRIMARY KEY,
    external_id CHAR(36) NOT NULL,
    name VARCHAR(120) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

DELIMITER $$
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON tb_product
FOR EACH ROW
BEGIN
    SET NEW.updated_at = CURRENT_TIMESTAMP;
END;
$$
DELIMITER ;

INSERT INTO tb_product (external_id, name, quantity, price) VALUES 
    (UUID(), 'Notebook 13 XPS', 10, 14500.00),
    (UUID(), 'Notebook 15', 15, 2500.00),
    (UUID(), 'Notebook 14', 20, 2500.00),
    (UUID(), 'Notebook 15', 12, 3500.00),
    (UUID(), 'Notebook 13', 8, 4568.00),
    (UUID(), 'Tablet', 5, 8450.00),
    (UUID(), 'Macbook 13 Pro M1', 30, 18500.00),
    (UUID(), 'TV 55', 25, 4500.00),
    (UUID(), 'TV 45', 18, 3500.00),
    (UUID(), 'TV 32', 22, 2500.00),
    (UUID(), 'TV 60', 10, 6500.00),
    (UUID(), 'TV 50', 15, 4800.00);
