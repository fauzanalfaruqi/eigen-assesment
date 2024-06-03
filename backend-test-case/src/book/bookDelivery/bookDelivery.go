package bookDelivery

import (
	"backend_test_case/model/bookModel"
	"backend_test_case/model/dto/json"
	"backend_test_case/pkg/constants"
	"backend_test_case/src/book"

	"github.com/gin-gonic/gin"
)

type bookDelivery struct {
	bookUsecase book.BookUsecase
}

func NewBookDelivery(v1Group *gin.RouterGroup, bookUC book.BookUsecase) {
	handler := bookDelivery{bookUC}

	bookGroup := v1Group.Group("/books")
	{
		bookGroup.GET("", handler.getAvailableBooks)
		bookGroup.GET(":code", handler.getBookByCode)
		bookGroup.POST("", handler.postBook)
		bookGroup.POST(":code/borrow", handler.borrowBook)
		bookGroup.POST(":code/return", handler.returnBook)
	}
}

func (delivery *bookDelivery) postBook(c *gin.Context) {
	var req bookModel.BookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(c, err.Error(), constants.BookService, "01")
		return
	}

	err := delivery.bookUsecase.Insert(req)
	if err != nil {
		json.NewResponseBadRequest(c, []json.ValidationField{}, err.Error(), constants.BookService, "02")
		return
	}

	json.NewResponseCreated(c, nil, "data inserted successfully", constants.BookService, "00")
}

func (delivery *bookDelivery) getAvailableBooks(c *gin.Context) {
	books, err := delivery.bookUsecase.GetAvailableBooks()
	if err != nil {
		json.NewResponseError(c, err.Error(), constants.BookService, "01")
		return
	}

	json.NewResponseSuccess(c, books, "data received successfully", constants.BookService, "00")
}

func (delivery *bookDelivery) getBookByCode(c *gin.Context) {
	code := c.Param("code")

	book, err := delivery.bookUsecase.GetBookByCode(code)
	if err != nil {
		if err.Error() == constants.ErrBookDoesNotExist {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName: code, Message: constants.ErrBookDoesNotExist}}, "bad request", constants.BookService, "01")
			return
		}

		json.NewResponseError(c, err.Error(), constants.BookService, "02")
		return
	}

	json.NewResponseSuccess(c, book, "data received successfully", constants.BookService, "00")
}

func (delivery *bookDelivery) borrowBook(c *gin.Context) {
	code := c.Param("code")
	var req bookModel.BorrowedBooksLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(c, err.Error(), constants.BookService, "01")
		return
	}

	err := delivery.bookUsecase.BorrowBook(code, req)
	if err != nil {
		if err.Error() == constants.ErrBookNotAvailableToBorrow {
			json.NewResponseBadRequest(c, []json.ValidationField{}, constants.ErrBookNotAvailableToBorrow, constants.BookService, "02")
			return
		}

		if err.Error() == constants.ErrMemberDoesNotExist {
			json.NewResponseBadRequest(c, []json.ValidationField{}, constants.ErrMemberDoesNotExist, constants.BookService, "03")
			return
		}

		if err.Error() == constants.ErrMemberReachedMaxBorrow {
			json.NewResponseBadRequest(c, []json.ValidationField{}, constants.ErrMemberReachedMaxBorrow, constants.BookService, "04")
			return
		}

		json.NewResponseError(c, err.Error(), constants.BookService, "05")
		return
	}

	json.NewResponseSuccess(c, nil, "successfully borrow book", constants.BookService, "00")
}

func (delivery *bookDelivery) returnBook(c *gin.Context) {
	code := c.Param("code")
	var req bookModel.BorrowedBooksLogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(c, err.Error(), constants.BookService, "01")
		return
	}

	err := delivery.bookUsecase.ReturnBook(code, req)
	if err != nil {
		if err.Error() == constants.ErrWrongMemberToReturnBook {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName: req.MemberCode, Message: constants.ErrWrongMemberToReturnBook}}, "Bad request", constants.BookService, "02")
			return
		}

		json.NewResponseError(c, err.Error(), constants.BookService, "03")
		return
	}

	json.NewResponseSuccess(c, nil, "successfully borrow book", constants.BookService, "00")
}
