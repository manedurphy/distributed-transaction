#! /bin/bash

docker exec -i customers_db mysql -u customers_db_user -pcustomers_db_password customers < sql/create_customers_table.sql
docker exec -i orders_db mysql -u orders_db_user -porders_db_password orders < sql/create_orders_table.sql
docker exec -i payments_db mysql -u payments_db_user -ppayments_db_password payments < sql/create_payments_table.sql
