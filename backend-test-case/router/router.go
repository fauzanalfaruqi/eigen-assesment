package router

import (
	"backend_test_case/src/book/bookDelivery"
	"backend_test_case/src/book/bookRepository"
	"backend_test_case/src/book/bookUsecase"
	"backend_test_case/src/member/memberDelivery"
	"backend_test_case/src/member/memberRepository"
	"backend_test_case/src/member/memberUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	memberRepo := memberRepository.NewMemberRespository(db)
	memberUsecase := memberUsecase.NewMemberUsecase(memberRepo)
	memberDelivery.NewMemberDelivery(v1Group, memberUsecase)

	bookRepo := bookRepository.NewBookRepository(db)
	bookUsecase := bookUsecase.NewBookUsecase(bookRepo, memberRepo)
	bookDelivery.NewBookDelivery(v1Group, bookUsecase)
}
