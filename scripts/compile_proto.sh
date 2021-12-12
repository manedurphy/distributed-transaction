#! /bin/bash

make -C api/customers/v1/proto regen-go
make -C api/orders/v1/proto regen-go
make -C api/payments/v1/proto regen-go
