package mocks

import (
	"context"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/stretchr/testify/mock"
)

type mockNoteRepository struct {
	mock.Mock
}

func NewNoteRepository() *mockNoteRepository {
	return new(mockNoteRepository)
}

func (mock *mockNoteRepository) GetNotes(ctx context.Context) ([]domain.Note, error) {
	args := mock.Called(ctx)
	return args.Get(0).([]domain.Note), args.Error(1)
}
