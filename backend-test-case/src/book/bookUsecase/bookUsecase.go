package bookUsecase

import (
	"backend_test_case/model/bookModel"
	"backend_test_case/pkg/constants"
	"backend_test_case/src/book"
	"backend_test_case/src/member"
	"errors"
	"time"
)

type bookUsecase struct {
	bookRepo   book.BookRepository
	memberRepo member.MemberRepository
}

func NewBookUsecase(bookRepo book.BookRepository, memberRepo member.MemberRepository) book.BookUsecase {
	return &bookUsecase{bookRepo, memberRepo}
}

func (usecase *bookUsecase) Insert(req bookModel.BookRequest) error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	var book = bookModel.Book{
		Code:      req.Code,
		Title:     req.Title,
		Author:    req.Author,
		Stock:     1,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	err := usecase.bookRepo.Insert(book)
	return err
}

func (usecase *bookUsecase) GetAvailableBooks() ([]bookModel.Book, error) {
	books, err := usecase.bookRepo.RetrieveAvailableBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (usecase *bookUsecase) GetBookByCode(code string) (bookModel.Book, error) {
	bookExist, err := usecase.bookRepo.BookExist(code)
	if err != nil {
		return bookModel.Book{}, err
	}

	if !bookExist {
		return bookModel.Book{}, errors.New(constants.ErrBookDoesNotExist)
	}

	book, err := usecase.bookRepo.RetrieveBook(code)
	if err != nil {
		return bookModel.Book{}, err
	}

	return book, nil
}

func (usecase *bookUsecase) BorrowBook(bookCode string, logReq bookModel.BorrowedBooksLogRequest) error {
	// Reduce book stock
	availableToBorrow, err := usecase.bookRepo.AvailableToBorrow(bookCode)
	if err != nil {
		return err
	}

	if !availableToBorrow {
		return errors.New(constants.ErrBookNotAvailableToBorrow)
	}

	err = usecase.bookRepo.ReduceStock(bookCode)
	if err != nil {
		return err
	}

	// Check if member is exist
	memberExist, err := usecase.memberRepo.MemberExist(logReq.MemberCode)
	if err != nil {
		return err
	}

	if !memberExist {
		return errors.New(constants.ErrMemberDoesNotExist)
	}

	// Get member data
	member, err := usecase.memberRepo.RetrieveMember(logReq.MemberCode)
	if err != nil {
		return err
	}

	// Check if member already reach max allowed borrow books total
	if member.TotalBooksBorrowed >= 2 {
		return errors.New(constants.ErrMemberReachedMaxBorrow)
	}

	// Check if PenalizedEndDate is not empty
	if member.PenalizedEndDate != "" {
		// Parse PenalizedEndDate
		penalizedEndDate, err := time.Parse("2006-01-02T15:04:05Z", member.PenalizedEndDate)
		if err != nil {
			return err
		}

		// Check if penalizedEndDate is in the future
		if penalizedEndDate.After(time.Now()) {
			return errors.New(constants.ErrMemberIsPenalized)
		}
	}

	// Increase borrowed_books_total value
	err = usecase.memberRepo.IncreaseBorrowedBooksTotal(logReq.MemberCode)
	if err != nil {
		return err
	}

	// Generate borrow log
	tableEmpty, err := usecase.bookRepo.TableIsEmpty()
	if err != nil {
		return err
	}

	logCode, err := usecase.bookRepo.GenerateNewLogCode(tableEmpty)
	if err != nil {
		return err
	}

	now := time.Now()
	nowFormated := now.Format("2006-01-02 15:04:05")
	borrowEndDate := now.AddDate(0, 0, 7).Format("2006-01-02 15:04:05")

	var log = bookModel.BorrowedBooksLog{
		Code:            logCode,
		BookCode:        bookCode,
		MemberCode:      logReq.MemberCode,
		BorrowStartDate: nowFormated,
		BorrowEndDate:   borrowEndDate,
		Returned:        false,
		CreatedAt:       nowFormated,
		UpdatedAt:       nowFormated,
	}

	err = usecase.bookRepo.InsertBorrowLog(log)
	if err != nil {
		return err
	}

	return nil
}
