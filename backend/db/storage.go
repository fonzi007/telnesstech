package db

import (
	"database/sql"
	"fmt"
	"telnesstech/model"
)

type Storage struct {
	db             *sql.DB
	preparedSelect *sql.Stmt
}

func NewStorage(client *sql.DB) (*Storage, error) {
	preparedSelect, err := client.Prepare("SELECT msisdn, grade, type, reserved_at, expires_at FROM phone_numbers WHERE (msisdn LIKE CONCAT(?, '%') AND msisdn LIKE CONCAT('%', ?) AND msisdn LIKE CONCAT('%', ?, '%')) AND (grade LIKE ? OR ? LIKE '') AND (type LIKE ? OR ? LIKE '')")
	if err != nil {
		fmt.Errorf("failed to prepare insert query: %w", err)
	}

	return &Storage{
		db:             client,
		preparedSelect: preparedSelect,
	}, err
}

func (s *Storage) FindBy(request model.PhoneNumberSearchRequest) ([]model.PhoneNumber, error) {
	rows, err := s.preparedSelect.Query(
		request.Prefix,
		request.Suffix,
		request.Substring,
		request.FilterBy.Grade, request.FilterBy.Grade,
		request.FilterBy.Type, request.FilterBy.Type)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	phoneNumbers := make([]model.PhoneNumber, 0)
	for rows.Next() {
		var phoneNumber model.PhoneNumber
		rows.Scan(&phoneNumber.Msisdn, &phoneNumber.Grade, &phoneNumber.Type, &phoneNumber.ReservedAt, &phoneNumber.ExpiresAt)
		phoneNumbers = append(phoneNumbers, phoneNumber)
	}
	return phoneNumbers, nil
}
