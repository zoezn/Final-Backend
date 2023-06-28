package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/zoezn/Final-Backend/internal/domain"
)

type jsonStore struct {
	pathToFile string
}

// loadProducts carga los productos desde un archivo json
func (s *jsonStore) loadProducts() ([]domain.Dentista, error) {
	var products []domain.Dentista
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// saveProducts guarda los productos en un archivo json
func (s *jsonStore) saveProducts(products []domain.Dentista) error {
	bytes, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

// NewJsonStore crea un nuevo store de products
func NewJsonStore(path string) StoreInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (s *jsonStore) Read(id int) (domain.Dentista, error) {
	products, err := s.loadProducts()
	if err != nil {
		return domain.Dentista{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Dentista{}, errors.New("product not found")
}

func (s *jsonStore) Create(product domain.Dentista) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	product.Id = len(products) + 1
	products = append(products, product)
	return s.saveProducts(products)
}

func (s *jsonStore) Update(product domain.Dentista) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i] = product
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

func (s *jsonStore) Delete(id int) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == id {
			products = append(products[:i], products[i+1:]...)
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

func (s *jsonStore) Exists(codeValue string) bool {
	products, err := s.loadProducts()
	if err != nil {
		return false
	}
	for _, p := range products {
		if p.Matricula == codeValue {
			return true
		}
	}
	return false
}
