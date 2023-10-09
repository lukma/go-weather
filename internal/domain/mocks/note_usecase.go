package mocks

import (
	"context"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type mockNoteUsecase struct {
	mock.Mock
}

func NewNoteUsecase() *mockNoteUsecase {
	return new(mockNoteUsecase)
}

func (mock *mockNoteUsecase) GetNotes(ctx context.Context) ([]domain.Note, error) {
	args := mock.Called(ctx)
	return args.Get(0).([]domain.Note), args.Error(1)
}

func (mock *mockNoteUsecase) CreateNote(ctx context.Context, note domain.Note) error {
	args := mock.Called(ctx, note)
	return args.Error(0)
}

func (mock *mockNoteUsecase) UpdateNote(ctx context.Context, note domain.Note) error {
	args := mock.Called(ctx, note)
	return args.Error(0)
}

func (mock *mockNoteUsecase) DeleteNote(ctx context.Context, id uuid.UUID) error {
	args := mock.Called(ctx, id)
	return args.Error(0)
}
