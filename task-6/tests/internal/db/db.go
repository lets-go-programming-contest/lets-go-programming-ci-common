package db

import (
	"database/sql"
)

type Database interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type DBService struct {
	DB Database
}

func New(db Database) DBService {
	return DBService{DB: db}
}

func (service DBService) GetNames() ([]string, error) {
	query := "SELECT name FROM users"

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return names, nil
}

func (service DBService) GetUniqueNames() ([]string, error) {
	query := "SELECT DISTINCT name FROM users"

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []string

	for rows.Next() {
		var value string

		if err := rows.Scan(&value); err != nil {
			return nil, err
		}

		values = append(values, value)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return values, nil
}
