CREATE TABLE IF NOT EXISTS orders_table (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	customer_id INT NOT NULL,
	total BIGINT NOT NULL,
	status VARCHAR(25) NOT NULL
);