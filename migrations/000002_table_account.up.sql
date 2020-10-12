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