DROP TABLE IF EXISTS customer;

CREATE TABLE customer (
	customer_number int4 NOT NULL,
	"name" varchar NOT NULL
);

INSERT INTO customer (customer_number,"name") VALUES
(1001,'Bob Martin')
,(1002,'Linus Torvalds')
;

DROP TABLE IF EXISTS account;

CREATE TABLE account (
	account_number int4 NOT NULL,
	customer_number int4 NOT NULL,
	balance int8 NOT NULL
);

INSERT INTO account (account_number,customer_number,balance) VALUES
(555002,1002,10000)
,(555001,1001,5000)
;