package tickets

import (
	"context"
	"desafio-goweb-julie-padilla/internal/domain"
)

// Desarrollar los métodos correspondientes a la estructura Ticket. Uno de ellos debe devolver la cantidad de tickets de un 
//destino. El  otro método debe devolver el promedio de personas que viajan a un país determinado en un dia:

// Paso 1. Se debe generar la interface Service con todos sus métodos.
type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByCountry(ctx context.Context, destination string) ([]domain.Ticket, error)
	//AverageDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

// Paso 2. Se debe generar la estructura service que contenga el repositorio.
type service struct {
	repository Repository
}

// Paso 3. Se debe generar una función que devuelva el Servicio.
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Paso 4. Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, GetTicketByDestination).
func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *service) GetTicketByCountry(ctx context.Context, destination string) ([]domain.Ticket, error){
	return s.repository.GetTicketByCountry(ctx, destination)
}

// func (s *service) AverageDestination(ctx context.Context, destination string) ([]domain.Ticket, error){
// 	return s.repository.AverageDestination(ctx, destination)
// }
