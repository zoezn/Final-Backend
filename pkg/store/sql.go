package store

import (
	"database/sql"

	"github.com/zoezn/Final-Backend/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Read(id int) (*domain.Dentista, error) {
	var productReturn domain.Dentista

	query := "SELECT * FROM dentistas WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&productReturn.Id, &productReturn.Nombre, &productReturn.Apellido, &productReturn.Matricula)
	if err != nil {
		return nil, err
	}
	return &productReturn, nil
}

func (s *SqlStore) Create(product domain.Dentista) (*domain.Dentista, error) {
	query := "INSERT INTO dentistas (Apellido, Nombre, Matricula) VALUES (?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(product.Apellido, product.Nombre, product.Matricula)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lid, _ := res.LastInsertId()
	product.Id = int(lid)
	return &product, nil
}

// func (s *SqlStore) Exist(codeValue string) bool {
// 	var exist bool
// 	var id int

// 	query := "SELECT id FROM dentistas WHERE code_value = ?;"
// 	row := s.DB.QueryRow(query, id)
// 	err := row.Scan(&id)
// 	if err != nil {
// 		return exist
// 	}

// 	if id > 0 {
// 		exist = true
// 	}

// 	return exist

// }
