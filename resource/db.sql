DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS customers;

CREATE TABLE customers (
    customer_id     int NOT NULL AUTO_INCREMENT,
    name            varchar(100) NOT NULL,
    date_of_birth   date NOT NULL,
    city            varchar(100) NOT NULL,
    postal_code     varchar(10) NOT NULL,
    status          smallint NOT NULL DEFAULT 1,
    PRIMARY KEY (customer_id)
);

CREATE TABLE accounts (
    account_id      int NOT NULL AUTO_INCREMENT,
    customer_id     int NOT NULL,
    opening_date    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    account_type    varchar(10) NOT NULL,
    amount          int NOT NULL,
    status          smallint NOT NULL DEFAULT 1,
    PRIMARY KEY (account_id),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);


CREATE TABLE transactions (
    transaction_id      int NOT NULL AUTO_INCREMENT,
    account_id          int NOT NULL,
    amount              int NOT NULL,
    transaction_type    varchar(10) NOT NULL,
    transaction_date    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (transaction_id),
    FOREIGN KEY (account_id) REFERENCES accounts(account_id)
);

CREATE TABLE users (
	user_id				int NOT NULL AUTO_INCREMENT,
    username            varchar(32) NOT NULL,
    password            varchar(100) NOT NULL,
    role                varchar(8) NOT NULL,
    customer_id         int,
    creation_date       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

