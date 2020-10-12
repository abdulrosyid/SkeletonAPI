DROP TABLE IF EXISTS customer;

CREATE TABLE customer (
	customer_number int4 NOT NULL,
	"name" varchar NOT NULL
);

INSERT INTO customer (customer_number,"name") VALUES
(1001,'Bob Martin')
,(1002,'Linus Torvalds')
;