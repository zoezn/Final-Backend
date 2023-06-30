package store

import (
	"github.com/zoezn/Final-Backend/internal/domain"
)

type StoreInterface interface {
	// Read devuelve un producto por su id
	Read(id int) (domain.Dentista, error)
	// Create agrega un nuevo producto
	Create(product domain.Dentista) error
	// Update actualiza un producto
	// Update(product domain.Dentista) error
	// Delete elimina un producto
	// Delete(id int) error
	// Exists verifica si un producto existe
	// Exists(codeValue string) bool
}

type PacienteInterface interface {
	// Read devuelve un producto por su id
	Read(id int) (domain.Paciente, error)
	// Create agrega un nuevo producto
	Create(product domain.Paciente) error
	// Update actualiza un producto
	Update(product domain.Paciente) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) bool
}
