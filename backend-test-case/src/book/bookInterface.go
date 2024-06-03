package book

import "backend_test_case/model/bookModel"

type BookRepository interface {
	Insert(req bookModel.Book) error
	RetrieveAvailableBooks() ([]bookModel.Book, error)
	RetrieveBook(code string) (bookModel.Book, error)
	BookExist(code string) (bool, error)
	ReduceStock(code string) error
	IncreaseStock(code string) error
	AvailableToBorrow(code string) (bool, error)
	InsertBorrowLog(req bookModel.BorrowedBooksLog) error
	BookTableIsEmpty() (bool, error)
	LogTableIsEmpty() (bool, error)
	GenerateNewLogCode(tableEmpty bool) (string, error)
	RetrieveBorrowLogByBookCode(bookCode string) (bookModel.BorrowedBooksLog, error)
}

type BookUsecase interface {
	Insert(req bookModel.BookRequest) error
	GetAvailableBooks() ([]bookModel.Book, error)
	GetBookByCode(code string) (bookModel.Book, error)
	BorrowBook(bookCode string, logReq bookModel.BorrowedBooksLogRequest) error
	ReturnBook(bookCode string, logReq bookModel.BorrowedBooksLogRequest) error
}
