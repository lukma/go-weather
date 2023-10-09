package mocks

import (
	"context"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/google/uuid"
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

func (mock *mockNoteRepository) CreateNote(ctx context.Context, note domain.Note) error {
	args := mock.Called(ctx, note)
	return args.Error(0)
}

func (mock *mockNoteRepository) UpdateNote(ctx context.Context, note domain.Note) error {
	args := mock.Called(ctx, note)
	return args.Error(0)
}

func (mock *mockNoteRepository) DeleteNote(ctx context.Context, id uuid.UUID) error {
	args := mock.Called(ctx, id)
	return args.Error(0)
}
