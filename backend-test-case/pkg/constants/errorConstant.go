package constants

const (
	ErrSecretKeyNotSet          = "jwt secret key is not set"
	ErrIssuerNotSet             = "jwt issuer is not set"
	ErrAuthIsMissing            = "authorization header is missing"
	ErrInvalidIDType            = "invalid type for id, require number"
	ErrUsernameDoesNotExist     = "member with requested username does not exist"
	ErrUsernameAlreadyExist     = "member with requested username already exist"
	ErrPasswordDoesNotMatch     = "password does not match"
	ErrMemberDoesNotExist       = "member with requested id does not exist"
	ErrMemberReachedMaxBorrow   = "member associated with this code has already borrowed the maximum allowable number of books"
	ErrWrongMemberToReturnBook  = "wrong member to return the book associated with this code"
	ErrMemberIsPenalized        = "member associated with this code is penalized"
	ErrBookDoesNotExist         = "book with requested code does not exist"
	ErrBookNotAvailableToBorrow = "requested book is not available to borrow"
)
