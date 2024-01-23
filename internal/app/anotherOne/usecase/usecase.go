package anotherOne

import (
	anotherOne "app-module/internal/app/anotherOne/domain"
)

// `usecase` instance
type usecase struct {
	// Define dependecies
	r anotherOne.Repository
}

// New `usecase` factory
func New(r anotherOne.Repository) anotherOne.Usecase {
	// Inject dependencies
	return &usecase{
		r: r,
	}
}
