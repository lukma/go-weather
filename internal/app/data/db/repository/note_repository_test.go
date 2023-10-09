package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"altechomega.lukma.dev/notes/internal/app/data/db/repository"
	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func TestGetNotes(t *testing.T) {
	selectQuery := "select (.+) from notes"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database")
	}
	defer db.Close()

	repository := repository.NewNoteRepository(db)

	t.Run("success", func(t *testing.T) {
		createdAt, _ := time.Parse("2006-01-02T15:04:05-0700", "2023-10-09T12:00:00.000124Z")
		updatedAt := createdAt.Add(time.Hour * time.Duration(2))
		expected := []domain.Note{
			{
				ID:        uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"),
				Title:     "Hello World",
				Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				CreatedAt: createdAt,
				UpdatedAt: &updatedAt,
			},
		}

		row := sqlmock.NewRows([]string{
			"id",
			"title",
			"content",
			"created_at",
			"updated_at",
		}).
			AddRow(
				expected[0].ID,
				expected[0].Title,
				expected[0].Content,
				expected[0].CreatedAt,
				expected[0].UpdatedAt,
			)
		mock.ExpectQuery(selectQuery).
			WillReturnRows(row)

		actual, err := repository.GetNotes(context.TODO())

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mock.ExpectQuery(selectQuery).
			WillReturnError(sql.ErrConnDone)

		actual, err := repository.GetNotes(context.TODO())

		assert.Empty(t, actual)
		assert.Error(t, err)
	})
}

func TestCreateNote(t *testing.T) {
	insertQuery := "insert into notes"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database")
	}
	defer db.Close()

	repository := repository.NewNoteRepository(db)

	t.Run("success", func(t *testing.T) {
		note := domain.Note{
			Title:   "Hello World",
			Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		}

		mock.ExpectExec(insertQuery).
			WithArgs(note.Title, note.Content).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repository.CreateNote(context.TODO(), note)

		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mock.ExpectExec(insertQuery).
			WillReturnError(sql.ErrConnDone)

		err := repository.CreateNote(context.TODO(), domain.Note{})

		assert.Error(t, err)
	})
}

func TestUpdateNote(t *testing.T) {
	updateQuery := "update notes"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database")
	}
	defer db.Close()

	repository := repository.NewNoteRepository(db)

	t.Run("success", func(t *testing.T) {
		note := domain.Note{
			ID:      uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"),
			Title:   "Hello World",
			Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		}

		mock.ExpectExec(updateQuery).
			WithArgs(note.Title, note.Content, note.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repository.UpdateNote(context.TODO(), note)

		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mock.ExpectExec(updateQuery).
			WillReturnError(sql.ErrConnDone)

		err := repository.UpdateNote(context.TODO(), domain.Note{})

		assert.Error(t, err)
	})
}

func TestDeleteNote(t *testing.T) {
	deleteQuery := "delete from notes"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database")
	}
	defer db.Close()

	repository := repository.NewNoteRepository(db)

	t.Run("success", func(t *testing.T) {
		id := uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")

		mock.ExpectExec(deleteQuery).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repository.DeleteNote(context.TODO(), id)

		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mock.ExpectExec(deleteQuery).
			WillReturnError(sql.ErrConnDone)

		err := repository.DeleteNote(context.TODO(), uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))

		assert.Error(t, err)
	})
}
