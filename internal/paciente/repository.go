package paciente

import (
	"errors"
	"strconv"

	"github.com/zoezn/Final-Backend/internal/domain"
	"github.com/zoezn/Final-Backend/pkg/store"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domain.Paciente, error)
	// Create agrega un nuevo paciente
	Create(p domain.Paciente) (domain.Paciente, error)
	// Update actualiza un paciente
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage store.PacienteInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.PacienteInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Paciente, error) {
	paciente, err := r.storage.Read(id)
	if err != nil {
		return domain.Paciente{}, errors.New("paciente not found")
	}
	return paciente, nil

}

func (r *repository) Create(p domain.Paciente) (domain.Paciente, error) {
	if r.storage.Exists(strconv.Itoa(p.DNI)) {
		return domain.Paciente{}, errors.New("dni already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Paciente{}, errors.New("error creating paciente")
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	if r.storage.Exists(strconv.Itoa(p.DNI)) {
		return domain.Paciente{}, errors.New("dni already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		return domain.Paciente{}, errors.New("error updating paciente")
	}
	return p, nil
}
