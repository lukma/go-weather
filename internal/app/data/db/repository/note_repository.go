package repository

import (
	"context"
	"database/sql"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/google/uuid"
)

type noteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) domain.NoteRepository {
	return &noteRepository{
		db: db,
	}
}

func (repository *noteRepository) GetNotes(ctx context.Context) ([]domain.Note, error) {
	notes := []domain.Note{}
	var err error

	rows, err := repository.db.QueryContext(ctx, "select * from notes")
	if err != nil {
		return notes, err
	}

	for rows.Next() {
		var note domain.Note

		if err := rows.Scan(
			&note.ID,
			&note.Title,
			&note.Content,
			&note.CreatedAt,
			&note.UpdatedAt,
		); err != nil {
			return notes, err
		}

		notes = append(notes, note)
	}

	return notes, err
}

func (repository *noteRepository) CreateNote(ctx context.Context, note domain.Note) error {
	_, err := repository.db.ExecContext(ctx, "insert into notes(title, content) values (?, ?)", note.Title, note.Content)
	return err
}

func (repository *noteRepository) UpdateNote(ctx context.Context, note domain.Note) error {
	_, err := repository.db.ExecContext(ctx, "update notes set title = ?, content = ? where id = UUID_TO_BIN(?)", note.Title, note.Content, note.ID)
	return err
}

func (repository *noteRepository) DeleteNote(ctx context.Context, id uuid.UUID) error {
	_, err := repository.db.ExecContext(ctx, "delete from notes where id = UUID_TO_BIN(?)", id)
	return err
}
