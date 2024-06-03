package bookModel

type (
	Book struct {
		Code      string `json:"code,omitempty"`
		Title     string `json:"title,omitempty"`
		Author    string `json:"author,omitempty"`
		Stock     int    `json:"stock,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		DeletedAt string `json:"deleted_at,omitempty"`
	}

	BookRequest struct {
		Code   string `json:"code" validate:"required"`
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
	}

	BorrowedBooksLog struct {
		Code            string `json:"code,omitempty"`
		BookCode        string `json:"book_code,omitempty"`
		MemberCode      string `json:"member_code,omitempty"`
		BorrowStartDate string `json:"borrow_start_date,omitempty"`
		BorrowEndDate   string `json:"borrow_end_date,omitempty"`
		Returned        bool   `json:"returned,omitempty"`
		CreatedAt       string `json:"created_at,omitempty"`
		UpdatedAt       string `json:"updated_at,omitempty"`
		DeletedAt       string `json:"deleted_at,omitempty"`
	}

	BorrowedBooksLogRequest struct {
		MemberCode string `json:"member_code" validate:"required"`
	}
)
