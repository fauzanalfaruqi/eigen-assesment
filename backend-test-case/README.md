# Backend Test Case - Solution

The test case requirements can be found [here](https://github.com/eigen3dev/backend-test-case). The stacks used for this solution is using Golang for server side scripting, and PostgreSQL for the database.

## Get started

- #### Tables creation

    Create required table for this program to run properly with this sample ddl data, which also can be found [here](./ddl.sql):

    ```sql
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

- #### Sample mock data

    To populate the database with the initial data, you can copy provided dml sample data query, which also can be found [here](./dml.sql):

    ```sql
    