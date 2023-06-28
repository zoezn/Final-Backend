package dentista

import (
	"github.com/zoezn/Final-Backend/internal/domain"
)

type Service interface {
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Dentista, error)
	// Create agrega un nuevo producto
	Create(p domain.Dentista) (domain.Dentista, error)
	// Delete elimina un producto
	Delete(id int) error
	// Update actualiza un producto
	Update(id int, p domain.Dentista) (domain.Dentista, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *service) Create(p domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Matricula != "" {
		p.Matricula = u.Matricula
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
