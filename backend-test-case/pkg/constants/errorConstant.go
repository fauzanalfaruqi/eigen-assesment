package constants

const (
	ErrSecretKeyNotSet          = "jwt secret key is not set"
	ErrIssuerNotSet             = "jwt issuer is not set"
	ErrAuthIsMissing            = "authorization header is missing"
	ErrInvalidIDType            = "invalid type for id, require number"
	ErrMemberDoesNotExist       = "member with requested id does not exist"
	ErrUsernameDoesNotExist     = "member with requested username does not exist"
	ErrUsernameAlreadyExist     = "member with requested username already exist"
	ErrMemberReachedMaxBorrow   = "member associated with this code has already borrowed the maximum allowable number of books"
	ErrPasswordDoesNotMatch     = "password does not match"
	ErrBookDoesNotExist         = "book with requested code does not exist"
	ErrBookNotAvailableToBorrow = "requested book is not available to borrow"
	ErrMemberIsPenalized        = "member associated with this code is penalized"
)
