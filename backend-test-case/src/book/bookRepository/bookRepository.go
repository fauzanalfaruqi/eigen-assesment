package bookRepository

import (
	"backend_test_case/model/bookModel"
	"backend_test_case/src/book"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) book.BookRepository {
	return &bookRepository{db}
}

func (repo *bookRepository) Insert(req bookModel.Book) error {
	query :=
		`INSERT INTO book (code, title, author, stock, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := repo.db.Exec(query, req.Code, req.Title, req.Author, req.Stock, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo *bookRepository) RetrieveAvailableBooks() ([]bookModel.Book, error) {
	minStock := 1

	query :=
		`SELECT code, title, author, stock, created_at, updated_at
		FROM book WHERE stock >= $1 and deleted_at IS NULL;`

	rows, err := repo.db.Query(query, minStock)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books, err := scanbooks(rows)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (repo *bookRepository) RetrieveBook(code string) (bookModel.Book, error) {
	query :=
		`SELECT code, title, author, stock, created_at, updated_at
		FROM book WHERE code = $1;`

	book, err := scanBook(repo.db.QueryRow(query, code))
	if err != nil {
		return bookModel.Book{}, err
	}

	return book, nil
}

func (repo *bookRepository) BookExist(code string) (bool, error) {
	var count int

	query :=
		`SELECT COUNT(*) FROM book WHERE code = $1;`

	err := repo.db.QueryRow(query, code).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *bookRepository) ReduceStock(code string) error {
	query :=
		`UPDATE book SET stock = stock - 1 WHERE code = $1;`

	_, err := repo.db.Exec(query, code)

	return err
}

func (repo *bookRepository) IncreaseStock(code string) error {
	query :=
		`UPDATE book SET stock = stock + 1 WHERE code = $1;`

	_, err := repo.db.Exec(query, code)

	return err
}

func (repo *bookRepository) AvailableToBorrow(code string) (bool, error) {
	var stock int

	query :=
		`SELECT stock FROM book WHERE code = $1 AND deleted_at IS NULL;`

	err := repo.db.QueryRow(query, code).Scan(&stock)

	return stock > 0, err
}

func (repo *bookRepository) BookTableIsEmpty() (bool, error) {
	count := 0

	query :=
		`SELECT COUNT(*) FROM book;`

	err := repo.db.QueryRow(query).Scan(&count)

	return count <= 0, err
}

// Borrowed book log

func (repo *bookRepository) InsertBorrowLog(req bookModel.BorrowedBooksLog) error {
	query :=
		`INSERT INTO borrowed_books_log (code, book_code, member_code, borrow_start_date, borrow_end_date, returned, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := repo.db.Exec(query, req.Code, req.BookCode, req.MemberCode, req.BorrowStartDate, req.BorrowEndDate, req.Returned, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo *bookRepository) RetrieveBorrowLogByBookCode(bookCode string) (bookModel.BorrowedBooksLog, error) {
	var log bookModel.BorrowedBooksLog
	query :=
		`SELECT code, book_code, member_code, borrow_start_date, borrow_end_date, returned, created_at, updated_at
		FROM borrowed_books_log WHERE book_code = $1 ORDER BY borrow_end_date DESC LIMIT 1;`

	err := repo.db.QueryRow(query, bookCode).Scan(
		&log.Code,
		&log.BookCode,
		&log.MemberCode,
		&log.BorrowStartDate,
		&log.BorrowEndDate,
		&log.Returned,
		&log.CreatedAt,
		&log.UpdatedAt,
	)

	if err != nil {
		return bookModel.BorrowedBooksLog{}, err
	}

	return log, nil

}
func (repo *bookRepository) LogTableIsEmpty() (bool, error) {
	count := 0

	query :=
		`SELECT COUNT(*) FROM borrowed_books_log;`

	err := repo.db.QueryRow(query).Scan(&count)

	return count <= 0, err
}

func (repo *bookRepository) GenerateNewLogCode(tableEmpty bool) (string, error) {
	var maxID string

	query :=
		`SELECT MAX(code) FROM borrowed_books_log;`

	if tableEmpty {
		return "BBL0001", nil
	}

	err := repo.db.QueryRow(query).Scan(&maxID)
	if err != nil {
		return "", errors.New("error gen 1")
	}

	num, err := strconv.Atoi(maxID[3:])
	if err != nil {
		return "", errors.New("error gen 2")
	}

	num++
	return fmt.Sprintf("BBL%04d", num), nil
}

//-------------------------------------------------------------------
// Repo layer utils
//-------------------------------------------------------------------

func scanBook(row *sql.Row) (bookModel.Book, error) {
	var book bookModel.Book
	err := row.Scan(
		&book.Code,
		&book.Title,
		&book.Author,
		&book.Stock,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		return bookModel.Book{}, err
	}

	return book, nil
}

func scanbooks(rows *sql.Rows) ([]bookModel.Book, error) {
	var books []bookModel.Book
	for rows.Next() {
		var book bookModel.Book
		err := rows.Scan(
			&book.Code,
			&book.Title,
			&book.Author,
			&book.Stock,
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
