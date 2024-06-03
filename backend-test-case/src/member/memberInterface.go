package member

import "backend_test_case/model/memberModel"

type MemberRepository interface {
	InsertMember(req memberModel.Member) error
	RetrieveMembers() ([]memberModel.Member, error)
	RetrieveMember(code string) (memberModel.Member, error)
	NameExist(name string) (bool, error)
	TableIsEmpty() (bool, error)
	GenerateNewCode(tableEmpty bool) (string, error)
	RetrieveMemberByName(name string) (memberModel.Member, error)
	MemberExist(code string) (bool, error)
	IncreaseBorrowedBooksTotal(code string) error
	PenalizedMember(code, startDate, endDate string) error
}

type MemberUsecase interface {
	Register(req memberModel.RegisterRequest) (memberModel.Member, error)
	Login(req memberModel.AuthRequest) (string, error)
	GetMembers() ([]memberModel.Member, error)
	GetMemberByCode(code string) (memberModel.Member, error)
}
