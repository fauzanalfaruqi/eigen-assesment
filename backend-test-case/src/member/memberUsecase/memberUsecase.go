package memberUsecase

import (
	"backend_test_case/model/memberModel"
	"backend_test_case/pkg/constants"
	"backend_test_case/pkg/utils"
	"backend_test_case/src/member"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type memberUsecase struct {
	memberRepo member.MemberRepository
}

func NewMemberUsecase(memberRepo member.MemberRepository) member.MemberUsecase {
	return &memberUsecase{memberRepo}
}

func (usecase *memberUsecase) Register(req memberModel.RegisterRequest) (memberModel.Member, error) {
	nameExist, err := usecase.memberRepo.NameExist(req.Username)
	if err != nil {
		return memberModel.Member{}, err
	}

	if nameExist {
		return memberModel.Member{}, errors.New(constants.ErrUsernameAlreadyExist)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return memberModel.Member{}, err
	}

	tableEmpty, err := usecase.memberRepo.TableIsEmpty()
	if err != nil {
		return memberModel.Member{}, err
	}

	newCode, err := usecase.memberRepo.GenerateNewCode(tableEmpty)
	if err != nil {
		return memberModel.Member{}, err
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	timeZero := time.Time{}.Format("2006-01-02 15:04:05")

	var member = memberModel.Member{
		Code:               newCode,
		Username:           req.Username,
		Password:           string(hashPassword),
		Role:               "MEMBER",
		TotalBooksBorrowed: 0,
		PenalizedStartDate: timeZero,
		PenalizedEndDate:   timeZero,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
		DeletedAt:          timeZero,
	}

	err = usecase.memberRepo.InsertMember(member)
	if err != nil {
		return memberModel.Member{}, err
	}

	return member, nil
}

func (usecase *memberUsecase) Login(req memberModel.AuthRequest) (string, error) {
	nameExist, err := usecase.memberRepo.NameExist(req.Username)
	if err != nil {
		return "", err
	}

	if !nameExist {
		return "", errors.New(constants.ErrUsernameDoesNotExist)
	}

	member, err := usecase.memberRepo.RetrieveMemberByName(req.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(member.Code, member.Username, member.Role)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (usecase *memberUsecase) GetMembers() ([]memberModel.Member, error) {
	members, err := usecase.memberRepo.RetrieveMembers()
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (usecase *memberUsecase) GetMemberByCode(code string) (memberModel.Member, error) {
	memberExist, err := usecase.memberRepo.MemberExist(code)
	if err != nil {
		return memberModel.Member{}, err
	}

	if !memberExist {
		return memberModel.Member{}, errors.New(constants.ErrMemberDoesNotExist)
	}

	member, err := usecase.memberRepo.RetrieveMember(code)
	if err != nil {
		return memberModel.Member{}, err
	}

	return member, nil
}
