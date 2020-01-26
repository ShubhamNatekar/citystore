#! /bin/bash
mongorestore -d customers -c customers /backup/customers/customers/customers.bson
mongorestore -d shops -c shops /backup/shops/shops/shops.bson
mongorestore -d products -c products /backup/products/products/products.bson
mongorestore -d addresses -c addresses /backup/addresses/addresses/addresses.bson
