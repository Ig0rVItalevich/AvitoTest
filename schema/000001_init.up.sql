create table users
(
	id      serial primary key,
	balance numeric default 0 check ( balance >= 0 ) not null
);

create table transactions
(
	id           serial primary key,
	user_id      int references users (id)      not null,
	amount       numeric                        not null,
	order_id     int         default 0          not null,
	product_id   int         default 0          not null,
	status       varchar(50) default 'reserved' not null,
	created_date timestamp   default now()      not null
);

create table history
(
	id            serial primary key,
	user_from     int references users (id) not null,
	user_to       int references users (id) not null,
	amount        numeric                   not null,
	order_id      int       default 0       not null,
	product_id    int       default 0       not null,
	description   varchar(50)               not null,
	accepted_date timestamp default now()   not null
);

INSERT INTO users (balance)
VALUES ('0');