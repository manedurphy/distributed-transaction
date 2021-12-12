CREATE TABLE IF NOT EXISTS payments_table (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	credit_card_number VARCHAR(16) NOT NULL,
	expiration DATE NOT NULL,
	cvc INT NOT NULL,
	customer_id INT NOT NULL
);