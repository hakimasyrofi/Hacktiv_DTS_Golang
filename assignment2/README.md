### SQL Query Create Table
```sh
-- Tabel untuk menyimpan pesanan
CREATE TABLE orders (
order_id SERIAL PRIMARY KEY,
customer_name VARCHAR(255),
ordered_at TIMESTAMP
);
```
```sh
-- Tabel untuk menyimpan detail item dalam pesanan
CREATE TABLE items (
item_id SERIAL PRIMARY KEY,
item_code VARCHAR(255),
description TEXT,
quantity INTEGER,
order_id INTEGER,
FOREIGN KEY (order_id) REFERENCES orders(order_id)
);
```

### Request Body Example

CreateOrder
```sh
{
"orderedAt": "2024-03-04T12:30:00Z",
"customerName": "Hakim",
"items": [
{
"itemCode": "ABC123",
"description": "Deskripsi item 1",
"quantity": 2
},
{
"itemCode": "DEF456",
"description": "Deskripsi item 2",
"quantity": 1
}
]
}
```

UpdateOrder
```sh
{
"orderedAt": "2024-03-04T12:30:00Z",
"customerName": "Hakim",
"items": [
{
"lineItemID": 1,
"itemCode": "ABC123",
"description": "Deskripsi item 1",
"quantity": 9
},
{
"lineItemID": 2,
"itemCode": "DEF459",
"description": "Deskripsi item 2",
"quantity": 22
}
]
}
```
