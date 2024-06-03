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
    -- Insert members
    INSERT INTO member
        (code, username, password, role, total_books_borrowed, penalized_start_date, penalized_end_date, created_at, updated_at)
    VALUES
        ('M001', 'Angga', '$2a$10$d3zy1Y3kbAOJMva8Du4/Ve/OTDFhYVbR14TtL1O4Iz5JeYNbyOtUS', 'MEMBER', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), -- pass: 12345
        ('M002', 'Ferry', '$2a$10$MYXE49b3KZjc4jbCrFTUzOcb3SowEeuqy6ol06B.lniSLiMfN7HaG', 'MEMBER', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), -- pass: 54321
        ('M003', 'Putri', '$2a$10$/a4YNcei6iqXZRxyA6hgN.ngclNEyQzOXuS5B1KX1s0HdX3xNV7bO', 'MEMBER', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); -- pass: 010101

    -- Insert books
    INSERT INTO book
        (code, title, author, stock, created_at, updated_at)
    VALUES
        ('JK-45', 'Harry Potter', 'J.K Rowling', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
        ('SHR-1', 'A Study in Scarlet', 'Arthur Conan Doyle', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
        ('TW-11', 'Twilight', 'Stephenie Meyer', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
        ('HOB-83', 'The Hobbit, or There and Back Again', 'J.R.R. Tolkien', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
        ('NRN-7', 'The Lion, the Witch and the Wardrobe', 'C.S. Lewis', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

- #### Create `.env` file

    Create .env file in the root directory of this project. Copy and modify env settings below into your created .env file:

    ```.env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASS=
    DB_NAME=
    MAX_IDLE=1
    MAX_CONN=2
    MAX_LIFE_TIME=1h

    PORT=8081
    LOG_MODE=1

    JWT_SECRET_KEY=
    JWT_ISSUER=
    ```

- #### Run the program

    To run the program, simply execute this command:

    ```shell
    > go run .
    ```

## Endpoints

- #### Members

  | Method | Description                    | Endpoint                 |
  | ------ | ------------------------------ | ------------------------ |
  | GET    | Get all members                | /api/v1/members          |
  | GET    | Get member by given code param | /api/v1/members/{:code}  |
  | POST   | Register new member            | /api/v1/members/register |
  | POST   | Login member                   | /api/v1/members/login    |

- #### Books

	| Method | Description                   | Endpoint                     |
	| ------ | ----------------------------- | ---------------------------- |
	| GET    | Get available books to borrow | /api/v1/books                |
	| GET    | Get book by given code param  | /api/v1/books/{:code}        |
	| POST   | Post a new book               | /api/v1/books                |
	| POST   | Borrow a book                 | /api/v1/books/{:code}/borrow |
	| POST   | Return a book                 | /api/v1/books/{:code}/return |
	
	

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): Gin is a HTTP web framework written in Go (Golang).
- [GoDotEnv](https://github.com/joho/godotenv): A Go (golang) port of the Ruby [dotenv](https://github.com/bkeepers/dotenv) project (which loads env vars from a .env file).
- [pq](https://github.com/lib/pq): Pure Go Postgres driver for database/sql.
- [jwt-go](https://github.com/dgrijalva/jwt-go): Golang implementation of JSON Web Tokens (JWT).
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock): **sqlmock** is a mock library implementing [sql/driver](https://godoc.org/database/sql/driver).
- [validator](https://github.com/go-playground/validator/v10): Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.
- [crypto](https://golang.org/x/crypto): Go supplementary cryptography libraries.
- [zerolog](https://github.com/rs/zerolog): Zero allocation JSON logger.
- [cors](https://github.com/gin-contrib/cors): Official CORS gin's middleware.
- [logger](https://github.com/gin-contrib/logger): Gin middleware/handler to logger url path using rs/zerolog.