CREATE TABLE member (
	code varchar(100) NOT NULL PRIMARY KEY,
	username varchar(100),
    password varchar(255),
	role varchar(100),
	total_books_borrowed int,
	penalized_start_date timestamp,
	penalized_end_date timestamp,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);

CREATE TABLE book (
	code varchar(100) NOT NULL PRIMARY KEY,
	title varchar(255),
	author varchar(100),
	stock int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);

CREATE TABLE borrowed_books_log (
	code varchar(100) NOT NULL PRIMARY KEY,
	book_code varchar(100) NOT NULL REFERENCES book(code),
	member_code varchar(100) NOT NULL REFERENCES member(code),
	borrow_start_date timestamp,
	borrow_end_date timestamp,
	returned bool,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);

