package memberModel

type (
	Member struct {
		Code               string      `json:"code,omitempty"`
		Username           string      `json:"username,omitempty"`
		Password           string      `json:"-"`
		Role               string      `json:"role,omitempty"`
		TotalBooksBorrowed int         `json:"total_books_borrowed"`
		PenalizedStartDate string      `json:"penalized_start_date,omitempty"`
		PenalizedEndDate   string      `json:"penalized_end_date,omitempty"`
		CreatedAt          string      `json:"created_at,omitempty"`
		UpdatedAt          string      `json:"updated_at,omitempty"`
		DeletedAt          interface{} `json:"deleted_at,omitempty"`
	}

	AuthRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	RegisterRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	UpdateRequest struct {
		Code     string `json:"code"`
		Username string `json:"username"`
	}

	UpdatePasswordRequest struct {
		Code                 string `json:"code"`
		CurrentPassword      string `json:"current_password" validate:"required"`
		NewPassword          string `json:"new_password" validate:"required"`
		ConfirmationPassword string `json:"confirmation_password" validate:"required"`
	}
)
