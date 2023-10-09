package usecase

import (
	"context"
	"time"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/google/uuid"
)

type noteUsecase struct {
	contextTimeout time.Duration
	noteRepository domain.NoteRepository
}

func NewNoteUsecase(timeout time.Duration, noteRepository domain.NoteRepository) domain.NoteUsecase {
	return &noteUsecase{
		contextTimeout: timeout,
		noteRepository: noteRepository,
	}
}

func (usecase *noteUsecase) GetNotes(c context.Context) ([]domain.Note, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.noteRepository.GetNotes(ctx)
}

func (usecase *noteUsecase) CreateNote(c context.Context, note domain.Note) error {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.noteRepository.CreateNote(ctx, note)
}

func (usecase *noteUsecase) UpdateNote(c context.Context, note domain.Note) error {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.noteRepository.UpdateNote(ctx, note)
}

func (usecase *noteUsecase) DeleteNote(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	return usecase.noteRepository.DeleteNote(ctx, id)
}
