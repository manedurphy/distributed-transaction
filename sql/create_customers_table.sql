CREATE TABLE IF NOT EXISTS customers_table (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	wallet INT UNSIGNED NOT NULL,
	password VARCHAR(100) NOT NULL
);