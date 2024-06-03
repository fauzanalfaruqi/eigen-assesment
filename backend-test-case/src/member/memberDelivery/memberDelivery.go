package memberDelivery

import (
	"backend_test_case/model/dto/json"
	"backend_test_case/model/memberModel"
	"backend_test_case/pkg/constants"
	"backend_test_case/src/member"

	"github.com/gin-gonic/gin"
)

type memberDelivery struct {
	memberUsecase member.MemberUsecase
}

func NewMemberDelivery(v1Group *gin.RouterGroup, memberUC member.MemberUsecase) {
	handler := memberDelivery{memberUC}

	memberGroup := v1Group.Group("/members")
	{
		memberGroup.POST("/register", handler.register)
		memberGroup.POST("/login", handler.login)
		memberGroup.GET("", handler.getMembers)
		memberGroup.GET(":code", handler.getMemberByCode)
	}
}

func (delivery *memberDelivery) register(c *gin.Context) {
	var req memberModel.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(c, err.Error(), constants.MemberService, "01")
		return
	}

	member, err := delivery.memberUsecase.Register(req)
	if err != nil {
		if err.Error() == constants.ErrUsernameAlreadyExist {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName: "name", Message: constants.ErrUsernameAlreadyExist}}, "Bad request", constants.MemberService, "02")
			return
		}

		json.NewResponseError(c, err.Error(), constants.MemberService, "03")
		return
	}

	json.NewResponseCreated(c, member, "member registered successfully", constants.MemberService, "00")
}

func (delivery *memberDelivery) login(c *gin.Context) {
	var req memberModel.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(c, err.Error(), constants.MemberService, "01")
		return
	}

	token, err := delivery.memberUsecase.Login(req)
	if err != nil {
		if err.Error() == constants.ErrUsernameDoesNotExist {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName: "name", Message: constants.ErrMemberDoesNotExist}}, "Bad request", constants.MemberService, "02")
			return
		}

		json.NewResponseError(c, err.Error(), constants.MemberService, "03")
		return
	}

	json.NewResponseSuccess(c, token, "login successfully", constants.MemberService, "00")
}

func (delivery *memberDelivery) getMembers(c *gin.Context) {
	members, err := delivery.memberUsecase.GetMembers()
	if err != nil {
		json.NewResponseError(c, err.Error(), constants.MemberService, "01")
		return
	}

	json.NewResponseSuccess(c, members, "users retrieved successfully", constants.MemberService, "00")
}

func (delivery *memberDelivery) getMemberByCode(c *gin.Context) {
	code := c.Param("code")

	member, err := delivery.memberUsecase.GetMemberByCode(code)
	if err != nil {
		if err.Error() == constants.ErrMemberDoesNotExist {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName: code, Message: constants.ErrMemberDoesNotExist}}, "bad request", constants.MemberService, "01")
			return
		}

		json.NewResponseError(c, err.Error(), constants.MemberService, "02")
		return
	}

	json.NewResponseSuccess(c, member, "data received successfully", constants.MemberService, "00")
}
