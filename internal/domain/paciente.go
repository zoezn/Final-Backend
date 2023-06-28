package domain

type Paciente struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Domicilio string `json:"domicilio" binding:"required"`
	DNI       int    `json:"DNI" binding:"required"`
	Alta      string `json:"alta" binding:"required"`
}
