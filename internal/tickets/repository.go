package tickets

import (
	"context"
	"fmt"
	"desafio-goweb-julie-padilla/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByCountry(ctx context.Context, destination string) ([]domain.Ticket, error)
	//AverageDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByCountry(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

// func (r *repository) AverageDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

// 	var ticketsDest []domain.Ticket
// 	var sum = 0
// 	var avg = 0
	
// 	for index, value := range ticketsDest {
// 		sum += int(value.Time)
// 	}
// }
