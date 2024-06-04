package memberRepository

import (
	"backend_test_case/model/memberModel"
	"backend_test_case/src/member"
	"database/sql"
	"fmt"
	"strconv"
)

type memberRepository struct {
	db *sql.DB
}

func NewMemberRespository(db *sql.DB) member.MemberRepository {
	return &memberRepository{db}
}

func (repo *memberRepository) InsertMember(req memberModel.Member) error {
	query :=
		`INSERT INTO member (code, username, password, role, total_books_borrowed, penalized_start_date, penalized_end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	_, err := repo.db.Exec(query, req.Code, req.Username, req.Password, req.Role, req.TotalBooksBorrowed, req.PenalizedStartDate, req.PenalizedEndDate, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo *memberRepository) RetrieveMembers() ([]memberModel.Member, error) {
	query :=
		`SELECT code, username, total_books_borrowed, penalized_start_date, penalized_end_date, created_at, updated_at FROM member;`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members, err := scanMembers(rows)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (repo *memberRepository) RetrieveMember(code string) (memberModel.Member, error) {
	query :=
		`SELECT code, username, password, total_books_borrowed, penalized_start_date, penalized_end_date, created_at, updated_at FROM member WHERE code = $1;`

	member, err := scanMember(repo.db.QueryRow(query, code))
	if err != nil {
		return memberModel.Member{}, err
	}

	return member, nil
}

func (repo *memberRepository) RetrieveMemberByName(name string) (memberModel.Member, error) {
	query :=
		`SELECT code, username, password, total_books_borrowed, penalized_start_date, penalized_end_date, created_at, updated_at FROM member WHERE username = $1;`

	member, err := scanMember(repo.db.QueryRow(query, name))
	if err != nil {
		return memberModel.Member{}, err
	}

	return member, nil
}

func (repo *memberRepository) NameExist(name string) (bool, error) {
	count := 0

	query :=
		`SELECT COUNT(*) FROM member WHERE username = $1;`

	err := repo.db.QueryRow(query, name).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *memberRepository) MemberExist(code string) (bool, error) {
	var count int

	query :=
		`SELECT COUNT(*) FROM member WHERE code = $1;`

	err := repo.db.QueryRow(query, code).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *memberRepository) IncreaseBorrowedBooksTotal(code string) error {
	query :=
		`UPDATE member SET total_books_borrowed = total_books_borrowed + 1 WHERE code = $1;`

	_, err := repo.db.Exec(query, code)

	return err
}

func (repo *memberRepository) DecreaseBorrowedBooksTotal(code string) error {
	query :=
		`UPDATE member SET total_books_borrowed = total_books_borrowed - 1 WHERE code = $1;`

	_, err := repo.db.Exec(query, code)

	return err
}

func (repo *memberRepository) TableIsEmpty() (bool, error) {
	count := 0

	query :=
		`SELECT COUNT(*) FROM member;`

	err := repo.db.QueryRow(query).Scan(&count)

	return count <= 0, err
}

func (repo *memberRepository) PenalizedMember(code, startDate, endDate string) error {
	query :=
		`UPDATE member SET penalized_start_date = $2, penalized_end_date = $3 WHERE code = $1;`

	_, err := repo.db.Exec(query, code, startDate, endDate)

	return err
}

func (repo *memberRepository) GenerateNewCode(tableEmpty bool) (string, error) {
	var maxID string

	query :=
		`SELECT MAX(code) FROM member;`

	if tableEmpty {
		return "M001", nil
	}

	err := repo.db.QueryRow(query).Scan(&maxID)
	if err != nil {
		return "", err
	}

	num, err := strconv.Atoi(maxID[1:])
	if err != nil {
		return "", err
	}

	num++
	return fmt.Sprintf("M%03d", num), nil
}

//-------------------------------------------------------------------
// Repo layer utils
//-------------------------------------------------------------------

func scanMember(row *sql.Row) (memberModel.Member, error) {
	var member memberModel.Member
	err := row.Scan(
		&member.Code,
		&member.Username,
		&member.Password,
		&member.TotalBooksBorrowed,
		&member.PenalizedStartDate,
		&member.PenalizedEndDate,
		&member.CreatedAt,
		&member.UpdatedAt,
	)

	if err != nil {
		return memberModel.Member{}, err
	}

	return member, nil
}

func scanMembers(rows *sql.Rows) ([]memberModel.Member, error) {
	var members []memberModel.Member
	for rows.Next() {
		var member memberModel.Member
		err := rows.Scan(
			&member.Code,
			&member.Username,
			&member.TotalBooksBorrowed,
			&member.PenalizedStartDate,
			&member.PenalizedEndDate,
			&member.CreatedAt,
			&member.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}
