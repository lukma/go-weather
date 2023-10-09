package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type NoteRepository interface {
	GetNotes(ctx context.Context) ([]Note, error)
	CreateNote(ctx context.Context, note Note) error
	UpdateNote(ctx context.Context, note Note) error
	DeleteNote(ctx context.Context, id uuid.UUID) error
}

type NoteUsecase interface {
	GetNotes(ctx context.Context) ([]Note, error)
	CreateNote(ctx context.Context, note Note) error
	UpdateNote(ctx context.Context, note Note) error
	DeleteNote(ctx context.Context, id uuid.UUID) error
}
