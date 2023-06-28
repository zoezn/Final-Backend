package domain

type Turno struct {
	Id          int      `json:"id"`
	Paciente    Paciente `json:"paciente" binding:"required"`
	Dentista    Dentista `json:"dentista" binding:"required"`
	Fecha       string   `json:"fecha" binding:"required"`
	Hora        string   `json:"hora" binding:"required"`
	Descripcion string   `json:"descripcion" binding:"optional"`
}
