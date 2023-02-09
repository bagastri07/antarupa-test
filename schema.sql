CREATE TABLE user_data
(
    user_id   INT PRIMARY KEY UNIQUE ,
    username  VARCHAR(20) NOT NULL,
    join_date TIMESTAMP   NULL
);

INSERT INTO user_data (user_id, username, join_date)
VALUES  (1, 'Kings885', '2022-07-22 00:00:00'),
        (2, 'Cool_Guy', '2019-03-07 11:00:40'),
        (3, 'Vicious_Hydra', '2021-04-11 13:04:06');

CREATE TABLE currencies
(
    currency_id   INT PRIMARY KEY,
    currency_name VARCHAR(10) NOT NULL,
    CONSTRAINT currencies_currency_id_uindex
        UNIQUE (currency_id)
);

INSERT INTO currencies (currency_id, currency_name)
VALUES  (1, 'Gold'),
        (2, 'Silver'),
        (3, 'Copper');

CREATE TABLE user_currency
(
    user_currency_id INT PRIMARY KEY,
    user_id          INT           NOT NULL,
    currency_type    INT DEFAULT 0 NULL,
    amount           INT           NULL,
    CONSTRAINT user_currency_user_currency_id_uindex
        UNIQUE (user_currency_id),
    CONSTRAINT user_currency___fk_currency_type
        FOREIGN KEY (currency_type) REFERENCES currencies (currency_id)
            ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT user_currency___fk_user_id
        FOREIGN KEY (user_id) REFERENCES user_data (user_id)
            ON UPDATE CASCADE ON DELETE CASCADE
);

INSERT INTO user_currency (user_currency_id, user_id, currency_type, amount)
VALUES  (1, 1, 1, 40),
        (2, 1, 2, 100),
        (3, 1, 3, 50000),
        (4, 2, 1, 3),
        (5, 2, 2, 50),
        (6, 2, 3, 1000),
        (7, 3, 1, 10),
        (8, 3, 2, 500),
        (9, 3, 3, 20050);

CREATE TABLE items
(
    item_id   INT PRIMARY KEY,
    item_name VARCHAR(20) NOT NULL,
    CONSTRAINT items_item_id_uindex
        UNIQUE (item_id)
);

INSERT INTO items (item_id, item_name)
VALUES  (1, 'healing potion'),
        (2, 'revival beads'),
        (3, 'invicible amulet'),
        (4, 'name changer');
       
create table shops (
	id int primary key,
	item_id int not null,
	max_owned int not null,
	price int not null,
	currency_type int not null,
	
	foreign key (item_id) references items (item_id)
		 ON UPDATE CASCADE ON DELETE cascade,
	foreign key (currency_type) references currencies (currency_id)
		ON UPDATE CASCADE ON DELETE cascade
)

insert into shops (id, item_id, max_owned, price, currency_type)
values 	(1, 1, 50, 100, 2),
		(2, 1, 10, 2, 1),
		(3, 1, 3, 3000, 3)
		
create table user_items (
	id bigint primary key,
	user_id int not null,
	item_id int not null,
	total int not null,
	foreign key (item_id) references items (item_id)
		 ON UPDATE CASCADE ON DELETE cascade,
	foreign key (user_id) references user_data (user_id)
		 ON UPDATE CASCADE ON DELETE cascade
	
)
